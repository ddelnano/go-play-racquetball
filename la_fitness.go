package racquetball

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	rtime "github.com/ddelnano/racquetball/time"
)

const (
	baseReservationsUrl = "/LAF_S4.5.2/Services/Private.svc/GetUpComingAppointments"
	makeReservationUrl  = "/LAF_S4.5.2/Services/Private.svc/CreateAmenityAppointment_2"
	iso8601Format       = "2006-01-02T15:04:00.000"
)

type LaFitnessRequest struct {
	Request LaRequestBody `json:"request"`
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

type MakeReservationResponse struct {
	CurrentServerTime string
	Message           string
	// ServerTimeZoneOffset
	Detail  string `json:"Detail,omitempty"`
	Success bool
	// Value   interface{}
}

type MakeReservationRequest struct {
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

// TODO: Should probably be deleted
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
	requestBody := LaFitnessRequest{
		Request: *NewLaRequestBody(nil),
	}
	// fmt.Println(requestBody, c.Credentials)
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

	if !reservations.Success {
		// fmt.Println(reservations)
		panic("We failed")
	}
	return transformReservations(reservations.Value.AmenityAppointments), err
}

func NewMakeReservationRequest(res Reservation) *LaFitnessRequest {
	makeResReq := MakeReservationRequest{
		ClubID:              "1010",
		ClubDescription:     "PITTSBURGH-PENN AVE",
		Duration:            res.Duration,
		AmenitiesApptTypeID: "1",
		AmenityID:           "0",
		StartDate:           res.StartTime.ISO8601(),
		StartDateUTC:        res.StartTime.ISO8601UTC(),
	}
	return &LaFitnessRequest{
		Request: *NewLaRequestBody(makeResReq),
	}
}

func (c *LaFitnessClient) MakeReservation(r Reservation) error {
	baseUrl := c.BaseUrl.String()
	url := fmt.Sprintf("%s%s", baseUrl, makeReservationUrl)
	requestBody := NewMakeReservationRequest(r)
	body, err := EncodeBody(requestBody)

	if err != nil {
		panic("Encoding json body failed")
	}

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Credentials.Username, c.Credentials.Password)
	res, err := c.Client.Do(req)
	defer res.Body.Close()

	// fmt.Println(requestBody.Value)
	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(data))

	if err != nil {
		panic(err.Error())
	}

	makeResResponse := MakeReservationResponse{}
	json.NewDecoder(res.Body).Decode(&makeResResponse)

	if !makeResResponse.Success {
		fmt.Println(makeResResponse)
		panic("Attempted to make reservation but received message " + makeResResponse.Message + ` and detail ` + makeResResponse.Detail)
	}

	return nil
}

// TODO: Needs tests
func transformReservations(r []amenityAppointment) []Reservation {
	res := make([]Reservation, 0)
	for _, appt := range r {
		startTime, _ := time.Parse(iso8601Format, appt.StartTime)
		endTime, _ := time.Parse(iso8601Format, appt.EndTime)
		reservation := Reservation{
			Day:       startTime.Weekday().String(),
			StartTime: rtime.UTCTime{Time: startTime},
			EndTime:   fmt.Sprintf("%d:%02d", endTime.Hour(), endTime.Minute()),
		}
		res = append(res, reservation)
	}
	return res
}

type LaRequestBody struct {
	Client laClient    `json:"Client"`
	Value  interface{} `json:"Value"`
}

type laClient struct {
	OSName string `json:"OSName"`
}

func NewLaRequestBody(v interface{}) *LaRequestBody {
	return &LaRequestBody{
		Client: laClient{
			OSName: "iPhone",
		},
		Value: v,
	}
}
