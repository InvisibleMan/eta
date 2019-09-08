package main

import "testing"

type CaseError struct {
	Form     map[string][]string
	ErrorMsg string
	Lat      float64
	Lng      float64
	Limit    int64
}

func TestFetchError(t *testing.T) {
	params := EtaParams{}

	cases := []CaseError{
		CaseError{
			Form:     map[string][]string{},
			ErrorMsg: "lat required",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"-91"}},
			ErrorMsg: "lat must be greater than -90",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"91"}},
			ErrorMsg: "lat must be less than 90",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"55"}},
			ErrorMsg: "lng required",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"55"}, "lng": []string{"-192"}},
			ErrorMsg: "lng must be greater than -180",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"55"}, "lng": []string{"190"}},
			ErrorMsg: "lng must be less than 180",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"55"}, "lng": []string{"37"}, "limit": []string{"sd"}},
			ErrorMsg: "wrong limit",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"55"}, "lng": []string{"37"}, "limit": []string{"0"}},
			ErrorMsg: "limit must be greater than 1",
		},
		CaseError{
			Form:     map[string][]string{"lat": []string{"55"}, "lng": []string{"37"}, "limit": []string{"200"}},
			ErrorMsg: "limit must be less than 100",
		},
	}

	for _, cs := range cases {
		err := params.Fetch(cs.Form)
		hasErr(t, err, cs.ErrorMsg)
	}
}

func TestFetchSuccess(t *testing.T) {
	cases := []CaseError{
		CaseError{
			Form:  map[string][]string{"lat": []string{"55.752992"}, "lng": []string{"37.618333"}},
			Lat:   55.752992,
			Lng:   37.618333,
			Limit: 3,
		},
		CaseError{
			Form:  map[string][]string{"lat": []string{"55.752992"}, "lng": []string{"37.618333"}, "limit": []string{""}},
			Lat:   55.752992,
			Lng:   37.618333,
			Limit: 3,
		},
		CaseError{
			Form:  map[string][]string{"lat": []string{"55.752992"}, "lng": []string{"37.618333"}, "limit": []string{"32"}},
			Lat:   55.752992,
			Lng:   37.618333,
			Limit: 32,
		},
	}

	for _, cs := range cases {
		params := EtaParams{}
		err := params.Fetch(cs.Form)

		ok(t, err)
		equals(t, params.Lat, cs.Lat)
		equals(t, params.Lng, cs.Lng)
		equals(t, params.Limit, cs.Limit)
	}
}
