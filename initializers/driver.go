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

		var sum_elapsed_time float64

		fmt.Println("g.RequestsCount: ", g.RequestsCount)

		// Make 10 requests.
		for i := range g.RequestsCount {

			elapsed := requests.MakeRequest(&g, fmt.Sprintf("request number %v", i))

			sum_elapsed_time += elapsed
		}

		avg_time := sum_elapsed_time / float64(g.RequestsCount)

		//TODO
		fmt.Printf("Average time for the requests is: %.2f \n", avg_time)

		return

	}

	// Make a normal request.
	requests.MakeRequest(&g, "")
}
