package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewReservation(t *testing.T) {
	today := time.Now().Weekday()
	r := NewReservation()
	assert.Equal(t, 14, r.threshold, "Threshold should be two weeks")
	assert.Equal(t, r.day, today)
}

func TestUnmarchalJSON(t *testing.T) {

	// assert.Equal(t, r.day, time.
}
