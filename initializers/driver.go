package initializers

import (
	"fmt"
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

	// If benchmark mode
	if g.Benchmark {

		// Make 10 requests.
		//TODO - Should be able to choose how many.
		for i := range 10 {

			elapsed := requests.MakeRequest(&g)

			fmt.Printf("Request number %v took %.2f seconds.\n", i, elapsed)

		}

		//TODO - Display some sort of average time and statistics and stuff based on all the times.

		return

	}

	elapsed := requests.MakeRequest(&g)

	fmt.Printf("Request took %.2f seconds.\n", elapsed)

}
