package flag_parsers

import (
	"fmt"

	"github.com/Kurler3/gurl/utils"
)

// Parser of the method flag.
func ParseMethod(method string) (string, error) {

	_, ok := utils.AVAILABLE_METHODS[method]

	if !ok {
		return "", fmt.Errorf("method %v not available", method)
	}

	return method, nil
}
