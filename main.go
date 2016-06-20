// Package main provides ...
package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const (
	laFitnessBaseUri = "https://publicapi.lafitness.com"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		_, ok := err.(*os.PathError)
		if ok {
			panic("You must have an .env file with LA_USERNAME and LA_PASSWORD defined")
		}
		panic(err)
	}
	laUsername := os.Getenv("LA_USERNAME")
	laPassword := os.Getenv("LA_PASSWORD")
	config, err := Load("./config.json", ValidateReservations)
	if err != nil {
		panic(err)
	}
	reservations := config.ReservationsForWeek()
	baseUrl, _ := url.Parse(laFitnessBaseUri)
	httpClient := http.DefaultClient
	cred := Credentials{
		Username: laUsername,
		Password: laPassword,
	}
	laClient := NewLaFitnessClient(httpClient, baseUrl, cred)

	for _, res := range reservations {
		laClient.MakeReservation(res)
	}
}
