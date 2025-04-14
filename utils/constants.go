package utils

import "net/http"

const (
	MethodFlag   = "method"
	UrlFlag      = "url"
	ProtocolFlag = "protocol"

	HTTP  = "http"
	HTTPS = "https"
)

var AVAILABLE_FLAGS = map[string]struct{}{
	UrlFlag:      {},
	MethodFlag:   {},
	ProtocolFlag: {},
}

var AVAILABLE_METHODS = map[string]struct{}{
	http.MethodGet: {},
}

// var FLAG_TO_GURL_FIELD = map[string]string{
// 	MethodFlag:   "Method",
// 	UrlFlag:      "Url",
// 	ProtocolFlag: "Protocol",
// }
