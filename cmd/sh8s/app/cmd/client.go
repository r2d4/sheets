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

	"github.com/r2d4/sh8s/cmd/sh8s/app/cmd/client"
	"github.com/r2d4/sh8s/pkg/datastore"
	"github.com/spf13/cobra"
)

func NewCmdClient(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "A client to help test the server",
	}

	datastore.InitDatastore("127.0.0.1:6379")

	cmd.AddCommand(client.NewCmdUploadFile(out))
	cmd.AddCommand(client.NewCmdUploadEnv(out))
	cmd.AddCommand(client.NewCmdRun(out))
	cmd.AddCommand(client.NewCmdMount(out))

	return cmd
}
