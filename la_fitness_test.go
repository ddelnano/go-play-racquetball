package racquetball

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	rtime "github.com/ddelnano/racquetball/time"
	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
)

func getReservationsResponse() []byte {
	return []byte(`
	{
	  "Message": "",
	  "CurrentServerTime": "04-29-16 21:44:29",
	  "ServerTimeZoneOffset": "-07:00:00",
	  "Success": true,
	  "UTCServerTime": "2016-04-30T04:44:29.8074915Z",
	  "Value": {
	    "AmenityAppointments": [
	    {
		"AmenititesAppointmentID": 7242424,
		"AmenityDescription": "1 RACQUETBALL COURT 1",
		"ClubDescription": "PITTSBURGH-PENN AVE",
		"ClubID": 1010,
		"EndTime": "2016-05-01T14:00:00.000",
		"Notes": "",
		"OtherCustomerNames": "",
		"StartTime": "2016-05-01T12:30:00.000"
	      },
	      {
		"AmenititesAppointmentID": 7242424,
		"AmenityDescription": "1 RACQUETBALL COURT 1",
		"ClubDescription": "PITTSBURGH-PENN AVE",
		"ClubID": 1010,
		"EndTime": "2016-05-01T14:00:00.000",
		"Notes": "",
		"OtherCustomerNames": "",
		"StartTime": "2016-05-01T12:30:00.000"
	      }
	    ],
	    "TrainingAppointments": []
	  }
	}
	`)
}

func TestNewReservation(t *testing.T) {
	r := NewReservation()
	assert.Equal(t, r.Day, "Sunday")
}

func TestClientCanGetReservations(t *testing.T) {
	// h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, client")
	// })

	// ts := httptest.NewServer(h)
	// defer ts.Close()

	// res, err := http.Get(ts.URL)
}

func TestNewLaFitnessClient(t *testing.T) {
	// New client has crendentials
	username := "ddelnano"
	password := "password"
	baseUrl, _ := url.Parse("base url")
	httpClient := http.DefaultClient
	cred := Credentials{Username: username, Password: password}

	client := NewLaFitnessClient(httpClient, baseUrl, cred)

	if client.Credentials.Username != username {
		t.Errorf("Credentials username must be set")
	}

	if client.Credentials.Password != password {
		t.Errorf("Credentials password must be set")
	}
}

// TODO: These need tests
func testRequestMethod(t *testing.T, r *http.Request, method string) {
	if r.Method != "POST" {
		t.Fail()
	}

	if r.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type must be json")
	}
}

// TODO: These need tests
func testBasicAuthSet(t *testing.T, r *http.Request) {
	if username, pass, _ := r.BasicAuth(); username == "" && pass == "" {
		t.Errorf("Username and password should be set")
	}
}

func TestGetReservations(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(baseReservationsUrl, func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "POST")
		testBasicAuthSet(t, r)

		var v LaFitnessRequest
		err := json.NewDecoder(r.Body).Decode(&v)
		defer r.Body.Close()

		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		expected := LaFitnessRequest{
			Request: *NewLaRequestBody(nil),
		}
		fmt.Printf("Testing: %+v\n\n\n", expected)
		if !reflect.DeepEqual(expected, v) {
			t.Errorf("Request body = %#v, expected %#v", v, expected)
		}

		w.Write(getReservationsResponse())
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	client := http.DefaultClient
	baseUrl, _ := url.Parse(s.URL)
	cred := Credentials{Username: "ddelnano", Password: "password"}
	laClient := NewLaFitnessClient(client, baseUrl, cred)
	reservations, ok := laClient.GetReservations()

	// TODO: Add DeepEqual here with expected struct
	assert.Equal(t, 2, len(reservations))
	assert.Equal(t, "Sunday", reservations[0].Day)
	// assert.Equal(t, "12:30", reservations[0].StartTime)
	assert.IsType(t, rtime.UTCTime{}, reservations[0].StartTime)
	assert.Nil(t, ok)
}

func TestMakeNewReservationRequest(t *testing.T) {
	res := Reservation{
		Duration: "60",
		StartTime: rtime.UTCTime{
			time.Now(),
		},
	}
	// expectedRequest := LaFitnessRequest{
	// 	Value: MakeReservationRequest{
	// 		ClubID:              "1010",
	// 		ClubDescription:     "PITTSBURGH-PENN AVE",
	// 		Duration:            duration,
	// 		AmenitiesApptTypeID: "1",
	// 		AmenityID:           "0",
	// 		StartDate:           "2016-05-10T09:00:00.000",
	// 		StartDateUTC:        "2016-05-10T13:00:00.000Z",
	// 	},
	// }
	// var startTime rtime.UTCTime
	// data := []byte(`{"Time":"6:30"}`)
	// json.Unmarshal(data, startTime)
	// res := Reservation{
	// 	Day:       "Sunday",
	// 	StartTime: startTime,
	// 	Duration:  duration,
	// }
	resRequest := NewMakeReservationRequest(res)
	fmt.Println(resRequest)
	// makeResRequest := resRequest.Request.Value.(MakeReservationRequest)
	// assert.Equal(t, res.Duration, makeResRequest.Duration, "Duration should make Reservation's")
	// assert.Equal(t, res.StartTime.ISO8601(), makeResRequest.StartDate)
	// assert.Equal(t, res.StartTime.ISO8601UTC(), makeResRequest.StartDateUTC)
}

// TODO: Handle errors and if success key in response is false
func TestMakeReservation(t *testing.T) {
	res := new(Reservation)
	mux := http.NewServeMux()
	mux.HandleFunc(makeReservationUrl, func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "POST")
		testBasicAuthSet(t, r)

		var v LaFitnessRequest
		json.NewDecoder(r.Body).Decode(&v)
		defer r.Body.Close()

		expected := *NewMakeReservationRequest(*res)
		if !reflect.DeepEqual(expected.Request.Client, v.Request.Client) {
			t.Errorf("Request body = %#v, expected %#v", v, expected)
		}

		if !reflect.DeepEqual(structs.Map(expected.Request.Value), v.Request.Value) {
			t.Errorf("Request.Value = %#v, expected %#v", v.Request.Value, expected.Request.Value)
		}
		data := []byte(
			`{
			  "CurrentServerTime": "05-08-16 15:42:58",
			  "Message": "Message",
			  "ServerTimeZoneOffset": "-07:00:00",
			  "Success": true,
			  "UTCServerTime": "2016-05-08T22:42:58.3208260Z"
			 }
			`)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	client := http.DefaultClient
	baseUrl, _ := url.Parse(s.URL)
	cred := Credentials{Username: "ddelnano", Password: "password"}
	laClient := NewLaFitnessClient(client, baseUrl, cred)
	laClient.MakeReservation(*res)
}

func TestMakeReservationWhenLaRespondsWithSuccessFalse(t *testing.T) {
	res := new(Reservation)
	mux := http.NewServeMux()
	mux.HandleFunc(makeReservationUrl, func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "POST")
		testBasicAuthSet(t, r)

		var v LaFitnessRequest
		json.NewDecoder(r.Body).Decode(&v)
		defer r.Body.Close()

		data := []byte(
			`{
			  "CurrentServerTime": "05-08-16 15:42:58",
			  "Message": "Message",
			  "ServerTimeZoneOffset": "-07:00:00",
			  "Success": false,
			  "UTCServerTime": "2016-05-08T22:42:58.3208260Z"
			 }
			`)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	client := http.DefaultClient
	baseUrl, _ := url.Parse(s.URL)
	cred := Credentials{Username: "ddelnano", Password: "password"}
	laClient := NewLaFitnessClient(client, baseUrl, cred)

	assert.Panics(t, func() {
		laClient.MakeReservation(*res)
	})
}

func TestLaFitnessRequest(t *testing.T) {
	requestBody := NewLaRequestBody(nil)
	assert.Equal(t, "iPhone", requestBody.Client.OSName)
}

func Test_transformReservations(t *testing.T) {
	appts := []amenityAppointment{
		amenityAppointment{
			StartTime: "2016-05-10T14:04:00.000",
			EndTime:   "2016-05-10T15:04:00.000",
		},
		amenityAppointment{
			StartTime: "2016-05-01T12:30:00.000",
			EndTime:   "2016-05-01T13:30:00.000",
		},
	}

	res := transformReservations(appts)

	assert.Equal(t, 2, len(res))
	// assert.Equal(t, "14:04", res[0].StartTime)
	// assert.Equal(t, "12:30", res[1].StartTime)
	assert.IsType(t, rtime.UTCTime{}, res[0].StartTime)
	assert.IsType(t, rtime.UTCTime{}, res[1].StartTime)
	assert.Equal(t, "15:04", res[0].EndTime)
	assert.Equal(t, "13:30", res[1].EndTime)
}
