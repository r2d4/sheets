package api

// UploadFileRequest uploads a file to redis
type UploadFileRequest struct {
	Filename string
	Data     string
}

// UploadCellResponse returns the key that points to an uploaded file
type UploadCellResponse struct {
	Pointer string
}

// UploadEnvironmentRequest associated a set of pointers to a list
type UploadEnvironmentRequest struct {
	ID    string
	Range []string
}

// RunRequest runs a function with the supplied arguments
// on the cluster
type RunRequest struct {
	Function string
	Args     []string
}

// RunResponse returns the logs from the job that was ran
type RunResponse struct {
	Data string
}
