#Parsetime

The parsetime package is designed as a helper to get around the obscurity of Go's time parsing. Coming from a PHP
background I find it easier to implement the standards laid out by PHP so I have created this helper to help when parsing
or formatting times.

## Getting started

Download the package using glide by adding the following into your glide.yaml file

```
import:
- package: github.com:gophreak/parsetime
  repo: git@github.com:gophreak/parsetime.git
  vcs: git
```

Then run `glide up -v` to install the package.

## Using the package

There are two exported functions available for use. `Parse(format, value string)` and `Format(t time.Time, format string)`.

The Parse function will parse a date format and a date value as arguments and return a time in the format of `time.Time`
and an `error` in the event that there was an error parsing the date.

The Format function will format a `time.Time` structure in the format given, and return a string representative of the 
time in the format requested.

Samples of date can be found below:

```
// Days
'd': Day of month (01, 02, 03, ... 30, 31, etc)
'j': Day of month with single digit (1, 2, 3, ... 10, 11, etc)
'D': Day of week - shortname (Mon, Tue, Wed, etc)
'l': Day of week - name (Monday, Tuesday, Wednesday, etc)

// Months
'F': Month - Full name (January, February, March, etc)
'M': Month - Short name (Jan, Feb, Mar, etc)
'm': Month number (01, 02, 03, ... 11, 12)
'n': Month number - Single digit (1, 2, 3, ... 11, 12)

// Years
'Y': Year (2006, 2015, 2016, etc) 
'y': Year in two digits (06, 09, 11, 17, etc)

// Time
'a': Morning or evening in AM/PM format
'A': Morning or evening in AM/PM format
'h': Hour number in 12 hour format (1 - 12)
'H': Hour number in 24 hour format (0 - 23)
'i': Minute number (0 - 60)
's': Second number (0 - 60)
'u': Microseconds (up to 6 digits),

// Timezone
'e': Timezone name (UTC, GMT, etc)
'T': Timezzone offset (±07:00)
```