package parsetime_test

import (
	"testing"
	"time"

	"github.com/gophreak/parsetime"
	"github.com/stretchr/testify/assert"
)

const (
	timestampTest   = 1500885555
	timeRFC3339Test = "2017-07-24T09:39:15+01:00"
)

func TestFormat_WithYearMonthDay(test *testing.T) {
	t := time.Unix(timestampTest, 0)

	formatted := parsetime.Format(t, "Y-m-d")

	assert.Equal(test, "2017-07-24", formatted)
}

func TestFormat_WithHourMinuteSecond(test *testing.T) {
	t := time.Unix(timestampTest, 0)

	formatted := parsetime.Format(t, "H:i:s")

	assert.Equal(test, "09:39:15", formatted)
}

func TestFormat_WithTimezone(test *testing.T) {
	t := time.Unix(timestampTest, 0)

	formatted := parsetime.Format(t, "e")

	assert.Equal(test, "BST", formatted)
}

func TestFormat_WithTimezoneOffset(test *testing.T) {
	t := time.Unix(timestampTest, 0)

	formatted := parsetime.Format(t, "T")

	assert.Equal(test, "+01:00", formatted)
}

func TestFormat_WithEscapedCharacters(test *testing.T) {
	t := time.Unix(timestampTest, 0)

	formatted := parsetime.Format(t, "[Date]: Y-m-d")

	assert.Equal(test, "Date: 2017-07-24", formatted)
}

func TestFormat_WithRFC3339(test *testing.T) {
	t := time.Unix(timestampTest, 0)

	formatted := parsetime.Format(t, "Y-m-d[T]H:i:sT")

	assert.Equal(test, timeRFC3339Test, formatted)
}

func TestParse_WithRFC3339(test *testing.T) {
	t, e := parsetime.Parse("Y-m-d[T]H:i:sT", timeRFC3339Test)
	assert.NoError(test, e)
	assert.Equal(test, int64(timestampTest), t.Unix())
}

func TestParse_WithFullDateTime(test *testing.T) {
	t, e := parsetime.Parse("Y-m-d[T]H:i:s", "2017-11-24T08:39:15")
	assert.NoError(test, e)
	assert.Equal(test, int64(1511512755), t.Unix())
}

func TestParse_WithRFC3339NegativeTimezone(test *testing.T) {
	t, e := parsetime.Parse("Y-m-d[T]H:i:sT", "2017-11-24T08:39:15-05:00")
	assert.NoError(test, e)
	assert.Equal(test, int64(1511530755), t.Unix())
}
