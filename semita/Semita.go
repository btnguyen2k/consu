// Package semita provides utility functions to access data from a hierarchy structure.
//
// Notation:
//   - . (the dot character): path separator
//   - Name: access attribute of a map/struct specified by 'Name'
//   - [i]: access i'th element of a slice/array (0-based)
//
// Sample of a path: `Employees.[1].first_name`. The dot right before `[]` can be omitted: `Employees[1].first_name`.
package semita

import (
	"errors"
	"github.com/btnguyen2k/consu/reddo"
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
	root *node
}

// NewSemita wraps the supplied 'data' inside a Semita instance and returns pointer to the Semita instance.
// data must be either array, slice, map or struct (or pointer fo them).
func NewSemita(data interface{}) *Semita {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		v := reflect.ValueOf(data)
		return &Semita{
			&node{
				prev:     nil,
				prevType: nil,
				key:      "",
				value:    v,
			},
		}
	case reflect.Ptr:
		switch v.Elem().Kind() {
		case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
			p := unsafe.Pointer(v.Pointer())
			r := reflect.NewAt(v.Elem().Type(), p)
			v := r.Elem()
			return &Semita{
				&node{
					prev:     nil,
					prevType: nil,
					key:      "",
					value:    v,
				},
			}
			// case reflect.Interface:
			// 	vi := reflect.ValueOf(v.Elem().Interface())
			// 	switch vi.Kind() {
			// 	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
			// 		p := unsafe.Pointer(v.Pointer())
			// 		r := reflect.NewAt(vi.Type(), p)
			// 		v := r.Elem()
			// 		return &Semita{
			// 			&node{
			// 				prev:     nil,
			// 				prevType: nil,
			// 				key:      "",
			// 				value:    v,
			// 			},
			// 		}
			// 	}
		}
	}
	return nil
}

/*----------------------------------------------------------------------*/

// Unwrap returns the underlying data
func (s *Semita) Unwrap() interface{} {
	// if s.root == nil || s.root.value.Kind() == reflect.Invalid {
	// 	return nil
	// }
	return s.root.value.Interface()
}

// seek seeks along the path and returns (prevNode, node, error)
func (s *Semita) seek(path string) (prevCursor *node, cursor *node, err error) {
	paths := SplitPath(path)
	prevCursor, cursor = s.root, s.root
	for _, index := range paths {
		cursor, err = cursor.next(index)
		if err != nil {
			return nil, nil, err
		}
		if cursor == nil {
			return prevCursor, nil, nil
		}
		prevCursor = cursor
	}
	return cursor.prev, cursor, nil
}

// GetValue returns a value located at 'path'.
//
// Notes:
//
//   - Wrapped data must be either a struct, map, array or slice
//   - map's keys must be strings
//   - Nested structure is supported (e.g. array inside a map, inside a struct, inside a slice, etc)
//   - Getting value of struct's unexported field is supported
//   - If index is out-of-bound, (nil, nil) is returned
//
// Example:
//
//   data := map[string]interface{}{
//     "Name": "Monster Corp.",
//     "Year": 2003,
//     "Employees": []map[string]interface{}{
//       {
//         "first_name": "Mike",
//         "last_name" : "Wazowski",
//         "email"     : "mike.wazowski@monster.com",
//         "age"       : 29,
//         "options"   : map[string]interface{}{
//           "work_hours": []int{9, 10, 11, 12, 13, 14, 15, 16},
//           "overtime"  : false,
//         },
//       },
//       {
//         "first_name": "Sulley",
//         "last_name" : "Sullivan",
//         "email"     : "sulley.sullivan@monster.com",
//         "age"       : 30,
//         "options"   : map[string]interface{}{
//           "work_hours": []int{13, 14, 15, 16, 17, 18, 19, 20},
//           "overtime"  :   true,
//         },
//       },
//     },
//   }
//   s := NewSetima(data)
//   s.GetValue("Name")              // "Monster Corp."
//   s.GetValue("Employees[0].age")  // 29
func (s *Semita) GetValue(path string) (interface{}, error) {
	_, cursor, err := s.seek(path)
	if err != nil {
		return nil, err
	}
	if cursor == nil {
		return nil, nil
	}
	return cursor.unwrap(), nil
}

// GetValueOfType retrieves value located at 'path', converts the value to target's type and returns it.
//
// Notes:
//
//   - Wrapped data must be either a struct, map, array or slice
//   - map's keys must be strings
//   - Nested structure is supported (e.g. array inside a map, inside a struct, inside a slice, etc)
//   - Getting value of struct's unexported field is supported
//   - If index is out-of-bound, (nil, nil) is returned
//
// Example:
//
//   data := map[string]interface{}{
//     "Name": "Monster Corp.",
//     "Year": 2003,
//     "Employees": []map[string]interface{}{
//       {
//         "first_name": "Mike",
//         "last_name" : "Wazowski",
//         "email"     : "mike.wazowski@monster.com",
//         "age"       : 29,
//         "options"   : map[string]interface{}{
//           "work_hours": []int{9, 10, 11, 12, 13, 14, 15, 16},
//           "overtime"  : false,
//         },
//       },
//       {
//         "first_name": "Sulley",
//         "last_name" : "Sullivan",
//         "email"     : "sulley.sullivan@monster.com",
//         "age"       : 30,
//         "options"   : map[string]interface{}{
//           "work_hours": []int{13, 14, 15, 16, 17, 18, 19, 20},
//           "overtime"  :   true,
//         },
//       },
//     },
//   }
//   s := NewSetima(data)
//   var Name string = s.GetValueOfType("Name", reddo.ZeroString).(string)          // "Monster Corp."
//   var age int64   = s.GetValueOfType("Employees[0].age", reddo.ZeroInt).(int64)  // 29
func (s *Semita) GetValueOfType(path string, target interface{}) (interface{}, error) {
	v, e := s.GetValue(path)
	if v == nil || e != nil {
		return v, e
	}
	return reddo.Convert(v, target)
}

// SetValue sets a value to position specified by 'path'.
//
// Notes:
//
//   - Wrapped data must be either a struct, map, array or slice
//   - map's keys must be strings
//   - If 'path' points to a map's key, the key must be exported
//   - Nested structure is supported (e.g. array inside a map, inside a struct, inside a slice, etc)
//   - If child nodes along the path does not exist, this function will create them
//   - If index is out-of-bound, this function returns error
//
// Example:
//
//   data := map[string]interface{}{}
//   s := NewSetima(data)
//   s.SetValue("Name", "Monster Corp.") // data is now {"Name":"Monster Corp."}
//   s.SetValue("Year", 2013)            // data is now {"Name":"Monster Corp.", "Year":2013}
//   s.SetValue("employees[0].age", 29)  // data is now {"Name":"Monster Corp.", "Year":2013, "employees":[{"age":29}]}
func (s *Semita) SetValue(path string, value interface{}) error {
	paths := SplitPath(path)
	var pathSoFar string
	// "seek" to the correct position
	for i, index := range paths[0 : len(paths)-1] {
		if i > 0 {
			pathSoFar = string(append([]byte(pathSoFar), PathSeparator))
		}
		pathSoFar += index
		prevCursor, cursor, err := s.seek(pathSoFar)
		if err != nil {
			return errors.New("error while getting value at path: " + pathSoFar)
		}
		if cursor == nil {
			// create node along the way
			var _newNode *node
			var err error
			nextIndex := paths[i+1]
			if patternIndex.MatchString(nextIndex) {
				_newNode, err = prevCursor.createChildSlice(index)
			} else {
				_newNode, err = prevCursor.createChildMap(index)
			}
			if err != nil {
				return errors.New("error while creating node at at path: " + pathSoFar)
			}
			prevCursor = _newNode.prev
		}
		if index == "[]" {
			// special case
			l := prevCursor.elem().Len()
			pathSoFar = pathSoFar[0:len(pathSoFar)-2] + "[" + strconv.Itoa(l-1) + "]"
		}
	}
	_, cursor, err := s.seek(pathSoFar)
	if err != nil || cursor == nil || cursor.unwrap() == nil {
		return errors.New("path not found: " + pathSoFar)
	}
	index := paths[len(paths)-1]
	_, err = cursor.setValue(index, reflect.ValueOf(value))
	return err
}
