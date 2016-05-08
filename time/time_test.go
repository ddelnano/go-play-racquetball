// Package main provides ...
package time

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type test struct {
	Time UTCTime
}

func getUnmarshalTimeForHour(hour string) string {
	expectedTime := time.Now().Format(format)
	return fmt.Sprintf("%s %s -0400 EDT", expectedTime, hour)
}

func TestUnmarshalUTCTime(t *testing.T) {

	tests := []struct {
		data     []byte
		expected string
	}{
		{
			data:     []byte(`{"Time": "6:00"}`),
			expected: getUnmarshalTimeForHour("06:00:00"),
		},
		{
			data:     []byte(`{"Time": "6:30"}`),
			expected: getUnmarshalTimeForHour("06:30:00"),
		},
	}

	for _, v := range tests {
		test := test{}

		err := json.Unmarshal(v.data, &test)
		str := test.Time.String()

		if err != nil {
			t.Errorf("UnmarshalJSON failed for %s with error", v.data, err)
		}

		if strings.Compare(v.expected, str) != 0 {
			t.Fatalf("String value of time %s not equal to %s", str, v.expected)
		}
	}
}

func TestUnmarshalUTCTimeFailsIfTimeDoesNotSpecifyHourAndMinutes(t *testing.T) {
	test := test{}

	assert.Panics(t, func() {
		json.Unmarshal([]byte(`{"Time": "6"}`), &test)
	}, "JSON value of `6` should not be parsable")
}

func TestMarshalUTC(t *testing.T) {
	test := test{
		UTCTime{
			time.Now(),
		},
	}
	data, err := json.Marshal(&test)

	if err != nil {
		t.Fatalf("json.Marshal failed with error: %s", err.Error())
	}

	iso8601Value := fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02d",
		test.Time.Year(),
		test.Time.Month(),
		test.Time.Day(),
		test.Time.Hour(),
		test.Time.Minute(),
		test.Time.Second(),
	)
	if strings.Compare(`{"Time":"`+iso8601Value+`"}`, string(data)) != 0 {
		t.Fatalf("String value from json.Marshal incorrect: %s", string(data))
	}
}
