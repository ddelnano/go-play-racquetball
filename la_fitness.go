package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseReservationsUrl = "/LAF_S4.5.2/Services/Private.svc/GetUpComingAppointments"
const iso8601Format = "2006-01-02T15:04:00.000"

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
	CurrentServerTime time.Time
	Message           string
	// ServerTimeZoneOffset
	Success       bool
	UTCServerTime time.Time
	Value         struct {
		AmenityAppointments  []amenityAppointment
		TrainingAppointments interface{}
	}
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

func NewLaFitnessClient(client *http.Client, baseUrl *url.URL, cred Credentials) *LaFitnessClient {
	return &LaFitnessClient{Client: client, BaseUrl: baseUrl, Credentials: cred}
}

func (c *LaFitnessClient) GetReservations() ([]Reservation, error) {
	baseUrl := c.BaseUrl.String()
	url := fmt.Sprintf("%s%s", baseUrl, baseReservationsUrl)
	requestBody := NewLaRequestBody()
	body, err := EncodeBody(requestBody)

	if err != nil {
		panic("Encoding json body failed")
	}

	res, err := c.Client.Post(url, "application/json", body)
	var reservations getReservationResponse
	err = json.NewDecoder(res.Body).Decode(&reservations)
	return transformReservations(reservations.Value.AmenityAppointments), err
}

func (*LaFitnessClient) MakeReservation(*Reservation) ([]Reservation, error) {
	return nil, nil
}

// TODO: Needs tests
func transformReservations(r []amenityAppointment) []Reservation {
	res := make([]Reservation, 0)
	for _, appt := range r {
		// fmt.Printf("%#v", appt)
		startTime, _ := time.Parse(iso8601Format, appt.StartTime)
		// endTime, _ := time.Parse(iso8601Format, appt.EndTime)
		time := strconv.Itoa(startTime.Hour())
		fmt.Println(startTime.Hour())
		reservation := Reservation{
			Day:  startTime.Weekday().String(),
			Time: time,
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
		return nil, err
	}
	return buf, nil
}

func ViewReservations() (bool, error) {
	data := []byte(`{
    "request": {
	"Client": {
	    "OSName": "iPhone"
	}
    }
}`)
	client := &http.Client{}
	// resp, err := client.Post("https://publicapi.lafitness.com/LAF_S4.5.2/Services/Private.svc/GetUpComingAppointments", "application/json", bytes.NewBuffer(data))
	req, err := http.NewRequest("POST", "https://publicapi.lafitness.com/LAF_S4.5.2/Services/Private.svc/GetUpComingAppointments", bytes.NewBuffer(data))

	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", "Basic ZGRlbG5hbm86UGl0dHRpZ2VyczJA")
	req.SetBasicAuth("ddelnano", "Pitttigers2@")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(req.Body)
	fmt.Println(resp)
	contents, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))
	return true, err
}
