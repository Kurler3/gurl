package requests

import (
	"fmt"

	"github.com/Kurler3/gurl/classes/gurl"
)

func MakeRequest(g *gurl.Gurl) {

	client, req, err := GetReqAndCLient(g)

	if err != nil {
		fmt.Println("Error setting up request:", err)
		return
	}

	MakeRequestWithClient(client, req, g)
}
