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

package server

import (
	"net/http"

	"github.com/r2d4/sh8s/pkg/sh8s/runner"
	"github.com/r2d4/sh8s/pkg/sh8s/upload"
	"github.com/r2d4/sh8s/pkg/sh8s/util"
)

func Run() error {
	http.HandleFunc("/upload/file", util.BaseHandler{fn: upload.UploadFileHandler})
	http.HandleFunc("/upload/environment", util.BaseHandler{fn: upload.UploadEnvHandler})

	http.HandleFunc("/run", runner.Handler)

	return http.ListenAndServe(":8080", nil)
}
