package upload

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/r2d4/sh8s/pkg/datastore"
	"github.com/r2d4/sh8s/pkg/sh8s/api"
)

func UploadEnvHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	decoder := json.NewDecoder(r.Body)

	var req *api.UploadEnvironmentRequest
	if err := decoder.Decode(&req); err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "unmarshaling upload run request")
	}
	resp, err := UploadEnvironment(req)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "uploading file")
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

func UploadEnvironment(req *api.UploadEnvironmentRequest) (*api.UploadCellResponse, error) {
	key := fmt.Sprintf("environment:%s", req.ID)
	if err := datastore.DefaultDatastore.SetList(key, req.Range); err != nil {
		return nil, errors.Wrap(err, "setting list")
	}
	return &api.UploadCellResponse{
		Pointer: key,
	}, nil
}
