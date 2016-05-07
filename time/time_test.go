// Package main provides ...
package time

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

type test struct {
	Time UTCTime
}

func TestUnmarshalUTCTime(t *testing.T) {
	var data = `
	{"Time": "6:00"}
`
	test := test{}

	err := json.Unmarshal([]byte(data), &test)
	str := test.Time.String()

	if err != nil {
		t.Fatalf("json.Marshal failed with error: %s", err.Error())
	}

	expectedTime := time.Now().Format(format)
	expected := expectedTime + ` 06:00:00 -0400 EDT`
	if strings.Compare(expected, str) != 0 {
		t.Fatalf("String value of time %s not equal to %s", str, expected)
	}
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
