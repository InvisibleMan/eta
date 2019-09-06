package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	client = &http.Client{Timeout: time.Second}
)

func TestApis(t *testing.T) {

	handler := NewEtaService(nil)
	ts := httptest.NewServer(handler)

	req, _ := http.NewRequest("GET", ts.URL, nil)
	resp, _ := client.Do(req)

	equals(t, resp.StatusCode, http.StatusNotFound)

	req, _ = http.NewRequest("POST", ts.URL+"/eta", nil)
	resp, _ = client.Do(req)

	equals(t, resp.StatusCode, http.StatusMethodNotAllowed)
}
