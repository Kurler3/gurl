package requests

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Kurler3/gurl/classes/gurl"
)

func GetReqAndCLient(g *gurl.Gurl) (*http.Client, *http.Request, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: g.SkipTLSVerification,
		},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   g.Timeout,
	}

	url := g.Protocol + "://" + g.Url

	// Conditionally set the body
	var bodyReader io.Reader

	// If there is a body and the method is not GET
	if g.Data != "" && g.Method != http.MethodGet {
		bodyReader = strings.NewReader(g.Data)
	}

	// Create a new req.
	req, err := http.NewRequest(g.Method, url, bodyReader)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, nil, err
	}

	// Set the headers
	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	return client, req, nil
}
