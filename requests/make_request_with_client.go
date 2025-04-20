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

		// Write to the file, and create it if it wasn't yet
		//?? 644: 6 (r+w), 4 (r)
		err = os.WriteFile(g.Output, []byte(output.String()), 0644)
		if err != nil {
			fmt.Println("Failed to write response to file:", err)
		}

		fmt.Println("Full response written to:", g.Output)
	}

	return elapsed
}
