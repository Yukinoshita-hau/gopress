package gopress

import (
	"encoding/json"
)

func JsonErrorResponse(w Response, status int, message string) {
	w.HttpResponse.WriteHeader(status)
	json.NewEncoder(w.HttpResponse).Encode(map[string]string{
		"error": message,
	})
}

type ErrorHandler func(Response, *Request, error)

var errorHandler ErrorHandler = func(w Response, _ *Request, err error) {
	status, body := handleErr(err)
	JsonErrorResponse(w, status, string(body))
}

func (r *Router) SetErrorHandler(handler ErrorHandler) {
	errorHandler = handler
}