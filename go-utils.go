/*
utils are currently just a wrapper on top of github/segmentio's extremely fast
Camelcase and Snakecase functions, with an added PascalCase.

Thank you @tj for switching to Go just before we did! ;)
*/
package utils

import (
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
