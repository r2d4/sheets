package util

import (
	"net/http"
)

type BaseHandler struct {
	fn func(w http.ResponseWriter, r *http.Request) (int, error)
}

// Handler is just a wrapper for the actual handle function
func (b *BaseHandler) Handler(w http.ResponseWriter, r *http.Request) {
	code, err := b.fn(w, r)
	if err != nil {
		http.Error(w, err.Error(), code)
	}
}
