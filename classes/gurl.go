package classes

import (
	"fmt"
	"net/http"
	"reflect"
)

type Gurl struct {
	Url      string
	Method   string
	Protocol string
}

/////////////////////////////////////////////////////////
// METHODS //////////////////////////////////////////////
/////////////////////////////////////////////////////////

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

// GET
func (g *Gurl) MakeGetRequest() {

	resp, err := http.Get(g.Protocol + g.Url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

// POST

// PATCH

// DELETE

// PUT

// HEAD

//////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////
