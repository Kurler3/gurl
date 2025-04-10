package utils

import "net/http"

const (
	MethodFlag = "method"
	UrlFlag    = "url"
)

var AVAILABLE_FLAGS = map[string]struct{}{
	UrlFlag:    {},
	MethodFlag: {},
}

var AVAILABLE_METHODS = map[string]struct{}{
	http.MethodGet: {},
}
