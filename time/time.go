// Package main provides ...
package time

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006
const format = "2006-01-02"
const formatWithHours = "2006-01-02 15:04"
const ctLayout = "2006/01/02|15:04:05"
const iso8601Format = "2006-01-02T15:04:05.000"
const Format = "2006-01-02T15:04:05"
const jsonFormat = `"` + Format + `"`

type UTCTime struct {
	time.Time
}

func (t *UTCTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(jsonFormat)), nil
}

func (t *UTCTime) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	now := time.Now()
	ymd := now.Format(format)

	if !bytes.Contains(b, []byte(`:`)) {
		panic("json is in invalid format")
	}
	times := bytes.Split(b, []byte(`:`))
	hours, _ := strconv.ParseInt(string(times[0]), 10, 0)
	test := fmt.Sprintf("%sT%02d:%s:00.000", ymd, hours, string(times[1]))

	loc, _ := time.LoadLocation("America/New_York")
	t.Time, _ = time.ParseInLocation(Format, test, loc)
	return nil
}

func (t *UTCTime) ISO8601() string {
	return t.Format(iso8601Format)
}

func (t *UTCTime) ISO8601UTC() string {
	// TODO: Location line has no test asssertion
	loc, _ := time.LoadLocation("UTC")
	return t.In(loc).Format(iso8601Format) + `Z`
}
