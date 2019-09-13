package main

import (
	"errors"
	"eta/api-cars/models"
	"testing"
)

type CarsFunc func(float64, float64, int64) ([]models.Car, error)
type PredictFunc func(lat float64, lng float64, cars []models.Car) ([]int64, error)

type EtaServiceMock struct {
	CarsF    CarsFunc
	PredictF PredictFunc
}

func (m *EtaServiceMock) GetCars(lat float64, lng float64, limit int64) ([]models.Car, error) {
	return m.CarsF(lat, lng, limit)
}

func (m *EtaServiceMock) Predict(lat float64, lng float64, cars []models.Car) ([]int64, error) {
	return m.PredictF(lat, lng, cars)
}

func createCarsF(cars []models.Car, err error) CarsFunc {
	return func(lat float64, lng float64, limit int64) ([]models.Car, error) {
		return cars, err
	}
}

func createPredictF(mins []int64, err error) PredictFunc {
	return func(lat float64, lng float64, cars []models.Car) ([]int64, error) {
		return mins, err
	}
}

func TestMinimum(t *testing.T) {
	in := []int64{5, 4, 3, 2, 1}
	exIdx := 4

	min, idx := minimumWithIndex(in)

	equals(t, idx, 4)
	equals(t, min, in[exIdx])

	in = []int64{4}
	min, idx = minimumWithIndex(in)

	exIdx = 0
	equals(t, idx, exIdx)
	equals(t, min, in[exIdx])
}

func TestFind(t *testing.T) {
	var lat float64 = DEFAULT_LAT
	var lng float64 = DEFAULT_LNG
	var limit int64 = DEFAULT_LIMIT

	model := NewModel(&EtaServiceMock{
		CarsF: createCarsF(nil, errors.New("Timeout")),
	})

	_, err := model.Find(lat, lng, limit)
	hasErr(t, err, "Timeout")

	model = NewModel(&EtaServiceMock{
		CarsF: createCarsF([]models.Car{}, nil),
	})

	_, err = model.Find(lat, lng, limit)
	hasErr(t, err, "Cars not founts")

	model = NewModel(&EtaServiceMock{
		CarsF:    createCarsF([]models.Car{models.Car{}}, nil),
		PredictF: createPredictF(nil, errors.New("")),
	})

	_, err = model.Find(lat, lng, limit)
	hasErr(t, err, "")

	model = NewModel(&EtaServiceMock{
		CarsF:    createCarsF([]models.Car{models.Car{}}, nil),
		PredictF: createPredictF([]int64{}, nil),
	})

	_, err = model.Find(lat, lng, limit)
	hasErr(t, err, "Something wrong")

	model = NewModel(&EtaServiceMock{
		CarsF:    createCarsF([]models.Car{models.Car{ID: 10}, models.Car{ID: 20}, models.Car{ID: 30}}, nil),
		PredictF: createPredictF([]int64{5, 4, 2}, nil),
	})

	res, err := model.Find(lat, lng, limit)
	ok(t, err)
	equals(t, &EtaResult{CarID: 30, Mins: 2}, res)

}
