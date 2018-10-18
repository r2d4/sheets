package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"github.com/r2d4/sh8s/pkg/sh8s/api"
	"github.com/r2d4/sh8s/pkg/sh8s/kubernetes"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	core_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func RunHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	decoder := json.NewDecoder(r.Body)

	var req *api.RunRequest
	if err := decoder.Decode(&req); err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "unmarshaling upload run request")
	}
	resp, err := Run(req)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "running pod request")
	}
	out, err := json.Marshal(resp)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "marshaling upload response")
	}
	if _, err := w.Write(out); err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "writing response")
	}
	return http.StatusOK, nil
}

func Run(req *api.RunRequest) (*api.RunResponse, error) {
	fn, ok := defaultFunctions[req.Function]
	if !ok {
		return nil, fmt.Errorf("unknown function")
	}
	pod, err := fn(req.Args)
	if err != nil {
		return nil, errors.Wrap(err, "evaluating pod function")
	}
	client, err := kubernetes.GetClientset()
	if err != nil {
		return nil, errors.Wrap(err, "getting k8s client")
	}
	pods := client.CoreV1().Pods("default")
	out, err := pods.Create(pod)
	if err != nil {
		return nil, errors.Wrap(err, "creating pod")
	}
	var b bytes.Buffer
	if err := getLogsForPod(pods, out.GetName(), &b); err != nil {
		return nil, errors.Wrap(err, "getting logs for pod")
	}

	go func() {
		if err := pods.Delete(out.GetName(), &meta_v1.DeleteOptions{}); err != nil {
			logrus.Warnf("cleaning up pod: ", out.GetName())
		}
	}()

	return &api.RunResponse{
		Data: b.String(),
	}, nil
}

func getLogsForPod(client core_v1.PodInterface, podName string, w io.Writer) error {
	opts := meta_v1.ListOptions{
		LabelSelector: "app=excel",
	}
	watchPod, err := client.Watch(opts)
	if err != nil {
		return err
	}
	logOpts := &v1.PodLogOptions{
		Follow: true,
	}
	for {
		event := <-watchPod.ResultChan()
		switch e := event.Type; e {
		case watch.Added:
		case watch.Modified:
			pod, ok := event.Object.(*v1.Pod)
			if !ok || pod.Name != podName {
				continue
			}
			// Either succeeded or failed, we still want the logs
			if pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
				logStream, err := client.GetLogs(pod.Name, logOpts).Stream()
				if err != nil {
					return err
				}
				defer logStream.Close()

				io.Copy(w, logStream)
				return nil
			}
		case watch.Deleted:
			return nil
		}
	}
}

type podFunction func(args []string) (*v1.Pod, error)

func unixTools(args []string) (*v1.Pod, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("unix_tools requires at least 1 argument")
	}

	return &v1.Pod{
		ObjectMeta: meta_v1.ObjectMeta{
			GenerateName: "unix-tools-",
			Labels: map[string]string{
				"app": "excel",
			},
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []v1.Container{
				{
					Image:   "alpine:3.6",
					Name:    "alpine",
					Command: args,
				},
			},
		},
	}, nil
}

var defaultFunctions = map[string]podFunction{
	"unix_tools": unixTools,
}
