package main

import (
	"flag"
)

var (
	DEFAULT_PORT         string = ":8082"
	DEFAULT_ETA_ENDPOINT string = "https://dev-api.wheely.com/fake-eta"
)

type Config struct {
	Port     string
	Endpoint string
}

// NewConfig is ...
func NewConfig() *Config {
	return &Config{
		Port:     DEFAULT_PORT,
		Endpoint: DEFAULT_ETA_ENDPOINT,
	}
}

func NewCmdConfig() *Config {
	cfg := Config{}

	flag.StringVar(&cfg.Port, "listen", DEFAULT_PORT, "HTTP listen spec")
	flag.StringVar(&cfg.Endpoint, "endpoint", DEFAULT_ETA_ENDPOINT, "Path to Todo-file")

	flag.Parse()

	return &cfg
}
