package gopress

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Request struct {
	httpRequest *http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		httpRequest: r,
	}
}

func (r *Request) GetHeader(key string) string {
	return r.httpRequest.Header.Get(key)
}

func (r *Request) GetBody() ([]byte, error) {
	body, err := io.ReadAll(r.httpRequest.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *Request) GetBodyAndConvertInJson() (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(r.httpRequest.Body).Decode(&data)
	if err != nil {
		return nil, errors.New("failed convert body to json")
	}
	return data, nil
}

func (r *Request) GetParam(key string) string {
	return r.httpRequest.Context().Value(key).(string)
}