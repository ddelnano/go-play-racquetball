package racquetball

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	rtime "github.com/ddelnano/racquetball/time"
	"github.com/joho/godotenv"
)

var laUsername, laPassword string
var cred Credentials
var client *LaFitnessClient

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	laUsername = os.Getenv("LA_USERNAME")
	laPassword = os.Getenv("LA_PASSWORD")

	cred = Credentials{Username: laUsername, Password: laPassword}
	baseUrl, _ := url.Parse("https://publicapi.lafitness.com")
	httpClient := http.DefaultClient
	client = NewLaFitnessClient(httpClient, baseUrl, cred)

	flag.Parse()
	os.Exit(m.Run())
}

func TestMakingReservations(t *testing.T) {
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
	fmt.Println(client, res)
	// client.MakeReservation(res)
	// reservations, err := client.GetReservations()
	// fmt.Println(reservations, err)
	// fmt.Println("oh yea")
}

func TestIntegrationForGetReservations(t *testing.T) {
	res, err := client.GetReservations()

	if err != nil {
		t.Errorf("Trying to get reservations failed with reservations: %#v and error: %#v", res, err)
	}
}
