package requests

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Kurler3/gurl/classes/gurl"
	"github.com/Kurler3/gurl/response"
	"github.com/Kurler3/gurl/utils"
)

func MakeRequestWithClient(
	client *http.Client,
	req *http.Request,
	g *gurl.Gurl,
	requestTitle string,
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

	output := response.PrintResponse(res, g, requestTitle, elapsed)

	// If defined an output path
	if g.Output != "" {

		// Make the dirs until the file.
		//?? 0755: 7 for owner (r+w+x), 5 (r+x) for groups and others
		err = os.MkdirAll(filepath.Dir(g.Output), 0755)
		if err != nil {
			fmt.Println("Failed to create directories:", err)
		}

		// Open the file in append mode. If it doesn't exist, create it.
		f, err := os.OpenFile(g.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to write response to file:", err)
		}
		defer f.Close()

		_, err = f.Write([]byte(output.String()))

		if err != nil {
			fmt.Printf("Error writing response to file: %v", err)
		}

		fmt.Println("Full response written to:", g.Output)
	}

	return elapsed
}
