/*
utils are currently just a wrapper on top of github/segmentio's extremely fast
Camelcase and Snakecase functions, with an added PascalCase.

Thank you @tj for switching to Go just before we did! ;)
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

	"github.com/segmentio/go-camelcase"
	"github.com/segmentio/go-snakecase"
)

func Slug(str string) string {
	return strings.Replace(snakecase.Snakecase(str), "_", "-", -1)
}

func UnCase(str string) string {
	str = strings.Replace(snakecase.Snakecase(str), "_", " ", -1)
	str = strings.ToUpper(str[0:1]) + str[1:]
	return str
}

func SnakeCase(str string) string {
	return snakecase.Snakecase(str)
}

func KebabCase(str string) string {
	return strings.Replace(snakecase.Snakecase(str), "_", "-", -1)
}

func CamelCase(str string) string {
	return camelcase.Camelcase(str)
}

func PascalCase(str string) string {
	out := camelcase.Camelcase(str)
	if len(out) > 0 {
		out = strings.ToUpper(out[0:1]) + out[1:]
	}
	return out
}

func StringInSlice(searchStr string, strs []string) bool {
	for _, str := range strs {
		if searchStr == str {
			return true
		}
	}
	return false
}

func UniqueInts(arr []int) (unique []int) {
	tmpMap := map[int]bool{}
	for i := 0; i < len(arr); i++ {
		tmpMap[arr[i]] = true
	}
	for val := range tmpMap {
		unique = append(unique, val)
	}
	return
}

func UniqueStrings(arr []string) (unique []string) {
	tmpMap := map[string]bool{}
	for i := 0; i < len(arr); i++ {
		tmpMap[arr[i]] = true
	}
	for val := range tmpMap {
		unique = append(unique, val)
	}
	return
}

func Unique(arr interface{}) (unique interface{}, err error) {
	arrType := reflect.TypeOf(arr)
	arrValue := reflect.ValueOf(arr)

	if arrType.Kind().String() != "slice" {
		return nil, errors.New("Not a slice")
	}

	tmpMap := map[interface{}]bool{}
	for i := 0; i < arrValue.Len(); i++ {
		tmpMap[arrValue.Index(i).Interface()] = true
	}

	newArr := reflect.MakeSlice(arrType, 0, arrValue.Len())
	for val := range tmpMap {
		newArr = reflect.Append(newArr, reflect.ValueOf(val))
	}

	unique = newArr.Interface()

	return
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

var callerRE *regexp.Regexp

func init() {
	// Matching e.g. (*ServiceName).ServiceMethod
	callerRE = regexp.MustCompile("(?:\\(\\*{0,1}([^\\)]*?)\\)|([^\\.]+))\\.([^\\.]+)$")
}

func GetCallerName(skip int) (callerName string) {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	pcFunc := runtime.FuncForPC(pc)
	matches := callerRE.FindStringSubmatch(pcFunc.Name())

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
	matches := callerRE.FindStringSubmatch(pcFunc.Name())

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
