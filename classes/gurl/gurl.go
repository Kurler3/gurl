package gurl

import (
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

func (g *Gurl) GetFlag(flag string) (interface{}, error) {
	switch flag {
	case utils.UrlFlag:
		return g.Url, nil
	case utils.MethodFlag:
		return g.Method, nil
	default:
		{

			field, err := utils.FindStructField(
				g,
				utils.CapitalizeFirst(flag),
				false,
				reflect.String,
			)

			if err != nil {
				return nil, err
			}

			return field.String(), nil

		}
	}

}

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

			field, err := utils.FindStructField(
				g,
				utils.CapitalizeFirst(flag),
				true,
				reflect.String,
			)

			if err != nil {
				return err
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
