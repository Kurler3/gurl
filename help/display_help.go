package help

import (
	"fmt"
	"sort"

	"github.com/Kurler3/gurl/utils"
)

func DisplayHelp() {
	fmt.Println("Usage:")
	fmt.Println("go run main.go [flags]")

	fmt.Println("\nRequired Flags:")
	for flag := range utils.REQUIRED_FLAGS {
		fmt.Printf("  --%-20s Required\n", flag)
	}

	fmt.Println("\nFlags are specified with the following format: --[flag]=[value]")

	fmt.Println("\nAvailable Flags:")
	// Collect and sort for consistent display
	var flags []string
	for flag := range utils.AVAILABLE_FLAGS {
		flags = append(flags, flag)
	}
	sort.Strings(flags)

	for _, flag := range flags {
		short := ""
		for k, v := range utils.SHORT_FLAG_TO_LONG_FLAG_MAP {
			if v == flag {
				short = k
				break
			}
		}
		if short != "" {
			fmt.Printf("  --%-2s, --%-20s\n", short, flag)
		} else {
			fmt.Printf("      --%-20s\n", flag)
		}
	}

	fmt.Println("\nBoolean Flags:")
	for flag := range utils.BOOL_FLAGS {
		fmt.Printf("  --%s\n", flag)
	}

	fmt.Println("\nDefault Values:")
	for flag, value := range utils.DEFAULT_FLAG_VALUES {
		fmt.Printf("  --%-20s %v\n", flag, value)
	}

	fmt.Println("\nSupported HTTP Methods:")
	for method := range utils.AVAILABLE_METHODS {
		fmt.Printf("  %s\n", method)
	}
}
