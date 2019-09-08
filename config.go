package main

import "net/url"

var (
	DEFAULT_PORT         string = ":8082"
	DEFAULT_ETA_ENDPOINT string = "https://dev-api.wheely.com/fake-eta"
)

type Config struct {
	Port        string
	EtaEndpoint *url.URL
}

// NewConfig is ...
func NewConfig() *Config {
	url, _ := url.Parse(DEFAULT_ETA_ENDPOINT)

	return &Config{
		Port:        DEFAULT_PORT,
		EtaEndpoint: url,
	}
}
