package racquetball

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	rtime "github.com/ddelnano/racquetball/time"
	"github.com/joho/godotenv"
)

func TestMakingReservations(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	laUsername := os.Getenv("LA_USERNAME")
	laPassword := os.Getenv("LA_PASSWORD")
	fmt.Println("Awe yea go", laUsername, laPassword)

	cred := Credentials{Username: laUsername, Password: laPassword}
	baseUrl, _ := url.Parse("https://publicapi.lafitness.com")
	httpClient := http.DefaultClient
	laClient := NewLaFitnessClient(httpClient, baseUrl, cred)

	// c, _ := Load("./res.json", mockReservationValiationPass)
	// res := c.ReservationsForWeek()
	loc, _ := time.LoadLocation("America/New_York")
	time, _ := time.ParseInLocation(iso8601Format, "2016-05-11T05:00:00.000", loc)
	res := Reservation{
		Day:             "Wednesday",
		StartTime:       rtime.UTCTime{time},
		EndTime:         "",
		Duration:        "60",
		ClubID:          "1010",
		ClubDescription: "PITTSBURGH-PENN AVE",
	}
	fmt.Println(res)
	laClient.MakeReservation(res)
	reservations, err := laClient.GetReservations()
	fmt.Println(reservations, err)
	fmt.Println("oh yea")
}
