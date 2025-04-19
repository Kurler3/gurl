package requests

import (
	"fmt"
	"net/http"

	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/response"
	"github.com/Kurler3/gurl/utils"
)

func MakeRequestWithClient(
	client *http.Client,
	req *http.Request,
	g *gurl.Gurl,
) float64 {

	var err error
	var res *http.Response

	elapsed := utils.WithTimer(
		func() {
			// Send request
			res, err = client.Do(req)
		},
	)

	if err != nil {
		fmt.Println("Error making request:", err)
		return elapsed
	}

	response.PrintResponse(res, g)

	return elapsed
}
