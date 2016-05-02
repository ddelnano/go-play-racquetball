// Package main provides ...
package racquetball

type Reservation struct {
	Day       string `json:"day,time.Weekday"`
	StartTime string `json:"time"`
	EndTime   string `json:"endTime"`
}

type GymClient interface {
	GetReservations() ([]Reservation, error)
	MakeReservation(*Reservation) ([]Reservation, error)
}
