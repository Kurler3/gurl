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
		if utils.IsEmpty(existingValue) {
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
		if utils.IsEmpty(value) {
			return fmt.Errorf("no value defined for required flag %v", flag)
		}

	}

	// Any warnings for any flags, like SkipTLSVerification
	for _, flagWarning := range utils.FLAG_WARNINGS {

		// Get the flag value.
		flagValue, err := g.GetFlag(flagWarning.Flag)

		if err != nil {
			return err
		}

		// Compare to the conditional value and display the warning if it matches.
		if flagValue == flagWarning.ConditionalValue {
			fmt.Printf("Warning for flag '%v': %v\n", flagWarning.Flag, flagWarning.Warning)
		}

	}

	return nil
}
