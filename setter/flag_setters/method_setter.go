package flag_setters

import (
	"strings"

	"github.com/Kurler3/gurl/classes"
)

func MethodSetter(g *classes.Gurl, val string) error {
	g.Method = strings.ToUpper(val)
	return nil
}
