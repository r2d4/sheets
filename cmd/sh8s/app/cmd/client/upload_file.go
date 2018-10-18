package client

import (
	"io"

	"github.com/r2d4/sh8s/pkg/sh8s/api"
	"github.com/spf13/cobra"
)

var filename, data string

func NewCmdUploadFile(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload-file",
		Short: "Uploads a file to the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUploadFile(out, cmd)
		},
	}

	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename for the request")
	cmd.Flags().StringVarP(&data, "data", "d", "", "Data for the request")

	return cmd
}

func RunUploadFile(out io.Writer, cmd *cobra.Command) error {
	req := &api.UploadFileRequest{
		Filename: filename,
		Data:     data,
	}
	return post(out, "upload/file", req)
}
