package gurl

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Kurler3/gurl/response"
	"github.com/Kurler3/gurl/utils"
)

type Gurl struct {
	Url      string
	Method   string
	Protocol string
	Verbose  bool
	Headers  map[string]string
}

/////////////////////////////////////////////////////////
// GETTERS & SETTERS ////////////////////////////////////
/////////////////////////////////////////////////////////

func (g *Gurl) GetFlag(flag string) (any, error) {
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
				reflect.Interface,
			)

			if err != nil {
				return nil, err
			}

			return field.Interface(), nil

		}
	}

}

func (g *Gurl) SetFlag(flag string, value any) error {

	switch flag {
	case utils.UrlFlag:
		{
			var protocolInUrl string

			if strings.HasPrefix(value.(string), "http://") {
				protocolInUrl = "http"
			}

			if strings.HasPrefix(value.(string), "https://") {
				protocolInUrl = "https"
			}

			// If there is no protocol in the url, save it as is.
			if protocolInUrl == "" {
				g.Url = value.(string)
				return nil
			}

			// If has protocol and there's already a protocol as well defined, check if the same.
			if g.Protocol != "" && g.Protocol != protocolInUrl {
				return fmt.Errorf("protocol specified in the flag \"%v\" not the same as the one specified in the url. Please either remove the one in the url or specify the same one in the \"%v\" flag", utils.ProtocolFlag, utils.ProtocolFlag)
			}

			g.Protocol = protocolInUrl

			// Strip the protocol from the url and save it.
			urlWithoutProtocol := strings.TrimPrefix(value.(string), protocolInUrl)

			// Set the url without protocol.
			g.Url = urlWithoutProtocol

			return nil
		}

	case utils.MethodFlag:
		{
			g.Method = strings.ToUpper(value.(string))
		}

	default:
		{

			field, err := utils.FindStructField(
				g,
				utils.CapitalizeFirst(flag),
				true,
				reflect.Interface,
			)

			if err != nil {
				return err
			}

			val := reflect.ValueOf(value)

			// Convert value if needed
			if val.Type().AssignableTo(field.Type()) {
				field.Set(val)
			} else if val.Type().ConvertibleTo(field.Type()) {
				field.Set(val.Convert(field.Type()))
			} else {
				return fmt.Errorf("cannot assign value of type %s to field of type %s", val.Type(), field.Type())
			}

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

	url := g.Protocol + "://" + g.Url

	// Create a new req.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	// Send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	response.PrintResponse(res, g.Verbose)
}

// POST

// PATCH

// DELETE

// PUT

// HEAD

//////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////
