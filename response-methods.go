package gopress

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Response struct {
	HttpResponse http.ResponseWriter
}

func (r Response) Json(data interface{}, statusCode int) {
	r.HttpResponse.Header().Set("Content-Type", "application/json")
	r.HttpResponse.WriteHeader(statusCode)
 
	if err := json.NewEncoder(r.HttpResponse).Encode(data); err != nil {
		http.Error(r.HttpResponse, `{"error": "Failed to encode JSON"}`, http.StatusInternalServerError)
	}
}

func (r Response) Download(path, name string) {
	file, err := os.Open(path)
	if err != nil {
		JsonErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	r.HttpResponse.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")
    r.HttpResponse.Header().Set("Content-Type", "application/octet-stream")
	
	if _, err := io.Copy(r.HttpResponse, file); err != nil {
		JsonErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (r Response) SendFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		JsonErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	if _, err := io.Copy(r.HttpResponse, file); err != nil {
		JsonErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
}