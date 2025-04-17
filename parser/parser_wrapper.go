package parser

// Need this to be able to use ANY value returned from a parser and still put it in the map in parser.go

func ParserWrapper[T any](parser func(string) (T, error)) func(string) (any, error) {

	return func(s string) (any, error) {
		return parser(s)
	}

}
