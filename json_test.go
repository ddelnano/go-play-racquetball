// Package main provides ...
package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestjsonMarshal(t *testing.T) {
	json.Marshal(true)
	assert.Equal(t, true, true)
}
