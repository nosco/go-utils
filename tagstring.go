package utils

import (
	"errors"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type TagString reflect.StructTag

// Used for sorting
type tagInfo struct {
	name string
	val  string
}

type byName []tagInfo

func (a byName) Len() int      { return len(a) }
func (a byName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool {
	// Favor json tags as they are what tags are generally used for
	if a[i].name == "json" {
		return true
	}
	if a[j].name == "json" {
		return false
	}
	return a[i].name < a[j].name
}

// Initialise the sort regexp
var tagStringRegExp = regexp.MustCompile(`(?:^| )([^:"]+):"([^"]*)"`)

// Sort the tags alphabetically, except favor the json tag
func (tag *TagString) Sort() {
	matches := tagStringRegExp.FindAllStringSubmatch(string(*tag), -1)
	if len(matches) > 0 {
		var tags []tagInfo
		for _, match := range matches {
			tags = append(tags, tagInfo{match[1], match[2]})
		}

		sort.Sort(byName(tags))

		tmpTag := ""
		for _, tagInfo := range tags {
			tmpTag += tagInfo.name + `:"` + tagInfo.val + `" `
		}

		*tag = TagString(tmpTag[:len(tmpTag)-1])
	}
}

func (tag TagString) Get(key string) string {
	return reflect.StructTag(tag).Get(key)
}

// TODO(morphar) All of the following functions can easily be made faster
// Look at reflect's func (tag StructTag) Get(key string) string
// It iterates instead of doing regexps
// These are actually also wrong, as the values can't have " in them

// Set ONLY sets the tag identified by key, if it already exists
// Returns true if a tag was found and changed
func (tag *TagString) Set(key string, val string) (success bool) {
	re, err := regexp.Compile(`(^| )` + key + `:"[^"]*"`)
	if err == nil && re.MatchString(string(*tag)) {
		with := "${1}" + key + `:` + strconv.Quote(val)
		newTag := re.ReplaceAllString(string(*tag), with)
		*tag = TagString(newTag)
		success = true
	}
	return
}

// SetMulti ONLY sets the tags identified by the keys, if they already exists
// Returns false if one or more tag was not changed
func (tag *TagString) SetMulti(keyVals map[string]string) (success bool) {
	success = true
	for key, val := range keyVals {
		res := tag.Set(key, val)
		if !res {
			success = false
		}
	}
	return
}

// Add adds the tag to the TagString
func (tag *TagString) Add(key string, val string) {
	re, err := regexp.Compile(`(^| )` + key + `:"[^"]*"`)
	if err == nil && re.MatchString(string(*tag)) {
		tag.Set(key, val)

	} else {
		newTag := strings.Join([]string{string(*tag), " ", key, `:`, strconv.Quote(val)}, "")
		*tag = TagString(strings.Trim(newTag, " "))
	}
}

// AddMulti adds multiple tags to the TagString
func (tag *TagString) AddMulti(keyVals map[string]string) {
	for key, val := range keyVals {
		tag.Add(key, val)
	}
}

// Remove removes tags identified by key
// Returns an error if the key was not understood
func (tag *TagString) Remove(key string) (err error) {
	re, err := regexp.Compile(`(^| )` + key + `:"[^"]*"`)
	if err == nil && re.MatchString(string(*tag)) {
		newTag := re.ReplaceAllString(string(*tag), "")
		*tag = TagString(newTag)
	}
	return
}

// RemoveMulti removes all the tags identified by the keys in the keys array
// Returns an error if one or more keys was not understood
func (tag *TagString) RemoveMulti(keys []string) (err error) {
	var errorKeys []string
	for _, key := range keys {
		err := tag.Remove(key)
		if err != nil {
			errorKeys = append(errorKeys, key)
		}
	}
	if len(errorKeys) > 0 {
		err = errors.New("The following key(s) were not recognised: " + strings.Join(errorKeys, ", "))
	}
	return
}
