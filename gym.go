// Package main provides ...
package main

import "time"

type Reservation struct {
	Day  string       `json:"day,time.Weekday"`
	day  time.Weekday `json:-`
	Time string       `json:"time"`
}

type GymClient interface {
	GetReservations() ([]Reservation, error)
	MakeReservation(*Reservation) ([]Reservation, error)
}
