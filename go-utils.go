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

func SnakeCase(in string) (out string) {
	return snakecase.Snakecase(in)
}

func CamelCase(in string) (out string) {
	return camelcase.Camelcase(in)
}

func PascalCase(in string) (out string) {
	out = camelcase.Camelcase(in)
	if len(out) > 0 {
		out = strings.ToUpper(out[0:1]) + out[1:len(out)]
	}
	return
}
