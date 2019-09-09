package main

import (
	carsClient "eta/api-cars/client"
	carsOps "eta/api-cars/client/operations"
	"eta/api-cars/models"
	carsModels "eta/api-cars/models"
	predictClient "eta/api-predict/client"
	predictOps "eta/api-predict/client/operations"
	predictModels "eta/api-predict/models"
	"net/url"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type ApiEtaService interface {
	GetCars(lat float64, lng float64, limit int64) ([]models.Car, error)
	Predict(lat float64, lng float64, cars []models.Car) ([]int64, error)
}

// const MAX_DURATION_CARS time.Duration = 301 * time.Millisecond

const (
	MAX_DURATION_CARS    time.Duration = 301 * time.Millisecond
	MAX_DURATION_PREDICT time.Duration = 501 * time.Millisecond
)

type ApiSwagger struct {
	carsService    *carsClient.CarsService
	predictService *predictClient.PredictService
}

func NewSwaggerApi(endpoint string) ApiEtaService {
	api := ApiSwagger{}
	url, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	transport := httptransport.New(url.Host, url.Path, []string{url.Scheme})
	api.carsService = carsClient.New(transport, strfmt.Default)
	api.predictService = predictClient.New(transport, strfmt.Default)

	return &api
}

func (sw *ApiSwagger) GetCars(lat float64, lng float64, limit int64) ([]models.Car, error) {
	res, err := sw.carsService.Operations.GetCars(carsOps.
		NewGetCarsParams().
		WithTimeout(MAX_DURATION_CARS).
		WithLat(lat).
		WithLng(lng).
		WithLimit(limit))

	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func carsToPredictOpts(lat float64, lng float64, cars []carsModels.Car) predictOps.PredictBody {
	src := make([]predictModels.Position, len(cars))
	for idx, c := range cars {
		src[idx] = predictModels.Position{Lat: c.Lat, Lng: c.Lng}
	}

	return predictOps.PredictBody{
		Source: src,
		Target: predictModels.Position{Lat: lat, Lng: lng},
	}
}

func (sw *ApiSwagger) Predict(lat float64, lng float64, cars []models.Car) ([]int64, error) {
	list := carsToPredictOpts(lat, lng, cars)

	res, err := sw.predictService.Operations.Predict(predictOps.
		NewPredictParams().
		WithTimeout(MAX_DURATION_PREDICT).
		WithPositionList(list))

	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}
