package initializers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

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

		//TODO - Use g.SetFlag instead. instead.

		err = g.SetFlag(flag, value)

		if err != nil {
			return g, fmt.Errorf("error while setting flag %v: %v", flag, err)
		}

	}

	//TODO - Should probably put this in a different function.
	//TODO - Essentially, the final check on the required flags and default assigns.

	if g.Url == "" {
		return g, errors.New("no url specified")
	}

	//TODO If there is no protocol => error
	if g.Protocol == "" {
		return g, errors.New("no protocol specified")
	}

	// If no method defined, by default assign to GET
	if g.Method == "" {
		g.Method = http.MethodGet
	}

	return g, nil
}
