package main

import (
	"time"
)

const RESERVATION_THRESHOLD int = 14

type Reservation struct {
	day       time.Weekday
	time      int
	threshold int
}

// func (r *Reservation) UnmarshalJSON(by []byte) erorr {

// }

func NewReservation() *Reservation {
	return &Reservation{
		threshold: RESERVATION_THRESHOLD,
		day:       time.Now().Weekday(),
	}
}

func FindReleventReservations(r []Reservation) []Reservation {
	releventReservations := make([]Reservation, 0)
	// if needsReserved() {
	//     append(releventReservations
	// }
	return releventReservations
}
