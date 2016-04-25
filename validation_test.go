// Package main provides ...
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValiationForValidReservationFile(t *testing.T) {
	result, err := ValidateReservations("./sample.json")
	assert.Equal(t, true, true)
	assert.Equal(t, true, result.Valid())
	assert.Equal(t, 0, len(result.Errors()))
	assert.Equal(t, nil, err)
}
