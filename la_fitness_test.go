package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getReservationsResponse() []byte {
	//"CurrentServerTime": "04-29-16 21:44:29",
	return []byte(`
	{
	  "Message": "",
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

func TestGetReservations(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(baseReservationsUrl, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fail()
		}
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
	assert.Equal(t, "12", reservations[0].Time)
	assert.Nil(t, ok)
}

func TestLaFitnessRequest(t *testing.T) {
	requestBody := NewLaRequestBody()
	assert.Equal(t, "iPhone", requestBody.Request.Client.OSName)
}

func TestEncodyBodyFailsIfBodyNil(t *testing.T) {
	assert.Panics(t, func() {
		EncodeBody(nil)
	})
}

func TestEncodeBodyEncodesToJSONBuffer(t *testing.T) {
	type testBody struct {
		Test    string `json:"test"`
		Another string `json:"another"`
	}
	body := testBody{
		Test:    "testing",
		Another: "string",
	}
	by, ok := EncodeBody(body)
	data := []byte(`{"test":"testing","another":"string"}`)
	expectedBuffer := bytes.NewBuffer(data)
	assert.Equal(t, strings.TrimSpace(expectedBuffer.String()), strings.TrimSpace(by.String()))
	assert.Nil(t, ok)
}
