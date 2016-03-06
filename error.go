package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e apiError) Error() string {
	return e.Message
}

func (e apiError) HTTPError(w http.ResponseWriter) {
	code := e.Code
	out, err := json.Marshal(e)
	if err != nil {
		out = []byte(fmt.Sprintf(`{"code": %d, "message": "An error occoured during your request, and during the rendering of the error: %s"}`, http.StatusInternalServerError, err.Error()))
		code = http.StatusInternalServerError
	}
	// Content-type should already be set by handleRequest
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, string(out))
}
