package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	Name string
}

func TestLoadReturnsReservationStructs(t *testing.T) {
	file := "./sample.json"
	config, err := Load(file)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(config.Reservations))
	assert.Equal(t, true, true)
}

func TestJsonmarshal(t *testing.T) {
	bytes := []byte(`{"name": "test"}`)
	var test Test
	err := json.Unmarshal(bytes, &test)
	assert.Equal(t, err, nil)
}

func TestReservationUnmarshal(t *testing.T) {
	bytes := []byte(`{ "day": "Wednesday", "time": 6, "threshold": 10}`)
	var reservation Reservation
	err := json.Unmarshal(bytes, &reservation)
	assert.Equal(t, err, nil)
}
