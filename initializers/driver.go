package initializers

import (
	"log"
	"os"

	"github.com/Kurler3/gurl/help"
	"github.com/Kurler3/gurl/requests"
)

func Init() {

	//  If second argument is "help", display all the commands available.
	if len(os.Args) == 2 && os.Args[1] == "help" {
		help.DisplayHelp()
		return
	}

	// - Init the Gurl struct by parsing the CMD args and assigning them to the gurl struct.
	g, err := InitGurl()

	if err != nil {
		log.Fatal(err)
	}

	// If benchmark mode
	if g.Benchmark {
		requests.MakeBenchmarkedRequests(&g)
		return
	}

	// Make a normal request.
	requests.MakeRequest(&g, "")
}
