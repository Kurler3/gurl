package checks

import (
	"fmt"

	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/utils"
)

func FinalFlagsCheck(g *gurl.Gurl) error {

	// For each flag in the default flag map => set it on the g instance.
	for flag, defaultValue := range utils.DEFAULT_FLAG_VALUES {

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
	for flag := range utils.REQUIRED_FLAGS {

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

	//TODO Depending on the method, the body might not be available, or other flags might also not be available.

	//TODO - Any warnings for any flags, like SkipTLSVerification

	return nil
}
