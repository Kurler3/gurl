package initializers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/Kurler3/gurl/parser"
	"github.com/Kurler3/gurl/utils"
)

type Gurl struct {
	Url    string
	Method string
}

type FlagSetter func(g *Gurl, val string) error

var flagSetters = map[string]FlagSetter{
	utils.UrlFlag: func(g *Gurl, val string) error {
		g.Url = val
		return nil
	},
	utils.MethodFlag: func(g *Gurl, val string) error {
		g.Method = strings.ToUpper(val)
		return nil
	},
}

func InitGurl() (Gurl, error) {

	g := Gurl{}

	// Parse the rest of the arguments
	for i := 1; i < len(os.Args); i++ {

		flag, value, err := parser.ParseCmdArg(os.Args[i])

		if err != nil {
			return g, err
		}

		// Assign flag and value to the g struct.
		flagSetter, ok := flagSetters[flag]

		if !ok {
			return g, fmt.Errorf("no flag setter found for flag %v", flag)
		}

		err = flagSetter(&g, value)

		if err != nil {
			return g, fmt.Errorf("error while setting flag %v: %v", flag, err)
		}

	}

	if g.Url == "" {
		return g, errors.New("no url specified")
	}

	return g, nil
}

func (g Gurl) String() string {

	val := reflect.ValueOf(g)
	typ := reflect.TypeOf(g)

	result := ""

	for i := range val.NumField() {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		result += fmt.Sprintf("%s: %v\n", fieldName, fieldValue)
	}
	return result
}

func (g *Gurl) MakeGetRequest() {

	resp, err := http.Get(g.Url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func (g *Gurl) MakeRequest() {

}
