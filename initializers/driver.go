package initializers

import (
	"fmt"
	"log"
)

func Init() {

	//TODO - Display basic usage of gurl.

	//TODO - If second argument is "help", display all the commands available.

	// - Init the Gurl struct by parsing the CMD args and assigning them to the gurl struct.
	g, err := InitGurl()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(g)

	// // Init map between the method and the handler
	// var methodToHandlerMap = map[string]func(){
	// 	http.MethodGet: g.MakeGetRequest,
	// }

	// // Depending on the method, call the correct function of the Gurl struct.
	// methodToHandlerMap[g.Method]()
}
