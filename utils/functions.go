package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func GetMapKeysAsArray[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func CapitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FindStructField(
	s any,
	fieldName string,
	checkWhetherCanSet bool,
	kindToCheck reflect.Kind,
) (reflect.Value, error) {

	v := reflect.ValueOf(s).Elem()

	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return field, errors.New("no such field: " + fieldName)
	}

	if checkWhetherCanSet && !field.CanSet() {
		return field, errors.New("cannot set field: " + fieldName)
	}

	if kindToCheck != reflect.Interface && field.Kind() != kindToCheck {
		return field, errors.New("unsupported field type")
	}

	return field, nil
}

func WithTimer(
	fn func(),
) float64 {
	start := time.Now()

	fn()

	now := time.Now()

	elapsed := now.Sub(start).Seconds()

	fmt.Printf("This function took %.2f seconds.", elapsed)

	return elapsed
}

func IsEmpty(x any) bool {
	if x == nil {
		return true
	}

	v := reflect.ValueOf(x)

	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}
