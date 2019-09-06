package main

var (
	DEFAULT_PORT             string = ":8082"
	DEFAULT_CAR_ENDPOINT     string = "https://dev-api.wheely.com/fake-eta/cars"
	DEFAULT_PREDICT_ENDPOINT string = "https://dev-api.wheely.com/fake-eta/predict"
)

type Config struct {
	Port            string
	CarEndpoint     string
	PredictEndpoint string
}

// NewConfig is ...
func NewConfig() *Config {
	return &Config{
		Port:            DEFAULT_PORT,
		CarEndpoint:     DEFAULT_CAR_ENDPOINT,
		PredictEndpoint: DEFAULT_PREDICT_ENDPOINT,
	}
}
