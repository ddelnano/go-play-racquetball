package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadReturnsReservationStructs(t *testing.T) {
	file := "./sample.json"
	config, err := Load(file)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(config.reservations))
}
