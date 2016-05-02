// Package config provides ...
package racquetball

import (
	"encoding/json"
	"io/ioutil"
	"time"
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

func (c *Configuration) DailyReservations() []Reservation {
	daily := make([]Reservation, 0)
	today := time.Now().Weekday().String()
	for _, v := range c.Reservations {
		if v.Day == today {
			daily = append(daily, v)
		}
	}
	return daily
}
