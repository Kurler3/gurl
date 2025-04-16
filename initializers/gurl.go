package initializers

import (
	"fmt"
	"os"

	"github.com/Kurler3/gurl/checks"
	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/parser"
)

func InitGurl() (gurl.Gurl, error) {

	g := gurl.Gurl{}

	// Parse the rest of the arguments
	for i := 1; i < len(os.Args); i++ {

		flag, value, err := parser.ParseCmdArg(os.Args[i])

		if err != nil {
			return g, err
		}

		err = g.SetFlag(flag, value)

		if err != nil {
			return g, fmt.Errorf("error while setting flag %v: %v", flag, err)
		}

	}

	err := checks.FinalFlagsCheck(&g)

	if err != nil {
		return g, err
	}

	return g, nil
}
