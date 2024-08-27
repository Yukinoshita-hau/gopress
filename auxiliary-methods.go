package gopress

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Request struct {
	HttpRequest *http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		HttpRequest: r,
	}
}

func (r *Request) GetHeader(key string) string {
	return r.HttpRequest.Header.Get(key)
}

func (r *Request) GetBody() ([]byte, error) {
	body, err := io.ReadAll(r.HttpRequest.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *Request) GetBodyAndConvertInJson() (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(r.HttpRequest.Body).Decode(&data)
	if err != nil {
		return nil, errors.New("failed convert body to json")
	}
	return data, nil
}

func (r *Request) GetParam(key string) string {
	return r.HttpRequest.Context().Value(key).(string)
}