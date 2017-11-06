package parsetime

import "time"

const (
	// Escape
	escapeCharacterOpen  = '['
	escapeCharacterClose = ']'

	// Times and timezones
	timezone = "MST"
	tzOffset = "Z07:00"
	amPM     = "PM"

	// Seconds
	microSecond = ".000000"
	second      = "05"

	// Minutes
	minute = "04"

	// Hours
	hour   = "15"
	hour12 = "3"

	// Days
	weekNameShort    = "Mon"
	weekName         = "Monday"
	dayOfMonthSingle = "_2"
	dayOfMonth       = "02"

	// Months
	monthName         = "Jan"
	monthNameFull     = "January"
	monthOfYearSingle = "1"
	monthOfYear       = "01"

	// Years
	year         = "2006"
	yearTwoDigit = "06"
)

var digitMap = map[rune]string{
	// Days
	'd': dayOfMonth,
	'D': weekNameShort,
	'l': weekName,
	'j': dayOfMonthSingle,
	// suffix missing
	// day of year missing

	// Months
	'F': monthNameFull,
	'M': monthName,
	'm': monthOfYear,
	'n': monthOfYearSingle,

	// Years
	'Y': year,
	'y': yearTwoDigit,

	// Time
	'a': amPM,
	'A': amPM,
	'h': hour12, // no support for 2 digit < 10
	'H': hour,
	'i': minute,
	's': second,
	'u': microSecond,

	// Timezone
	'e': timezone,
	'T': tzOffset,
}

// Format will format time t using the string format arguments, by overriding Go's native format library and using the
// inbuilt functionality to this library. An appropriate representation of time in string format will be returned.
func Format(t time.Time, format string) string {
	return t.Format(convertToNative(format))
}

// Parse will take value of string and a representation of the format of the value and return a time.Time structure
// and/or an error type in the event that the time is unable to be parsed. This function will wrap the native time.Parse
// func and return the output once the format is decoded to Go's standard library format.
func Parse(format, value string) (time.Time, error) {
	return time.Parse(convertToNative(format), value)
}

// convert the wrapped string representation of format into Go's native format for time, so it can be parsed by the
// native time libraries. Exclude any characters inside of the escape characters.
func convertToNative(format string) string {
	var real string
	var escaped bool

	for _, s := range format {
		if s == escapeCharacterOpen {
			escaped = true
		} else if s == escapeCharacterClose {
			escaped = false
		} else if r, ok := digitMap[s]; ok && !escaped {
			real += r
		} else {
			real += string(s)
		}
	}

	return real
}
