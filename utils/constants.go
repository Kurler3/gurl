package utils

import "net/http"

const (
	MethodFlag        = "method"
	ShortMethodFlag   = "M"
	UrlFlag           = "url"
	ShortUrlFlag      = "U"
	ProtocolFlag      = "protocol"
	ShortProtocolFlag = "P"

	HTTP  = "http"
	HTTPS = "https"
)

var SHORT_FLAG_TO_LONG_FLAG_MAP = map[string]string{
	ShortMethodFlag:   MethodFlag,
	ShortUrlFlag:      UrlFlag,
	ShortProtocolFlag: ProtocolFlag,
}

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
