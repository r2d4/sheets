package upload

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/r2d4/sh8s/pkg/datastore"
	"github.com/r2d4/sh8s/pkg/sh8s/api"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	decoder := json.NewDecoder(r.Body)

	var req *api.UploadFileRequest
	if err := decoder.Decode(&req); err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "unmarshaling upload run request")
	}
	resp, err := UploadFile(req)
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

func UploadFile(req *api.UploadFileRequest) (*api.UploadCellResponse, error) {
	key := fmt.Sprintf("file:%s", req.Filename)
	if err := datastore.DefaultDatastore.Set(key, req.Data); err != nil {
		return nil, errors.Wrap(err, "storing file")
	}
	return &api.UploadCellResponse{
		Pointer: key,
	}, nil
}
