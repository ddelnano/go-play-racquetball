// Package main provides ...
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	laUsername := os.Getenv("LA_USERNAME")
	laPassword := os.Getenv("LA_PASSWORD")
	fmt.Println("Awe yea go", laUsername, laPassword)
	// cred := Credentials{Username: laUsername, Password: laPassword}
	// baseUrl := "https://publicapi.lafitness.com"
	// laUrl, _ := url.Parse(baseUrl)
	// client := http.DefaultClient
	defer func() {
		// if r := recover(); r != nil {
		// 	fmt.Println("Recovered in f", r)
		// }
	}()
	// baseUrl, _ := url.Parse("https://publicapi.lafitness.com")
	// httpClient := http.DefaultClient
	// cred := Credentials{Username: laUsername, Password: laPassword}
	// laClient := NewLaFitnessClient(httpClient, baseUrl, cred)
	// reservations, err := laClient.GetReservations()
	// fmt.Println(reservations, err)
	// fmt.Println("oh yea")
}
