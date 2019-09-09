package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	clientHttp = &http.Client{Timeout: time.Second}
)

type CaseService struct {
	Method string
	Url    string
	Status int
}

func TestApis(t *testing.T) {
	handler := NewEtaService(nil)
	ts := httptest.NewServer(handler)

	cases := []CaseService{
		CaseService{Method: "GET", Url: ts.URL, Status: http.StatusNotFound},
		CaseService{Method: "POST", Url: ts.URL + "/eta", Status: http.StatusMethodNotAllowed},
		CaseService{Method: "GET", Url: ts.URL + "/eta", Status: http.StatusBadRequest},
	}

	for _, cs := range cases {
		req, _ := http.NewRequest(cs.Method, cs.Url, nil)
		resp, _ := clientHttp.Do(req)
		equals(t, resp.StatusCode, cs.Status)
	}
}
