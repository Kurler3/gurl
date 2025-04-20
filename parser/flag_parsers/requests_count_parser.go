package flag_parsers

import (
	"errors"
	"strconv"
)

func ParseRequestsCount(value string) (int, error) {

	// Parse into int
	count, err := strconv.Atoi(value)

	if err != nil {
		return 0, err
	}

	if count <= 0 {
		return 0, errors.New("requests count must be greater than 0")
	}

	return count, nil
}
