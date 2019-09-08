package main

import (
	"errors"
	"net/url"
	"strconv"
)

const DEFAULT_LIMIT int64 = 3

type EtaParams struct {
	Lat   float64
	Lng   float64
	Limit int64
}

func (ep *EtaParams) Fetch(f url.Values) error {
	lat, err := strconv.ParseFloat(f.Get("lat"), 64)

	if err != nil {
		return errors.New("lat required")
	}

	if lat < -90 {
		return errors.New("lat must be greater than -90")
	}
	if lat > 90 {
		return errors.New("lat must be less than 90")
	}

	lng, err := strconv.ParseFloat(f.Get("lng"), 64)

	if err != nil {
		return errors.New("lng required")
	}

	if lng < -180 {
		return errors.New("lng must be greater than -180")
	}
	if lng > 180 {
		return errors.New("lng must be less than 180")
	}

	limit := DEFAULT_LIMIT

	slimit := f.Get("limit")
	if len(slimit) != 0 {
		limit, err = strconv.ParseInt(slimit, 10, 32)
		if err != nil {
			return errors.New("wrong limit")
		}
	}

	if limit < 1 {
		return errors.New("limit must be greater than 1")
	}
	if limit > 100 {
		return errors.New("limit must be less than 100")
	}

	ep.Lat = lat
	ep.Lng = lng
	ep.Limit = limit

	return nil
}
