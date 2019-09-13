package main

import (
	"errors"
	"log"
	"time"
)

type NoContentError struct {
	Msg string
}

func (e NoContentError) Error() string {
	return e.Msg
}

type EtaModel struct {
	apiEta ApiEtaService
}

type EtaResult struct {
	CarID int64 `json:"carId"`
	Mins  int64 `json:"minutes"`
}

func NewModel(apiEta ApiEtaService) *EtaModel {
	return &EtaModel{apiEta: apiEta}
}

func minimumWithIndex(mins []int64) (int64, int) {
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

	cars, err := m.apiEta.GetCars(lat, lng, limit)
	dr1 := time.Since(start)

	if err != nil {
		return nil, err
	}

	if len(cars) < 1 {
		return nil, &NoContentError{Msg: "Cars not founts"}
	}

	t2 := time.Now()
	mins, err := m.apiEta.Predict(lat, lng, cars)
	dr2 := time.Since(t2)

	if err != nil {
		// log.Printf("Did calc Predict: %s.\n", dr2)
		return nil, err
	}

	if len(mins) < 1 {
		return nil, errors.New("Something wrong")
	}

	min, idx := minimumWithIndex(mins)

	log.Printf("Find took %s. Cars: %s, Predict: %s.\n", time.Since(start), dr1, dr2)
	return &EtaResult{CarID: cars[idx].ID, Mins: min}, nil
}
