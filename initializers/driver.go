package initializers

import (
	"fmt"
	"log"
	"net/http"
)

func Init() {

	//TODO - If second argument is "help", display all the commands available.

	// - Init the Gurl struct by parsing the CMD args and assigning them to the gurl struct.
	g, err := InitGurl()

	if err != nil {
		log.Fatal(err)
	}

	if g.Verbose {

		fmt.Println("------------------------------")
		fmt.Println("-------- Request data --------")
		fmt.Println("------------------------------")
		fmt.Println("\n", g)
	}

	// Init map between the method and the handler
	methodToHandlerMap := map[string]func(){
		http.MethodGet: g.MakeGetRequest,
	}

	// Depending on the method, call the correct function of the Gurl struct.
	if handler, ok := methodToHandlerMap[g.Method]; ok {
		handler()
		return
	}

	// If no handler found.
	log.Printf("Unsupported HTTP method: %s", g.Method)

}
