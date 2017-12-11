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

func TestGetStartOfDay(test *testing.T) {
	t := time.Date(2017, 10, 25, 18, 36, 45, 56, time.UTC)
	expected := time.Date(2017, 10, 25, 0, 0, 0, 0, time.UTC)
	actual := parsetime.GetStartOfDay(t)
	assert.Equal(test, expected, actual)
}

func TestGetEndOfDay(test *testing.T) {
	t := time.Date(2017, 10, 25, 18, 36, 45, 56, time.UTC)
	expected := time.Date(2017, 10, 25, 23, 59, 59, 999999999, time.UTC)
	actual := parsetime.GetEndOfDay(t)
	assert.Equal(test, expected, actual)
	newExpected := time.Date(2017, 10, 26, 0, 0, 0, 0, time.UTC)
	newActual := actual.Add(time.Nanosecond)
	assert.Equal(test, newExpected, newActual)
}

func TestInTimeZone_UTC(test *testing.T) {
	t := time.Date(2017, 10, 25, 18, 36, 45, 56, time.UTC)
	expected := time.Date(2017, 10, 25, 18, 36, 45, 56, time.UTC)
	actual, tzErr := parsetime.InTimeZone(t, "UTC")
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_London(test *testing.T) {
	var londonName = "Europe/London"
	london, _ := time.LoadLocation(londonName)

	t := time.Date(2017, 12, 25, 18, 36, 45, 56, time.UTC)
	expected := time.Date(2017, 12, 25, 18, 36, 45, 56, london)
	actual, tzErr := parsetime.InTimeZone(t, londonName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_LondonBST(test *testing.T) {
	var londonName = "Europe/London"
	london, _ := time.LoadLocation(londonName)

	t := time.Date(2017, 6, 25, 18, 36, 45, 56, time.UTC)
	expected := time.Date(2017, 6, 25, 19, 36, 45, 56, london)
	actual, tzErr := parsetime.InTimeZone(t, londonName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_Paris(test *testing.T) {
	var parisName = "Europe/Paris"
	paris, _ := time.LoadLocation(parisName)

	t := time.Date(2017, 9, 25, 18, 36, 45, 56, time.UTC)
	expected := time.Date(2017, 9, 25, 20, 36, 45, 56, paris)
	actual, tzErr := parsetime.InTimeZone(t, parisName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_HongKong(test *testing.T) {
	var hongkongName = "Hongkong"
	hkt, _ := time.LoadLocation(hongkongName)

	t := time.Date(2017, 9, 25, 18, 36, 45, 56, time.UTC)
	expected := time.Date(2017, 9, 26, 02, 36, 45, 56, hkt)
	actual, tzErr := parsetime.InTimeZone(t, hongkongName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_Sydney(test *testing.T) {
	var sydneyName = "Australia/Sydney"
	sydney, e := time.LoadLocation(sydneyName)
	if e != nil {
		panic(e)
	}
	t := time.Date(2017, 11, 25, 11, 36, 45, 56, time.UTC)
	expected := time.Date(2017,11, 25, 22, 36, 45, 56, sydney)
	actual, tzErr := parsetime.InTimeZone(t, sydneyName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_Sydney_Daylight(test *testing.T) {
	var sydneyName = "Australia/Sydney"
	sydney, e := time.LoadLocation(sydneyName)
	if e != nil {
		panic(e)
	}
	t := time.Date(2017, 6, 25, 11, 36, 45, 56, time.UTC)
	expected := time.Date(2017,6, 25, 21, 36, 45, 56, sydney)
	actual, tzErr := parsetime.InTimeZone(t, sydneyName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_Santiago(test *testing.T) {
	var chileName = "Chile/Continental"
	chile, e := time.LoadLocation(chileName)
	if e != nil {
		panic(e)
	}
	t := time.Date(2017, 6, 25, 17, 36, 45, 56, time.UTC)
	expected := time.Date(2017,6, 25, 13, 36, 45, 56, chile)
	actual, tzErr := parsetime.InTimeZone(t, chileName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestInTimeZone_SantiagoSummer(test *testing.T) {
	var chileName = "Chile/Continental"
	chile, e := time.LoadLocation(chileName)
	if e != nil {
		panic(e)
	}
	t := time.Date(2017, 03, 25, 10, 36, 45, 56, time.UTC)
	expected := time.Date(2017,03, 25, 7, 36, 45, 56, chile)
	actual, tzErr := parsetime.InTimeZone(t, chileName)
	assert.Nil(test, tzErr)
	assert.Equal(test, expected, actual)
}

func TestParseWithTimeZone(test *testing.T) {
	southAfrica := "Africa/Johannesburg"
	saTZ, e := time.LoadLocation(southAfrica)
	if e != nil {
		panic(e)
	}

	actual, e := parsetime.ParseWithTimeZone("Y-m-d[T]H:i:s", "2017-11-24T08:39:15", southAfrica)
	expected := time.Date(2017,11, 24, 8, 39, 15, 0, saTZ)

	assert.Nil(test, e)
	assert.Equal(test, expected, actual)
}