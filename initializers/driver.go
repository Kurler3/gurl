package initializers

import (
	"log"

	"github.com/Kurler3/gurl/requests"
)

func Init() {

	//TODO - If second argument is "help", display all the commands available.

	// - Init the Gurl struct by parsing the CMD args and assigning them to the gurl struct.
	g, err := InitGurl()

	if err != nil {
		log.Fatal(err)
	}

	// Make the request.
	requests.MakeRequest(&g)
}
