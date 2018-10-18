package runner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/r2d4/sh8s/pkg/sh8s/api"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RunHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	var b []byte
	if _, err := r.Body.Read(b); err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "reading request")
	}
	defer r.Body.Close()
	var req *api.RunRequest
	if err := json.Unmarshal(b, &req); err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "unmarshaling uploadCellRequest")
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
	return nil, nil
}

type podFunction func(args []string) (*v1.Pod, error)

func unixTools(args []string) (*v1.Pod, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("unix_tools requires at least 1 argument")
	}

	return &v1.Pod{
		ObjectMeta: meta_v1.ObjectMeta{
			GenerateName: "unix-tools-",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Image:   "alpine:3.6",
					Command: args,
				},
			},
		},
	}, nil
}

var defaultFunctions = map[string]podFunction{
	"unix_tools": unixTools,
}
