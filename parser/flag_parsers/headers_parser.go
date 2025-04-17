package flag_parsers

import (
	"errors"
	"fmt"
	"strings"
)

func ParseHeaders(value string) (any, error) {

	if value == "" {
		return nil, errors.New("headers are empty")
	}

	// Example: "Content-Type: application/json; Whatever: whatever"
	headersMap := map[string]string{}

	headerStrs := strings.Split(value, ";")

	if len(headerStrs) == 0 {
		return nil, errors.New("header lines need to be separated by semicolon. Example: \"Content-Type: application/json; John: Doe\" ")
	}

	for _, headerStr := range headerStrs {

		headerValues := strings.Split(headerStr, ":")

		if len(headerValues) < 2 {
			return nil, errors.New("header lines need to have a key and a value separated by ':'. Example: Content-Type: application/json")
		}

		key, value := headerValues[0], headerValues[1]

		// If key already exists, means its duplicated
		if _, ok := headersMap[key]; ok {
			return nil, fmt.Errorf("key \"%v\" is duplicated. Please define only header line for each key", key)
		}

		headersMap[key] = value

	}

	return headersMap, nil

}
