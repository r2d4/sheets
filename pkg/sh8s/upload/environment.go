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
	var b []byte
	if _, err := r.Body.Read(b); err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "reading request")
	}
	defer r.Body.Close()
	var req *api.UploadEnvironmentRequest
	if err := json.Unmarshal(b, &req); err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "unmarshaling uploadCellRequest")
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
	if err := datastore.DefaultDatastore.SetList(id, req.Range); err != nil {
		return nil, errors.Wrap(err, "setting list")
	}
	return &api.UploadCellResponse{
		Pointer: key,
	}, nil
}
