package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseReservationsUrl = "/LAF_S4.5.2/Services/Private.svc/GetUpComingAppointments"
	makeReservationUrl  = "/LAF_S4.5.2/Services/Private.svc/CreateAmenityAppointment_2"
	iso8601Format       = "2006-01-02T15:04:00.000"
)

type LaFitnessRequest struct {
	Value   interface{}
	Request LaRequestBody
}

type amenityAppointment struct {
	AmenitiesAppointmentID int
	AmenityDescription     string
	ClubDescription        string
	ClubID                 int
	EndTime                string
	Notes                  string
	OtherCustomerNames     string
	StartTime              string
}

type getReservationResponse struct {
	CurrentServerTime string
	Message           string
	// ServerTimeZoneOffset
	Success       bool
	UTCServerTime time.Time
	Value         struct {
		AmenityAppointments  []amenityAppointment
		TrainingAppointments interface{}
	}
}

type makeReservationRequest struct {
	ClubID              string
	ClubDescription     string
	Duration            string
	AmenitiesApptTypeID string
	AmenityID           string
	StartDate           string
	StartDateUTC        string
}

type Credentials struct {
	Username string
	Password string
}

func NewReservation() *Reservation {
	return &Reservation{
		Day: "Sunday",
	}
}

type Client struct {
	Gym     GymClient
	baseUrl *url.URL
}

type LaFitnessClient struct {
	Client      *http.Client
	BaseUrl     *url.URL
	Credentials Credentials
}

// TODO: Need to add http basic authentication so client is usable after creation
func NewLaFitnessClient(client *http.Client, baseUrl *url.URL, cred Credentials) *LaFitnessClient {
	return &LaFitnessClient{
		Client:      client,
		BaseUrl:     baseUrl,
		Credentials: cred,
	}
}

func (c *LaFitnessClient) GetReservations() ([]Reservation, error) {
	baseUrl := c.BaseUrl.String()
	url := fmt.Sprintf("%s%s", baseUrl, baseReservationsUrl)
	requestBody := NewLaRequestBody()
	body, err := EncodeBody(requestBody)

	if err != nil {
		panic("Encoding json body failed")
	}

	// TODO: This should be extracted for reuse
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Credentials.Username, c.Credentials.Password)
	res, err := c.Client.Do(req)

	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	var reservations getReservationResponse
	err = json.NewDecoder(res.Body).Decode(&reservations)
	return transformReservations(reservations.Value.AmenityAppointments), err
}

func NewMakeReservationRequest() *LaFitnessRequest {
	return &LaFitnessRequest{
		Value: makeReservationRequest{
			ClubID:              "1010",
			ClubDescription:     "PITTSBURGH-PENN AVE",
			Duration:            "60",
			AmenitiesApptTypeID: "1",
			AmenityID:           "0",
			StartDate:           "2016-05-10T09:00:00.000",
			StartDateUTC:        "2016-05-10T13:00:00.000Z",
		},
		Request: *NewLaRequestBody(),
	}
}

func (c *LaFitnessClient) MakeReservation(*Reservation) ([]Reservation, error) {
	baseUrl := c.BaseUrl.String()
	url := fmt.Sprintf("%s%s", baseUrl, makeReservationUrl)
	requestBody := NewMakeReservationRequest()
	body, err := EncodeBody(requestBody)

	if err != nil {
		panic("Encoding json body failed")
	}

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Credentials.Username, c.Credentials.Password)
	res, err := c.Client.Do(req)
	defer res.Body.Close()

	if err != nil {
		panic(err.Error())
	}
	return nil, nil
}

// TODO: Needs tests
func transformReservations(r []amenityAppointment) []Reservation {
	res := make([]Reservation, 0)
	for _, appt := range r {
		startTime, _ := time.Parse(iso8601Format, appt.StartTime)
		endTime, _ := time.Parse(iso8601Format, appt.EndTime)
		// endTime, _ := time.Parse(iso8601Format, appt.EndTime)
		reservation := Reservation{
			Day:       startTime.Weekday().String(),
			StartTime: fmt.Sprintf("%d:%02d", startTime.Hour(), startTime.Minute()),
			EndTime:   fmt.Sprintf("%d:%02d", endTime.Hour(), endTime.Minute()),
		}
		res = append(res, reservation)
	}
	return res
}

type LaRequestBody struct {
	Request laRequest `json:"request"`
}

type laRequest struct {
	Client laClient `json:"Client"`
}

type laClient struct {
	OSName string `json:"OSName"`
}

func NewLaRequestBody() *LaRequestBody {
	return &LaRequestBody{
		Request: laRequest{
			Client: laClient{
				OSName: "iPhone",
			},
		},
	}
}

// TODO: Utility function should be moved out into another package
func EncodeBody(body interface{}) (*bytes.Buffer, error) {
	if body == nil {
		panic("Body argument should not be nil")
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buf, nil
}
