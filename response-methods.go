package gopress

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	HttpResponse http.ResponseWriter
}

func (r Response) Json(data map[string]interface{}, statusCode int) {
	r.HttpResponse.WriteHeader(statusCode)
	if err := json.NewEncoder(r.HttpResponse).Encode(data); err != nil {
		JsonErrorResponse(r, http.StatusInternalServerError, err.Error())
	}
}