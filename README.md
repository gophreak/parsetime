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
'i': Minute number (00 - 60)
's': Second number (00 - 60)
'u': Microseconds (up to 6 digits),

// Timezone
'e': Timezone name (UTC, GMT, etc)
'T': Timezzone offset (±07:00)
```

If you would like to escape reserved characters for use during the formatting, you can wrap it in `[]square brackets] to
preserve the individual characters. See some examples below:

## Examples

Assuming a time of 24th July 2017, at 09:35:42
### Dates

If you want to print out a date in the format of `Year-Month-Date`, then the following would do that:
```
fmt.Println(parsetime.Format(t, "Y-m-d"))
```
Will print
```
"2017-07-24"
```

### Times
If you wanted to print out the time, with `Hour:Minute:Second`, then use the following:
```
fmt.Println(parsetime.Format(t, "H:i:s"))
```
Will print
```
"09:35:42"
```

### Escaped characters

If you want to print out a date in the format of `Time: Year-Month-DateTHour:Minute:Second`, then you would need to
escape the reserved characters. In this instance, it would be simpler to escape in blocks of words:

```
fmt.Println(parsetime.Format(t, "[Time: ]Y-m-d[T]H:i:s"))
```
Will print
```
"Time: 2017-07-24T09:35:42"
```