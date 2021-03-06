// Package main provides ...
package time

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const iso8601Regex = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z$`

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

func TestISO8601(t *testing.T) {
	now := time.Now()
	time := UTCTime{now}
	value := time.ISO8601()
	regex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}$`)

	if !regex.MatchString(value) {
		t.Errorf("Expected ISO 8601 format but received %s", value)
	}
}

func TestISO8601UTC(t *testing.T) {
	now := time.Now()
	time := UTCTime{now}
	value := time.ISO8601UTC()
	regex := regexp.MustCompile(iso8601Regex)

	if !regex.MatchString(value) {
		t.Errorf("Expected ISO 8601 format in UTC but received %s", value)
	}
}

func TestGetNextWeekday(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	n := time.Date(2016, 5, 16, 15, 30, 30, 0, loc)
	tim := UTCTime{n}

	days := []struct {
		Weekday time.Weekday
		Hour    int
	}{
		{
			Weekday: time.Sunday,
			Hour:    5,
		},
		{
			Weekday: time.Monday,
			Hour:    5,
		},
		{
			Weekday: time.Tuesday,
			Hour:    5,
		},
		{
			Weekday: time.Wednesday,
			Hour:    5,
		},
		{
			Weekday: time.Thursday,
			Hour:    5,
		},
		{
			Weekday: time.Friday,
			Hour:    5,
		},
		{
			Weekday: time.Saturday,
			Hour:    5,
		},
	}

	for _, v := range days {
		// Get the next weekday at 5:00 AM
		nextWeekday := tim.UpcomingWeekdayAtHour(v.Weekday, v.Hour)

		day := nextWeekday.Day()
		expectedDay := getDateOfUpcomingWeekday(v.Weekday)
		if day != getDateOfUpcomingWeekday(v.Weekday) {
			t.Errorf("Day should be %d, instead received %d", expectedDay, day)
		}

		hour := nextWeekday.Hour()
		if hour != v.Hour {
			t.Errorf("Hour should be %d, instead received %d", v.Hour, hour)
		}

		min := nextWeekday.Minute()
		if min != 0 {
			t.Errorf("Min should be 0, instead received %d", min)
		}

		secs := nextWeekday.Second()
		if secs != 0 {
			t.Errorf("Secs should be 0, instead received %d", secs)
		}
	}
}

func getDateOfUpcomingWeekday(weekday time.Weekday) int {
	switch weekday {
	case time.Sunday:
		return 22
	case time.Monday:
		return 23
	case time.Tuesday:
		return 17
	case time.Wednesday:
		return 18
	case time.Thursday:
		return 19
	case time.Friday:
		return 20
	case time.Saturday:
		return 21
	}
	return 0
}
