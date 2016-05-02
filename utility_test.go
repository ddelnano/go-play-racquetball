package racquetball

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodyBodyFailsIfBodyNil(t *testing.T) {
	assert.Panics(t, func() {
		EncodeBody(nil)
	})
}

func TestEncodeBodyEncodesToJSONBuffer(t *testing.T) {
	type testBody struct {
		Test    string `json:"test"`
		Another string `json:"another"`
	}
	body := testBody{
		Test:    "testing",
		Another: "string",
	}
	by, ok := EncodeBody(body)
	data := []byte(`{"test":"testing","another":"string"}`)
	expectedBuffer := bytes.NewBuffer(data)
	assert.Equal(t, strings.TrimSpace(expectedBuffer.String()), strings.TrimSpace(by.String()))
	assert.Nil(t, ok)
}
