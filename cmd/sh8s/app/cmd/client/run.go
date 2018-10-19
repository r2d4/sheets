package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/r2d4/sh8s/pkg/sh8s/api"
	"github.com/r2d4/sh8s/pkg/sh8s/constants"
	"github.com/spf13/cobra"
)

var (
	function string
	args     []string
)

func NewCmdRun(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Runs a pod request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunPod(out, cmd)
		},
	}

	cmd.Flags().StringVarP(&function, "function", "f", "", "Name of the function")
	cmd.Flags().StringArrayVarP(&args, "args", "a", nil, "Set of args")

	return cmd
}

func RunPod(out io.Writer, cmd *cobra.Command) error {
	req := &api.RunRequest{
		Function: function,
		Args:     args,
	}
	return post(out, "run", req)
}

func post(w io.Writer, route string, req interface{}) error {
	data, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "marshaling data")
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/%s", constants.DefaultPortAndAddress, route), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return errors.Wrap(err, "post:")
	}
	fmt.Println(string(data))
	io.Copy(w, resp.Body)
	w.Write([]byte("\n"))
	return resp.Body.Close()
}
