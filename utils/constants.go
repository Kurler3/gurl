package utils

import (
	"net/http"
	"time"
)

const (
	MethodFlag               = "method"
	ShortMethodFlag          = "M"
	UrlFlag                  = "url"
	ShortUrlFlag             = "U"
	ProtocolFlag             = "protocol"
	ShortProtocolFlag        = "P"
	VerboseFlag              = "verbose"
	ShortVerboseFlag         = "V"
	HeadersFlag              = "headers"
	ShortHeadersFlag         = "H"
	SkipTLSVerification      = "skipTLSVerification"
	ShortSkipTLSVerification = "K"
	OutputFlag               = "output"
	ShortOutputFlag          = "O"
	TimeoutFlag              = "timeout"
	ShortTimeoutFlag         = "T"
	DataFlag                 = "data"
	ShortDataFlag            = "D"
	BenchmarkFlag            = "benchmark"
	ShortBenchmarkFlag       = "B"
	RequestsCountFlag        = "requestsCount"
	ShortRequestsCountFlag   = "R"

	HTTP  = "http"
	HTTPS = "https"
)

var SHORT_FLAG_TO_LONG_FLAG_MAP = map[string]string{
	ShortMethodFlag:          MethodFlag,
	ShortUrlFlag:             UrlFlag,
	ShortProtocolFlag:        ProtocolFlag,
	ShortVerboseFlag:         VerboseFlag,
	ShortHeadersFlag:         HeadersFlag,
	ShortSkipTLSVerification: SkipTLSVerification,
	ShortOutputFlag:          OutputFlag,
	ShortTimeoutFlag:         TimeoutFlag,
	ShortDataFlag:            DataFlag,
	ShortBenchmarkFlag:       BenchmarkFlag,
	ShortRequestsCountFlag:   RequestsCountFlag,
}

var AVAILABLE_FLAGS = map[string]struct{}{
	UrlFlag:             {},
	MethodFlag:          {},
	ProtocolFlag:        {},
	HeadersFlag:         {},
	SkipTLSVerification: {},
	OutputFlag:          {},
	TimeoutFlag:         {},
	DataFlag:            {},
	BenchmarkFlag:       {},
	RequestsCountFlag:   {},
}

var AVAILABLE_METHODS = map[string]struct{}{
	http.MethodGet:    {},
	http.MethodPost:   {},
	http.MethodPatch:  {},
	http.MethodPut:    {},
	http.MethodDelete: {},
}

var REQUIRED_FLAGS = map[string]struct{}{
	UrlFlag:    {},
	MethodFlag: {},
}

var DEFAULT_FLAG_VALUES = map[string]any{
	MethodFlag:          http.MethodGet,
	ProtocolFlag:        HTTPS,
	VerboseFlag:         true,
	SkipTLSVerification: false,
	HeadersFlag:         map[string]string{},
	TimeoutFlag:         1000 * time.Hour,
	BenchmarkFlag:       false,
	RequestsCountFlag:   10,
}

var BOOL_FLAGS = map[string]struct{}{
	VerboseFlag:              {},
	ShortVerboseFlag:         {},
	SkipTLSVerification:      {},
	ShortSkipTLSVerification: {},
	BenchmarkFlag:            {},
	ShortBenchmarkFlag:       {},
}
