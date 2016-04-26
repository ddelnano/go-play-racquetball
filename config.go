// Package config provides ...
package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Reservations []Reservation `json:"reservations"`
}

func Load(filepath string, validate ReservationValidation) (*Configuration, error) {
	// TODO: Should handle the Result type
	_, err := validate(filepath)

	if err != nil {
		return nil, err
	}
	var config Configuration
	file, _ := ioutil.ReadFile(filepath)
	err = json.Unmarshal(file, &config)
	// TODO: Should be able to test this condition
	if err != nil {
		panic(err)
	}

	return &config, err
}
