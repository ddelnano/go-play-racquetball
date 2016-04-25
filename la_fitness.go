package main

import (
	"time"
)

const RESERVATION_THRESHOLD int = 14

type Reservation struct {
	Day       string       `json:"day,time.Weekday"`
	day       time.Weekday `json:-`
	Time      int          `json:"time"`
	Threshold int          `json:"threshold"`
}

func NewReservation() *Reservation {
	return &Reservation{
		Threshold: RESERVATION_THRESHOLD,
		Day:       "Sunday",
	}
}

func FindReleventReservations(r []Reservation) []Reservation {
	releventReservations := make([]Reservation, 0)
	// if needsReserved() {
	//     append(releventReservations
	// }
	return releventReservations
}
