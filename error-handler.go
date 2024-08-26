package gopress

import "net/http"

type ErrorHandler func(http.ResponseWriter, *http.Request, error)

var errorHandler ErrorHandler = func(w http.ResponseWriter, _ *http.Request, err error) {
	status, body := handleErr(err)
	w.WriteHeader(status)
	w.Write(body)
}

func (r *Router) setErrorHandler(handler ErrorHandler) {
	errorHandler = handler
}