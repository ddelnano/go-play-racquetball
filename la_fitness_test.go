package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

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

		var v LaRequestBody
		err := json.NewDecoder(r.Body).Decode(&v)
		defer r.Body.Close()

		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		expected := NewLaRequestBody()
		if !reflect.DeepEqual(*expected, v) {
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
	assert.Equal(t, "12:30", reservations[0].StartTime)
	assert.Nil(t, ok)
}

// TODO: Handle errors and if success key in response is false
func TestMakeReservation(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(makeReservationUrl, func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "POST")
		testBasicAuthSet(t, r)

		var v LaFitnessRequest
		json.NewDecoder(r.Body).Decode(&v)
		defer r.Body.Close()

		expected := *NewMakeReservationRequest()
		if !reflect.DeepEqual(expected.Request, v.Request) {
			t.Errorf("Request body = %#v, expected %#v", v, expected)
		}

		if !reflect.DeepEqual(structs.Map(expected.Value), v.Value) {
			t.Errorf("Request body = %#v, expected %#v", v.Value, expected.Value)
		}
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	client := http.DefaultClient
	baseUrl, _ := url.Parse(s.URL)
	cred := Credentials{Username: "ddelnano", Password: "password"}
	laClient := NewLaFitnessClient(client, baseUrl, cred)
	laClient.MakeReservation(nil)
}

func TestLaFitnessRequest(t *testing.T) {
	requestBody := NewLaRequestBody()
	assert.Equal(t, "iPhone", requestBody.Request.Client.OSName)
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
	assert.Equal(t, "14:04", res[0].StartTime)
	assert.Equal(t, "12:30", res[1].StartTime)
	assert.Equal(t, "15:04", res[0].EndTime)
	assert.Equal(t, "13:30", res[1].EndTime)
}
