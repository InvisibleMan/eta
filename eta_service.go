package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type EtaService struct {
}

func NewEtaService(cfg *Config) http.Handler {
	return &EtaService{}
}

type ApiError struct {
	HTTPStatus int
	Err        error
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

type Response struct {
	Error    string      `json:"error,omitempty"`
	Response interface{} `json:"response,omitempty"`
}

func responseError(w http.ResponseWriter, err error) {
	if errAPI, ok := err.(ApiError); ok {
		w.WriteHeader(errAPI.HTTPStatus)
		bytes, _ := json.Marshal(&Response{Error: errAPI.Error()})
		w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	bytes, _ := json.Marshal(&Response{Error: err.Error()})
	w.Write(bytes)
}

func responseSuccess(w http.ResponseWriter, resp *Response) {
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

func (s *EtaService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		responseError(w, ApiError{HTTPStatus: http.StatusMethodNotAllowed, Err: errors.New("unknown")})
		return
	}
	if r.URL.Path != "/eta" {
		responseError(w, ApiError{HTTPStatus: http.StatusNotFound, Err: errors.New("unknown method")})
		return
	}

	responseSuccess(w, &Response{Response: "10"})
}
