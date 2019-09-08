package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	DEFAULT_LAT = 55.752992
	DEFAULT_LNG = 37.618333
)

type EtaService struct {
	eta   *EtaModel
	cache Cache
}

func NewEtaService(cfg *Config) *EtaService {
	return &EtaService{eta: NewModel(cfg), cache: NewDummyCache(100, 30)}
}

type ApiError struct {
	HTTPStatus int
	Err        error
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

type Response struct {
	Error string      `json:"error,omitempty"`
	Eta   interface{} `json:"response,omitempty"`
}

func responseError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-type", "application/json")
	if errAPI, ok := err.(ApiError); ok {
		w.WriteHeader(errAPI.HTTPStatus)
		bytes, _ := json.Marshal(&Response{Error: errAPI.Error()})
		_, _ = w.Write(bytes)
		return
	} else if netError, ok := err.(*url.Error); ok {
		status := http.StatusForbidden
		if netError.Timeout() {
			status = http.StatusRequestTimeout
		}
		w.WriteHeader(status)
		bytes, _ := json.Marshal(&Response{Error: netError.Error()})
		_, _ = w.Write(bytes)
		return
	} else if errDomain, ok := err.(NoContentError); ok {
		w.WriteHeader(http.StatusNoContent)
		bytes, _ := json.Marshal(&Response{Error: errDomain.Error()})
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	bytes, _ := json.Marshal(&Response{Error: err.Error()})
	_, _ = w.Write(bytes)
}

func responseSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	raw, _ := json.Marshal(data)
	_, _ = w.Write(raw)
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

	params := &EtaParams{}

	err := params.Fetch(r.URL.Query())
	if err != nil {
		responseError(w, ApiError{HTTPStatus: http.StatusBadRequest, Err: err})
		return
	}

	key := r.URL.RawQuery
	cachedRes, exists := s.cache.Get(key)

	if !exists {
		res, err := s.eta.Find(params.Lat, params.Lng, params.Limit)
		if err != nil {
			responseError(w, err)
			return
		}
		cachedRes = &Response{Eta: res}
		s.cache.Put(key, cachedRes)
	}

	responseSuccess(w, cachedRes)
}

func (s EtaService) WarmUp() {
	_, _ = s.eta.Find(DEFAULT_LAT, DEFAULT_LNG, DEFAULT_LIMIT)
}
