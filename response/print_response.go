package response

import (
	"fmt"
	"io"
	"net/http"
)

func PrintResponse(res *http.Response, verbose bool) {

	fmt.Println("---------------------------------------")
	fmt.Println("---------- RESPONSE DATA --------------")
	fmt.Println("---------------------------------------")

	if res == nil {
		fmt.Println("No response received.")
		return
	}

	if verbose {

		// Status line
		fmt.Println("Status:", res.Status)

		// Headers
		fmt.Println("Headers:")
		for k, v := range res.Header {
			fmt.Printf("  %s: %s\n", k, v)
		}

	}

	// Body
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println("Body:")
	fmt.Println(string(bodyBytes))
}
