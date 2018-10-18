/*
Copyright 2018 COMPANY

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"io"

	"github.com/pkg/errors"
	"github.com/r2d4/sh8s/pkg/sh8s/constants"
	"github.com/r2d4/sh8s/pkg/sh8s/serve"
	"github.com/spf13/cobra"
)

func NewCmdServe(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Runs the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServe(out, cmd)
		},
	}

	return cmd
}

func RunServe(out io.Writer, cmd *cobra.Command) error {
	if err := serve.Run(constants.DefaultPort); err != nil {
		return errors.Wrap(err, "running server")
	}
	return nil
}
