package client

import (
	"io"

	"github.com/r2d4/sh8s/pkg/sh8s/api"
	"github.com/spf13/cobra"
)

var (
	id  string
	set []string
)

func NewCmdUploadEnv(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload-env",
		Short: "Uploads a environment to the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUploadEnv(out, cmd)
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Name of the set")
	cmd.Flags().StringArrayVarP(&set, "set", "s", nil, "Set of keys")

	return cmd
}

func RunUploadEnv(out io.Writer, cmd *cobra.Command) error {
	req := &api.UploadEnvironmentRequest{
		ID:    id,
		Range: set,
	}
	return post(out, "upload/environment", req)
}
