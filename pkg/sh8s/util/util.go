package util

import (
	"net/http"
)

// Handler is a wrapper that handles the http error codes
func Handler(f func(w http.ResponseWriter, r *http.Request) (int, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code, err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), code)
		}
	}
}
