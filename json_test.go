// Package main provides ...
package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestjsonMarshal(t *testing.T) {
	json.Marshal(true)
	assert.Equal(t, true, true)
}
