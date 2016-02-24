/*
utils are currently just a wrapper on top of github/segmentio's extremely fast
Camelcase and Snakecase functions, with a couple of added cases.

Thank you @tj for switching to Go just before we did! ;)

All case types takes Go's annoying ID convention into consideration...
I.e: id -> ID, Id -> ID, ID -> id, ID !-> i_d, ID -> i-d, Identifier -> identifier
*/
package utils

import (
	"errors"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var (
	reUpperID      *regexp.Regexp
	reUpperFirstID *regexp.Regexp
	reCaller       *regexp.Regexp
)

func init() {
	reUpperID = regexp.MustCompile("(^|\\W)id(\\W|$)")
	reUpperFirstID = regexp.MustCompile("^Id([A-Z]|$)")
	// Matching e.g. (*ServiceName).ServiceMethod
	reCaller = regexp.MustCompile("(?:\\(\\*{0,1}([^\\)]*?)\\)|([^\\.]+))\\.([^\\.]+)$")
}

// InterfaceToReflect helps ensure the reflect value is in an editable state
// It will check the type and get the correct reference if possible
// TODO(morphar) Make some tests
func InterfaceToReflect(val interface{}) (reflectValue reflect.Value, err error) {
	typ := reflect.TypeOf(val)

	// TODO(morphar) Is this correct?
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

func GetCallerName(skip int) (callerName string) {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	pcFunc := runtime.FuncForPC(pc)
	matches := reCaller.FindStringSubmatch(pcFunc.Name())

	if matches == nil || len(matches) != 4 {
		return
	}

	return matches[1] + matches[2] + "." + matches[3]
}

func GetCallerNames(skip int) (typeName, callerName string) {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	pcFunc := runtime.FuncForPC(pc)
	matches := reCaller.FindStringSubmatch(pcFunc.Name())

	if matches == nil || len(matches) != 4 {
		return
	}

	return matches[1] + matches[2], matches[3]
}

func GetCallStack() (stack []string) {
	pcs := make([]uintptr, 50)
	pcCount := runtime.Callers(2, pcs)

	pathRE := regexp.MustCompile("^.*/")

	for i := 0; i < pcCount; i++ {
		pcFunc := runtime.FuncForPC(pcs[i])
		file, line := pcFunc.FileLine(pcs[i])
		fileName := pathRE.ReplaceAllString(file, "")
		stack = append(stack, "["+fileName+":"+strconv.Itoa(line)+"]: "+pcFunc.Name())
	}

	return
}

func PrintCallStack() {
	stack := GetCallStack()
	log.Println("Call stack:\n", strings.Join(stack[1:len(stack)-1], "\n "))
}

func Slug(str string) string {
	return strings.Replace(SnakeCase(str), "_", "-", -1)
}

func UnCase(str string) string {
	str = strings.Replace(SnakeCase(str), "_", " ", -1)
	str = reUpperID.ReplaceAllString(str, "${1}ID${2}")
	str = strings.ToUpper(str[0:1]) + str[1:]
	return str
}

func KebabCase(str string) string {
	return strings.Replace(SnakeCase(str), "_", "-", -1)
}

func PascalCase(str string) string {
	str = CamelCase(str)
	if len(str) > 0 {
		str = strings.ToUpper(str[0:1]) + str[1:]
	}
	str = reUpperFirstID.ReplaceAllString(str, "ID$1")
	return str
}

// Snakecase the given `str`.
func SnakeCase(str string) string {
	var b [1024]byte
	max := 1024
	l := len(str)
	ret := ""
	bi := 0
	i := 0

	for i < l {
		for i < l && !isWord(str[i]) {
			i++
		}

		for i < l && isUpper(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && isPart(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && !isWord(str[i]) {
			i++
		}

		if strings.ToUpper(string(b[:2])) == "ID" && bi == 2 {
			ret += "id" + "_"
		} else if strings.ToUpper(string(b[:2])) == "ID" && bi > 2 && !isPart(b[2]) {
			ret += strings.ToLower(string(b[:2])) + "_"
			if bi > 3 {
				ret += strings.ToLower(string(b[2:bi])) + "_"
			}
		} else {
			ret += strings.ToLower(string(b[:bi])) + "_"
		}

		bi = 0
	}

	if len(ret) > 0 {
		ret = ret[:len(ret)-1]
	}

	return ret
}

// camelCase the given `str`.
func CamelCase(str string) string {
	var b [1024]byte
	max := 1024
	l := len(str)
	ret := ""
	bi := 0
	i := 0
	first := true

	for i < l {

		for i < l && !isWord(str[i]) {
			i++
		}

		for i < l && isUpper(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && isPart(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && !isWord(str[i]) {
			i++
		}

		if first {
			ret += strings.ToLower(string(b[:bi]))
			first = false
		} else if strings.ToUpper(string(b[:2])) == "ID" && bi == 2 {
			ret += "ID"
		} else if strings.ToUpper(string(b[:2])) == "ID" && bi > 2 && !isPart(b[2]) {
			ret += strings.ToUpper(string(b[:3]))
			if bi > 3 {
				ret += strings.ToLower(string(b[3:bi]))
			}
		} else {
			// .ToTitle is weird in Go
			ret += strings.ToUpper(string(b[:1]))
			ret += strings.ToLower(string(b[1:bi]))
		}
		bi = 0
	}

	if len(ret) > 0 {
		ret = ret[:]
	}

	return ret
}

func isPart(c byte) bool {
	return isLower(c) || isDigit(c)
}

func isWord(c byte) bool {
	return isLetter(c) || isDigit(c)
}

func isLetter(c byte) bool {
	return isLower(c) || isUpper(c)
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
