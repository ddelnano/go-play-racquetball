package main

import (
	"github.com/stretchr/testify/assert"
	// "io/ioutil"
	"testing"
)

func TestNewReservation(t *testing.T) {
	r := NewReservation()
	assert.Equal(t, r.Day, "Sunday")
}
