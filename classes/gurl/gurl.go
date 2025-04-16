package gurl

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Kurler3/gurl/utils"
)

type Gurl struct {
	Url      string
	Method   string
	Protocol string
}

/////////////////////////////////////////////////////////
// GETTERS & SETTERS ////////////////////////////////////
/////////////////////////////////////////////////////////

//TODO - Figure out.

// func (g *Gurl) GetFlag(flag string) error {

// }

func (g *Gurl) SetFlag(flag string, value string) error {

	switch flag {
	case utils.UrlFlag:
		{
			var protocolInUrl string

			if strings.HasPrefix(value, "http://") {
				protocolInUrl = "http"
			}

			if strings.HasPrefix(value, "https://") {
				protocolInUrl = "https"
			}

			// If there is no protocol in the url, save it as is.
			if protocolInUrl == "" {
				g.Url = value
				return nil
			}

			// If has protocol and there's already a protocol as well defined, check if the same.
			if g.Protocol != "" && g.Protocol != protocolInUrl {
				return fmt.Errorf("protocol specified in the flag \"%v\" not the same as the one specified in the url. Please either remove the one in the url or specify the same one in the \"%v\" flag", utils.ProtocolFlag, utils.ProtocolFlag)
			}

			g.Protocol = protocolInUrl

			// Strip the protocol from the url and save it.
			urlWithoutProtocol := strings.TrimPrefix(value, protocolInUrl)

			// Set the url without protocol.
			g.Url = urlWithoutProtocol

			return nil
		}

	case utils.MethodFlag:
		{
			g.Method = strings.ToUpper(value)
		}

	default:
		{
			// Get the reflect.Value of the struct
			v := reflect.ValueOf(g).Elem()

			// Convert flag to TitleCase (optional, depending on your field names)
			fieldName := utils.CapitalizeFirst(flag)

			field := v.FieldByName(fieldName)

			// Check if the field actually exists
			if !field.IsValid() {
				return errors.New("no such field: " + fieldName)
			}

			// Check if can set the field
			if !field.CanSet() {
				return errors.New("cannot set field: " + fieldName)
			}

			// Check if the type of the field is string
			if field.Kind() != reflect.String {
				return errors.New("unsupported field type: only string is supported")
			}

			// Set the field.
			field.SetString(value)

		}
	}

	return nil
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
