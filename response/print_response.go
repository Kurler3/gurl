package response

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kurler3/gurl/classes/gurl"
)

func PrintResponse(
	res *http.Response,
	g *gurl.Gurl,
) {

	// Init a string builder
	var output strings.Builder

	writeLine := func(s string) {
		fmt.Println(s)
		output.WriteString(s + "\n")
	}

	if g.Verbose {
		writeLine("------------------------------")
		writeLine("-------- Request data --------")
		writeLine("------------------------------")
		writeLine(fmt.Sprintf("%v", *g))
	}

	writeLine("---------------------------------------")
	writeLine("---------- RESPONSE DATA --------------")
	writeLine("---------------------------------------")

	if res == nil {
		writeLine("No response received.")
		return
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
		return
	}

	writeLine("Body:")
	writeLine(string(bodyBytes))

	// If defined an output path
	if g.Output != "" {

		// Make the dirs until the file.
		//?? 0755: 7 for owner (r+w+x), 5 (r+x) for groups and others
		err = os.MkdirAll(filepath.Dir(g.Output), 0755)
		if err != nil {
			fmt.Println("Failed to create directories:", err)
			return
		}

		// Write to the file, and create it if it wasn't yet
		//?? 644: 6 (r+w), 4 (r)
		err = os.WriteFile(g.Output, []byte(output.String()), 0644)
		if err != nil {
			fmt.Println("Failed to write response to file:", err)
			return
		}

		fmt.Println("Full response written to:", g.Output)
	}

}
