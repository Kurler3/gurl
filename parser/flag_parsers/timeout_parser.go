package flag_parsers

import (
	"errors"
	"strconv"
	"time"
)

func ParseTimeout(value string) (time.Duration, error) {

	// Parse into int
	milliseconds, err := strconv.Atoi(value)

	if err != nil {
		return time.Nanosecond, err
	}

	if milliseconds <= 0 {
		return time.Nanosecond, errors.New("specified timeout is invalid. Must be greater than 0 milliseconds")
	}

	return time.Duration(milliseconds) * time.Millisecond, nil

}
