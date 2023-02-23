//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Parser interface {
	Parse(string) error
}

type Time struct {
	Original string
	Hour     int
	Minute   int
	Second   int
}

func (t *Time) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
}

func (t *Time) Parse() error {
	if t.Original == "" {
		return errors.New("empty string")
	}
	parts := strings.Split(t.Original, ":")
	if len(parts) != 3 {
		return errors.New("invalid format")
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return err
	}
	if hours < 0 || minutes < 0 || seconds < 0 {
		return errors.New("invalid time")
	} else if hours > 23 || minutes > 59 || seconds > 59 {
		return errors.New("invalid time")
	}
	t.Hour = hours
	t.Minute = minutes
	t.Second = seconds
	return nil
}

func (t Time) New() Time {
	return Time{
		Original: t.Original,
		Hour:     t.Hour,
		Minute:   t.Minute,
		Second:   t.Second,
	}
}
