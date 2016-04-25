// Package config provides ...
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	Reservations []Reservation `json:"reservations"`
}

func Load(filepath string) (*Configuration, error) {
	file, err := ioutil.ReadFile(filepath)
	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return &config, err
}
