// Package main provides ...
package racquetball

import rtime "github.com/ddelnano/racquetball/time"

type Reservation struct {
	Day             string        `json:"day,time.Weekday"`
	StartTime       rtime.UTCTime `json:"time"`
	EndTime         string        `json:"endTime"`
	Duration        string
	ClubID          string
	ClubDescription string
	Test            rtime.UTCTime
}

type GymClient interface {
	GetReservations() ([]Reservation, error)
	MakeReservation(*Reservation) ([]Reservation, error)
}
