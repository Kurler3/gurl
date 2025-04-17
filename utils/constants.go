package utils

import "net/http"

const (
	MethodFlag        = "method"
	ShortMethodFlag   = "M"
	UrlFlag           = "url"
	ShortUrlFlag      = "U"
	ProtocolFlag      = "protocol"
	ShortProtocolFlag = "P"
	VerboseFlag       = "verbose"
	ShortVerboseFlag  = "V"
	HeadersFlag       = "headers"
	ShortHeadersFlag  = "H"

	HTTP  = "http"
	HTTPS = "https"
)

var SHORT_FLAG_TO_LONG_FLAG_MAP = map[string]string{
	ShortMethodFlag:   MethodFlag,
	ShortUrlFlag:      UrlFlag,
	ShortProtocolFlag: ProtocolFlag,
	ShortVerboseFlag:  VerboseFlag,
	ShortHeadersFlag:  HeadersFlag,
}

var AVAILABLE_FLAGS = map[string]struct{}{
	UrlFlag:      {},
	MethodFlag:   {},
	ProtocolFlag: {},
}

var AVAILABLE_METHODS = map[string]struct{}{
	http.MethodGet: {},
}

var REQUIRED_FLAGS = map[string]struct{}{
	UrlFlag:    {},
	MethodFlag: {},
}

var DEFAULT_FLAG_VALUES = map[string]any{
	MethodFlag:   http.MethodGet,
	ProtocolFlag: HTTPS,
	VerboseFlag:  true,
	HeadersFlag:  map[string]string{},
}

var BOOL_FLAGS = map[string]struct{}{
	VerboseFlag:      {},
	ShortVerboseFlag: {},
}
