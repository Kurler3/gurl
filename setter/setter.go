package setter

import (
	"errors"
	"reflect"

	"github.com/Kurler3/gurl/classes"
	"github.com/Kurler3/gurl/setter/flag_setters"
	"github.com/Kurler3/gurl/utils"
)

type FlagSetter func(g *classes.Gurl, val string) error

var FlagSetters = map[string]FlagSetter{
	utils.UrlFlag:    flag_setters.UrlSetter,
	utils.MethodFlag: flag_setters.MethodSetter,
}

func GetFlagSetter(flag string) FlagSetter {

	// Check if there's a setter in the flag setters.
	if flagSetter, ok := FlagSetters[flag]; ok {
		return flagSetter
	}

	// If there isn't => return a new func as a default setter.
	return func(g *classes.Gurl, val string) error {

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
		field.SetString(val)

		return nil
	}

}
