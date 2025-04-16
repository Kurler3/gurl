package checks

import (
	"fmt"
	"net/http"

	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/utils"
)

var REQUIRED_FLAGS = map[string]struct{}{
	utils.UrlFlag:    {},
	utils.MethodFlag: {},
}

var DEFAULT_FLAG_VALUES = map[string]string{
	utils.MethodFlag:   http.MethodGet,
	utils.ProtocolFlag: utils.HTTPS,
}

func FinalFlagsCheck(g *gurl.Gurl) error {

	// For each flag in the default flag map => set it on the g instance.
	for flag, defaultValue := range DEFAULT_FLAG_VALUES {

		// Get the value for this flag.
		existingValue, err := g.GetFlag(flag)

		if err != nil {
			return err
		}

		// If no value, set the default.
		if existingValue == "" {
			g.SetFlag(flag, defaultValue)
		}

	}

	// Check if every required flag is defined.
	for flag := range REQUIRED_FLAGS {

		// Get value for flag
		value, err := g.GetFlag(flag)

		if err != nil {
			return err
		}

		// If no value => error
		if value == "" {
			return fmt.Errorf("no value defined for required flag %v", flag)
		}

	}

	return nil
}
