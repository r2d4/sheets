package client

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/r2d4/sh8s/pkg/datastore"
	"github.com/spf13/cobra"
)

var (
	env string
)

func NewCmdMount(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mount",
		Short: "Mounts a folder",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunMount(out, cmd)
		},
	}

	cmd.Flags().StringVarP(&env, "environment", "e", "", "Name of the folder. Must begin with environment:")
	return cmd
}

func RunMount(out io.Writer, cmd *cobra.Command) error {
	ds := datastore.DefaultDatastore
	m, err := ds.List(env)
	if err != nil {
		return errors.Wrap(err, "listing files")
	}
	fmt.Println(m)
	for dest, contents := range m {
		if err := os.MkdirAll(filepath.Base(dest), os.ModePerm); err != nil {
			return errors.Wrap(err, "making directories")
		}
		f, err := os.Open(dest)
		if err != nil {
			return errors.Wrap(err, "opening file to write")
		}
		defer f.Close()
		if _, err := io.Copy(f, bytes.NewBufferString(contents)); err != nil {
			return errors.Wrap(err, "writing file")
		}
	}
	return nil
}
