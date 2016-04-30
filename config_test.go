package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

func mockReservationValidation(file string) (*gojsonschema.Result, error) {
	err := errors.New("validation has failed")
	return nil, err
}

func reservationForToday() Reservation {
	today := time.Now().Weekday().String()
	return Reservation{Day: today}
}

func reservationNotToday() Reservation {
	notToday := ((time.Now().Weekday() + 1) % 7).String()
	return Reservation{Day: notToday}
}

func mockReservationValiationPass(file string) (*gojsonschema.Result, error) {
	return &gojsonschema.Result{}, nil
}

func TestLoadReturnsErrorIfValidationFails(t *testing.T) {
	config, err := Load("./sample.json", mockReservationValidation)
	assert.Nil(t, config)
	assert.Equal(t, "validation has failed", err.Error())
}

func TestLoadReturnsConfigurationStructIfValidationPasses(t *testing.T) {
	config, err := Load("./sample.json", mockReservationValiationPass)
	assert.NotNil(t, config)
	assert.Equal(t, config.Reservations[0].Day, "Wednesday")
	assert.Equal(t, config.Reservations[1].Day, "Thursday")
	assert.Equal(t, config.Reservations[0].Time, "6")
	assert.Equal(t, config.Reservations[1].Time, "6")
	assert.Nil(t, err)
}

func TestDailyReservations(t *testing.T) {
	config := Configuration{
		Reservations: []Reservation{
			reservationForToday(),
			reservationForToday(),
			reservationNotToday(),
		},
	}
	assert.Equal(t, 2, len(config.DailyReservations()))
}
