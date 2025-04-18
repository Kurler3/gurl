package response

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func PrintResponse(
	res *http.Response,
	verbose bool,
	outputPath string,
) {

	// Init a string builder
	var output strings.Builder

	writeLine := func(s string) {
		fmt.Println(s)
		output.WriteString(s + "\n")
	}

	// Init function to write a line.

	writeLine("---------------------------------------")
	writeLine("---------- RESPONSE DATA --------------")
	writeLine("---------------------------------------")

	if res == nil {
		writeLine("No response received.")
		return
	}

	if verbose {

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
	if outputPath != "" {

		// Make the dirs until the file.
		//?? 0755: 7 for owner (r+w+x), 5 (r+x) for groups and others
		err = os.MkdirAll(filepath.Dir(outputPath), 0755)
		if err != nil {
			fmt.Println("Failed to create directories:", err)
			return
		}

		// Write to the file, and create it if it wasn't yet
		//?? 644: 6 (r+w), 4 (r)
		err = os.WriteFile(outputPath, []byte(output.String()), 0644)
		if err != nil {
			fmt.Println("Failed to write response to file:", err)
			return
		}

		fmt.Println("Full response written to:", outputPath)
	}

}
