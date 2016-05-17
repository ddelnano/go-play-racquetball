package racquetball

import (
	"flag"
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
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// Get reservation start time to be next friday at 5 AM
	// this should be one of the most available times during the
	// week.
	startTime := rtime.UTCTime{
		time.Now(),
	}
	res := Reservation{
		Day:             "Friday",
		StartTime:       startTime.UpcomingWeekdayAtHour(time.Friday, 5),
		EndTime:         "",
		Duration:        "60",
		ClubID:          "1010",
		ClubDescription: "PITTSBURGH-PENN AVE",
	}
	id := client.MakeReservation(res)

	if id <= 0 {
		t.Errorf("AmenititesAppointmentID %d must be valid ID", id)
	}

	client.DeleteReservation(id)
}

func TestIntegrationForGetReservations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	res, err := client.GetReservations()

	if err != nil {
		t.Errorf("Trying to get reservations failed with reservations: %#v and error: %#v", res, err)
	}
}
