// Package main provides ...
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	laUsername := os.Getenv("LA_USERNAME")
	laPassword := os.Getenv("LA_PASSWORD")
	fmt.Println("Awe yea go", laUsername, laPassword)
	ViewReservations()
}
