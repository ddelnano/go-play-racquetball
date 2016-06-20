package main

import (
	"errors"
	"testing"
	"time"

	rtime "github.com/ddelnano/racquetball/time"
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
	// assert.Equal(t, config.Reservations[0].StartTime, "6")
	assert.IsType(t, rtime.UTCTime{}, config.Reservations[0].StartTime)
	assert.Equal(t, config.Reservations[0].Duration, "60")
	assert.Equal(t, config.Reservations[0].ClubID, "1010")
	assert.Equal(t, config.Reservations[0].ClubDescription, "PITTSBURGH-PENN AVE")

	assert.Equal(t, config.Reservations[1].Day, "Thursday")
	// assert.Equal(t, config.Reservations[1].StartTime, "6")
	assert.IsType(t, rtime.UTCTime{}, config.Reservations[1].StartTime)
	assert.Equal(t, config.Reservations[1].Duration, "60")
	assert.Equal(t, config.Reservations[1].ClubID, "1010")
	assert.Equal(t, config.Reservations[1].ClubDescription, "PITTSBURGH-PENN AVE")
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

func TestReservationsForWeek(t *testing.T) {
	c := Configuration{
		Reservations: []Reservation{
			Reservation{
				Day:       time.Now().Weekday().String(),
				StartTime: rtime.UTCTime{time.Now()},
			},
			Reservation{
				Day:       ((time.Now().Weekday() + 1) % 7).String(),
				StartTime: rtime.UTCTime{time.Now()},
			},
		},
	}
	res := c.ReservationsForWeek()
	assert.Equal(t, 1, len(res))
	ymdFormat := "2006-02-01"
	assert.Equal(t, time.Now().AddDate(0, 0, 7).Format(ymdFormat), res[0].StartTime.Format(ymdFormat))
}
