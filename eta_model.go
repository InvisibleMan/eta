package main

import (
	"errors"
	"log"
	"time"

	carsClient "eta/api-cars/client"
	carsOps "eta/api-cars/client/operations"
	carsModels "eta/api-cars/models"

	predictClient "eta/api-predict/client"
	predictOps "eta/api-predict/client/operations"
	predictModels "eta/api-predict/models"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	MAX_DURATION_CARS    time.Duration = 301 * time.Millisecond
	MAX_DURATION_PREDICT time.Duration = 501 * time.Millisecond
)

type NoContentError struct {
	Msg string
}

func (e NoContentError) Error() string {
	return e.Msg
}

type EtaModel struct {
	carsApi    *carsClient.CarsService
	predictApi *predictClient.PredictService
}

type EtaResult struct {
	CarID int64 `json:"carId"`
	Mins  int64 `json:"minutes"`
}

func NewModel(cfg *Config) *EtaModel {
	model := &EtaModel{}

	transport := httptransport.New(cfg.EtaEndpoint.Host, cfg.EtaEndpoint.Path, []string{cfg.EtaEndpoint.Scheme})
	model.carsApi = carsClient.New(transport, strfmt.Default)
	model.predictApi = predictClient.New(transport, strfmt.Default)

	return model
}

func (m *EtaModel) CarsToPredictOpts(lat float64, lng float64, cars []carsModels.Car) predictOps.PredictBody {
	src := make([]predictModels.Position, len(cars))
	for idx, c := range cars {
		src[idx] = predictModels.Position{Lat: c.Lat, Lng: c.Lng}
	}

	return predictOps.PredictBody{
		Source: src,
		Target: predictModels.Position{Lat: lat, Lng: lng},
	}
}

func (m *EtaModel) getMinimumIn(mins []int64) (int64, int) {
	var idx int = 0
	var min int64 = mins[idx]

	for i, m := range mins {
		if m < min {
			min = m
			idx = i
		}
	}
	return min, idx
}

func (m *EtaModel) Find(lat float64, lng float64, limit int64) (*EtaResult, error) {
	start := time.Now()

	res1, err := m.carsApi.Operations.GetCars(carsOps.
		NewGetCarsParams().
		WithTimeout(MAX_DURATION_CARS).
		WithLat(lat).
		WithLng(lng).
		WithLimit(limit))
	dr1 := time.Since(start)

	if err != nil {
		log.Printf("Did R1: %s.\n", dr1)
		return nil, err
	}
	cars := res1.GetPayload()

	if len(cars) < 1 {
		return nil, &NoContentError{Msg: "Cars not founts"}
	}

	list := m.CarsToPredictOpts(lat, lng, cars)

	t2 := time.Now()
	res2, err := m.predictApi.Operations.Predict(predictOps.
		NewPredictParams().
		WithTimeout(MAX_DURATION_PREDICT).
		WithPositionList(list))
	dr2 := time.Since(t2)
	if err != nil {
		log.Printf("Did R2: %s.\n", dr2)
		return nil, err
	}

	mins := res2.GetPayload()
	if len(mins) < 1 {
		return nil, errors.New("Something wrong")
	}

	min, idx := m.getMinimumIn(mins)

	log.Printf("Find took %s. R1: %s, R2: %s.\n", time.Since(start), dr1, dr2)
	return &EtaResult{CarID: cars[idx].ID, Mins: min}, nil
}
