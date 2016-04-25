// Package main provides ...
package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

const (
	RESERVATION_JSON_SCHEMA string = "./reservation.json"
)

func ValidateReservations(filepath string) (*gojsonschema.Result, error) {
	loaderPath := fmt.Sprintf("file://%s", RESERVATION_JSON_SCHEMA)
	documentPath := fmt.Sprintf("file://%s", filepath)

	schemaLoader := gojsonschema.NewReferenceLoader(loaderPath)
	documentLoader := gojsonschema.NewReferenceLoader(documentPath)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	return result, err
}
