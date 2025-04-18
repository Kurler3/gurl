package requests

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/response"
)

func MakeGetRequest(g *gurl.Gurl) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: g.SkipTLSVerification,
		},
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second, //TODO - Change if timeout is specified.
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}

	url := g.Protocol + "://" + g.Url

	// Create a new req.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	// Send request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	response.PrintResponse(res, g)

}
