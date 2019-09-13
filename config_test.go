package main

import (
	"os"
	"testing"
)

func TestCreateConfigCmd(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "-listen", ":8888", "-endpoint", "http://localhost/fake-eta"}

	cfg := NewCmdConfig()

	equals(t, os.Args[2], cfg.Port)
	equals(t, os.Args[4], cfg.Endpoint)
}

func TestCreateConfigDefault(t *testing.T) {
	cfg := NewConfig()

	equals(t, DEFAULT_PORT, cfg.Port)
	equals(t, DEFAULT_ETA_ENDPOINT, cfg.Endpoint)
}
