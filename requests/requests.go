package requests

import (
	"log"
	"net/http"

	"github.com/Kurler3/gurl/classes/gurl"
)

var methodToHandlerMap = map[string]func(g *gurl.Gurl){
	http.MethodGet: MakeGetRequest,
}

func MakeRequest(g *gurl.Gurl) {
	if handler, ok := methodToHandlerMap[g.Method]; ok {
		handler(g)
		return
	}
	log.Printf("Unsupported HTTP Method: %s", g.Method)
}
