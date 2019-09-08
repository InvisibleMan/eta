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

func TestApis(t *testing.T) {
	handler := NewEtaService(NewConfig())
	ts := httptest.NewServer(handler)

	req, _ := http.NewRequest("GET", ts.URL, nil)
	resp, _ := clientHttp.Do(req)

	equals(t, resp.StatusCode, http.StatusNotFound)

	req, _ = http.NewRequest("POST", ts.URL+"/eta", nil)
	resp, _ = clientHttp.Do(req)

	equals(t, resp.StatusCode, http.StatusMethodNotAllowed)
}
