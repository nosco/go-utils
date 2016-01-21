/*
utils are currently just a wrapper on top of github/segmentio's extremely fast
Camelcase and Snakecase functions, with an added PascalCase.

Thank you @tj for switching to Go just before we did! ;)
*/
package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/segmentio/go-camelcase"
	"github.com/segmentio/go-snakecase"
)

func Slug(str string) string {
	return strings.Replace(snakecase.Snakecase(str), "_", "-", -1)
}

func SnakeCase(str string) string {
	return snakecase.Snakecase(str)
}

func CamelCase(str string) string {
	return camelcase.Camelcase(str)
}

func PascalCase(str string) string {
	out := camelcase.Camelcase(str)
	if len(out) > 0 {
		out = strings.ToUpper(out[0:1]) + out[1:len(out)]
	}
	return out
}

// InterfaceToReflect helps ensure the reflect value is in an editable state
// It will check the type and get the correct reference if possible
// @TODO Make some tests
func InterfaceToReflect(val interface{}) (reflectValue reflect.Value, err error) {
	typ := reflect.TypeOf(val)

	// @TODO Is this correct?
	if typ.String() == "reflect.Value" {
		reflectValue = val.(reflect.Value)

	} else if typ.String()[0:1] != "*" {
		err = errors.New("Please provide a reference to the value")
		return

	} else {
		reflectValue = reflect.ValueOf(val).Elem()
	}

	return
}
