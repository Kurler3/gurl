package requests

import (
	"fmt"

	"github.com/Kurler3/gurl/classes/gurl"
)

func MakeBenchmarkedRequests(g *gurl.Gurl) {
	var sum_elapsed_time float64

	// Make 10 requests.
	for i := range g.RequestsCount {

		elapsed := MakeRequest(g, fmt.Sprintf("request number %v", i))

		sum_elapsed_time += elapsed
	}

	avg_time := sum_elapsed_time / float64(g.RequestsCount)

	fmt.Printf("Average time for the requests is: %.2f \n", avg_time)
}
