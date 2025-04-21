package response

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Kurler3/gurl/classes/gurl"
)

func PrintResponse(
	res *http.Response,
	g *gurl.Gurl,
	requestTitle string,
	elapsed float64,
) strings.Builder {

	// Init a string builder
	var output strings.Builder

	writeLine := func(s string) {
		fmt.Println(s)
		output.WriteString(s + "\n")
	}

	// If is in benchmark mode, write some title for the request.
	if g.Benchmark && requestTitle != "" {
		writeLine("------------------------------------------------------")
		writeLine(fmt.Sprintf("---------------------%v-----------------------", requestTitle))
		writeLine("------------------------------------------------------")
	}

	if g.Verbose {
		writeLine("------------------------------")
		writeLine("-------- Request data --------")
		writeLine("------------------------------")
		writeLine(fmt.Sprintf("%v", *g))
		writeLine(fmt.Sprintf("Request %v took %.2f seconds.\n", requestTitle, elapsed))
	}

	writeLine("---------------------------------------")
	writeLine("---------- RESPONSE DATA --------------")
	writeLine("---------------------------------------")

	if res == nil {
		writeLine("No response received.")
		return output
	}

	if g.Verbose {

		// Status line
		writeLine(fmt.Sprintf("Status: %v", res.Status))

		// Headers
		writeLine("Headers:")
		for k, v := range res.Header {
			writeLine(fmt.Sprintf("  %s: %s\n", k, v))
		}

	}

	// Body
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		writeLine(fmt.Sprintf("Error reading body: %v", err))
		return output
	}

	writeLine("Body:")
	writeLine(string(bodyBytes))

	return output

}
