package requests

import (
	"fmt"
	"net/http"

	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/response"
)

func MakeRequestWithClient(
	client *http.Client,
	req *http.Request,
	g *gurl.Gurl,
) {

	// Send request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	response.PrintResponse(res, g)

}
