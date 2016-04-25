package main

import (
	"github.com/stretchr/testify/assert"
	// "io/ioutil"
	"testing"
)

func TestNewReservation(t *testing.T) {
	r := NewReservation()
	assert.Equal(t, 14, r.Threshold, "Threshold should be two weeks")
	assert.Equal(t, r.Day, "Sunday")
}
