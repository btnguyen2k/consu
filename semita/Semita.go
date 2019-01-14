// Package semita provides utility functions to access data from a hierarchy structure.
//
// Notation:
//   - . (the dot character): path separator
//   - name: access a map's attribute specified by name
//   - i]: access i'th element of a list/array (0-based)
//
// Sample of a path: `employees.[1].first_name`. The dot right before `[]` can be omitted: `employees[1].first_name`.
package semita

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

const (
	// Version defines version number of this package
	Version = "0.1.0"

	// PathSeparator separates path components
	PathSeparator = '.'
)

var (
	patternIndex    = regexp.MustCompile(`^\[(.*?)\]$`)
	patternEndIndex = regexp.MustCompile(`^(.*)(\[.*?\])$`)
)

// SplitPath splits a path into components.
//
// Examples:
//
//   SplitPath("a.b.c.[i].d")     returns []string{"a", "b", "c", "[i]", "d"}
//   SplitPath("a.b.c[i].d")      returns []string{"a", "b", "c", "[i]", "d"}
//   SplitPath("a.b.c.[i].[j].d") returns []string{"a", "b", "c", "[i]", "[j]", "d"}
//   SplitPath("a.b.c[i].[j].d")  returns []string{"a", "b", "c", "[i]", "[j]", "d"}
//   SplitPath("a.b.c[i][j].d")   returns []string{"a", "b", "c", "[i]", "[j]", "d"}
//   SplitPath("a.b.c.[i][j].d")  returns []string{"a", "b", "c", "[i]", "[j]", "d"}
func SplitPath(path string) []string {
	var result []string
	tokens := strings.Split(path, string(PathSeparator))
	for _, token := range tokens {
		var temp []string
		for match := patternEndIndex.FindStringSubmatch(token); len(match) > 0; match = patternEndIndex.FindStringSubmatch(token) {
			temp = append(temp[:0], append([]string{match[2]}, temp[0:]...)...)
			token = match[1]
		}
		if token != "" {
			temp = append(temp[:0], append([]string{token}, temp[0:]...)...)
		}
		result = append(result, temp...)
	}
	return result
}

// Semita struct wraps a underlying data store inside.
type Semita struct {
	Data interface{}
}

// NewSemita creates a new Data and wraps a data store inside it.
func NewSemita(data interface{}) *Semita {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		return &Semita{data}
	}
	return nil
}

func isExportedField(fieldName string) bool {
	return len(fieldName) >= 0 && string(fieldName[0]) == strings.ToUpper(string(fieldName[0]))
}

func getValue(target interface{}, path string) (interface{}, error) {
	if target == nil {
		return nil, nil
	}
	v := reflect.ValueOf(target)
	k := v.Kind()
	if match := patternIndex.FindStringSubmatch(path); len(match) > 0 {
		if k != reflect.Array && k != reflect.Slice {
			return nil, errors.New("required array or slice for path [" + path + "], but input is " + v.Type().String())
		}
		i, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, errors.New("invalid index " + match[1])
		}
		if i < 0 || i > v.Len()-1 {
			return nil, errors.New("array index [" + strconv.Itoa(i) + "] out of range")
		}
		return v.Index(i).Interface(), nil
	}
	switch k {
	case reflect.Map:
		entry := v.MapIndex(reflect.ValueOf(path))
		if entry.Kind() == reflect.Invalid {
			// non-exist index
			return nil, nil
		}
		return entry.Interface(), nil
	case reflect.Struct:
		f := v.FieldByName(path)
		if f.Kind() == reflect.Invalid {
			// non-exist field
			return nil, nil
		}
		if !isExportedField(path) {
			rv := reflect.New(v.Type()).Elem()
			rv.Set(v)
			f = rv.FieldByName(path)
			f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
		}
		return f.Interface(), nil
	}
	return nil, errors.New("required map or struct for path [" + path + "], but input is " + v.Type().String())
}

// GetValue returns a value located at 'path'
func (s *Semita) GetValue(path string) (interface{}, error) {
	paths := SplitPath(path)
	result := s.Data
	var err error
	for _, p := range paths {
		result, err = getValue(result, p)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
