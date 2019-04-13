/*
Package semita provides utility functions to access data from a hierarchy structure.

A "path" is used to specify the location of item in the hierarchy data. Sample of a path

	"Employees.[1].first_name"

where:

	- "." (the dot character): path separator
	- "Name": access attribute of a map/struct specified by "Name"
	- "[i]": access i'th element of a slice/array (0-based)
	- The dot right before "[]" can be omitted: "Employees[1].first_name" is equivalent to "Employees.[1].first_name".

Notes:

	- Supported nested arrays, slices, maps and structs.
	- Struct's un-exported fields can be read, but not written.
	- Unaddressable structs and arrays are read-only.

Example: (more examples at https://github.com/btnguyen2k/consu/tree/master/semita/examples)

	package main

	import (
		"encoding/json"
		"fmt"
		"github.com/btnguyen2k/consu/reddo"
		"github.com/btnguyen2k/consu/semita"
	)

	func main() {
		fmt.Println("-========== Semina demo ==========-")
		data := map[string]interface{}{
			"name": map[string]interface{}{
				"first": "Thanh",
				"last":  "ngn",
			},
			"yob":   1981,
			"alias": []string{"btnguyen2k", "thanhnb"},
		}
		s := semita.NewSemita(data)
		var err error
		var v interface{}

		// current data tree: {"alias":["btnguyen2k","thanhnb"],"name":{"first":"Thanh","last":"ngn"},"yob":1981}
		tree := s.Unwrap()
		js, _ := json.Marshal(tree)
		fmt.Println("Data tree:", string(js))

		// get nested value
		v, err = s.GetValue("name.first") // v should be "Thanh"
		if err == nil {
			fmt.Println("Firstname:", v)
		} else {
			fmt.Println("Error:", err)
		}

		// set nested value
		err = s.SetValue("name.last", "Nguyen") // v should be "Nguyen" (instead of "ngn")
		if err == nil {
			v, err = s.GetValue("name.last")
			if err == nil {
				fmt.Println("Lastname:", v)
			} else {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Error:", err)
		}

		// get a value and its type
		yob, err := s.GetValue("yob") // yob should be int(1981)
		if err == nil {
			fmt.Println("YOB:", yob.(int))
		} else {
			fmt.Println("Error:", err)
		}

		// get a value and type
		yob, err = s.GetValueOfType("yob", reddo.TypeUint) // yob should be uint64(1981)
		if err == nil {
			fmt.Println("YOB:", yob.(uint64)) // all uint types are returned as uint64
		} else {
			fmt.Println("Error:", err)
		}

		// append new item to end of slice
		err = s.SetValue("alias[]", "another")
		if err == nil {
			// either alias[2] or alias.[2] is accepted
			alias, err := s.GetValue("alias.[2]") // alias should be "another"
			if err == nil {
				fmt.Println("New Alias:", alias)
			} else {
				fmt.Println("Error:", err)
			}

			allAlias, err := s.GetValue("alias") // allAlias should be ["btnguyen2k","thanhnb","another"]
			if err == nil {
				fmt.Println("All Alias:", allAlias)
			} else {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Error:", err)
		}

		// create missing nodes along the path
		err = s.SetValue("a.b[].c.d", true)
		if err == nil {
			// missing nodes should be created
			// data tree should be: {"a":{"b":[{"c":{"d":true}}]},"alias":["btnguyen2k","thanhnb","another"],"name":{"first":"Thanh","last":"Nguyen"},"yob":1981}
			tree := s.Unwrap()
			js, _ := json.Marshal(tree)
			fmt.Println("Data tree:", string(js))
		} else {
			fmt.Println("Error:", err)
		}
	}
*/
package semita

import (
	"errors"
	"github.com/btnguyen2k/consu/reddo"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	// Version defines version number of this package
	Version = "0.1.4.1"

	// PathSeparator separates path components
	PathSeparator = '.'
)

var (
	patternIndex    = regexp.MustCompile(`^\[(.*?)\]$`)
	patternEndIndex = regexp.MustCompile(`^(.*)(\[.*?\])$`)
)

// concreteValue returns concrete value of target 't'.
func concreteValue(t interface{}) reflect.Value {
	v := reflect.ValueOf(t)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

/*
GetTypeOfMapKey returns type of map 'm''s key.

	- if 'm' is a map (or pointer to a map): type of map's key is returned
	- otherwise, nil is returned
*/
func GetTypeOfMapKey(m interface{}) reflect.Type {
	v := concreteValue(m)
	if v.Kind() == reflect.Map {
		return v.Type().Key()
	}
	return nil
}

/*
GetTypeOfElement returns type of element of target 't'.

	- if 't' is an array, slice, map or channel (or pointer to an array, slice, map or channel): element type is returned
	- otherwise, t's type is return
*/
func GetTypeOfElement(t interface{}) reflect.Type {
	v := concreteValue(t)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array || v.Kind() == reflect.Map || v.Kind() == reflect.Chan {
		return v.Type().Elem()
	}
	return v.Type()
}

/*
GetTypeOfStructAttibute returns type of a struct attribute.

	- if 's' is a struct (or pointer to a struct): type of attribute 'attr' is returned
	- if 's' is not a struct (or pointer to a struct) or attribute 'attr' does not exist: nil is returned
*/
func GetTypeOfStructAttibute(s interface{}, attr string) reflect.Type {
	v := concreteValue(s)
	if v.Kind() != reflect.Struct {
		return nil
	}
	f := v.FieldByName(attr)
	if f.Kind() != reflect.Invalid {
		return f.Type()
	}
	return nil
}

/*
CreateZero create 'zero' value of specified type

	- if 't' is primitive type (bool, ints, uints, floats, complexes, string, uintptr and unsafe-pointer): 'zero' value is created via reflect.Zero(t)
	- if 't' is array or slice: returns empty slice of type 't'
	- if 't' is map: returns empty map of type 't'
	- if 't' is struct: returns empty struct of type 't'
	- otherwise, return empty 'reflect.Value'
*/
func CreateZero(t reflect.Type) reflect.Value {
	if t == nil {
		return reflect.Value{}
	}
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.UnsafePointer:
		// primitive types --> 'zero' primitive
		return reflect.Zero(t)
	case reflect.Slice, reflect.Array:
		// slice/array --> empty slice
		elType := t.Elem()
		return reflect.MakeSlice(reflect.SliceOf(elType), 0, 0)
	case reflect.Map:
		// map --> 'zero' map
		return reflect.MakeMap(t)
	case reflect.Struct:
		// struct --> 'zero' struct
		return reflect.New(t).Elem()
	default:
		// Chan
		// Func
		// Interface
		// Ptr
	}
	return reflect.Value{}
}

/*
SplitPath splits a path into components.

Examples:

	SplitPath("a.b.c.[i].d")     returns []string{"a", "b", "c", "[i]", "d"}
	SplitPath("a.b.c[i].d")      returns []string{"a", "b", "c", "[i]", "d"}
	SplitPath("a.b.c.[i].[j].d") returns []string{"a", "b", "c", "[i]", "[j]", "d"}
	SplitPath("a.b.c[i].[j].d")  returns []string{"a", "b", "c", "[i]", "[j]", "d"}
	SplitPath("a.b.c[i][j].d")   returns []string{"a", "b", "c", "[i]", "[j]", "d"}
	SplitPath("a.b.c.[i][j].d")  returns []string{"a", "b", "c", "[i]", "[j]", "d"}
*/
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

/*
Semita struct wraps a underlying data store inside.
*/
type Semita struct {
	root *node
}

/*
NewSemita wraps the supplied 'data' inside a Semita instance and returns pointer to the Semita instance.
data must be either array, slice, map or struct (or pointer fo them).
*/
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
		case reflect.Interface:
			return NewSemita(v.Elem().Interface())
		}
	}
	return nil
}

/*----------------------------------------------------------------------*/

/*
Unwrap returns the underlying data
*/
func (s *Semita) Unwrap() interface{} {
	// if s.root == nil || s.root.value.Kind() == reflect.Invalid {
	// 	return nil
	// }
	return s.root.unwrap()
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

/*
GetValue returns a value located at 'path'.

Notes:

	- Wrapped data must be either a struct, map, array or slice
	- map's keys must be strings
	- Nested structure is supported (e.g. array inside a map, inside a struct, inside a slice, etc)
	- Getting value of struct's unexported field is supported
	- If index is out-of-bound, (nil, nil) is returned

Example:

	  data := map[string]interface{}{
		"Name": "Monster Corp.",
		"Year": 2003,
		"Employees": []map[string]interface{}{
		  {
			"first_name": "Mike",
			"last_name" : "Wazowski",
			"email"     : "mike.wazowski@monster.com",
			"age"       : 29,
			"options"   : map[string]interface{}{
			  "work_hours": []int{9, 10, 11, 12, 13, 14, 15, 16},
			  "overtime"  : false,
			},
		  },
		  {
			"first_name": "Sulley",
			"last_name" : "Sullivan",
			"email"     : "sulley.sullivan@monster.com",
			"age"       : 30,
			"options"   : map[string]interface{}{
			  "work_hours": []int{13, 14, 15, 16, 17, 18, 19, 20},
			  "overtime"  :   true,
			},
		  },
		},
	  }
	  s := NewSetima(data)
	  s.GetValue("Name")              // "Monster Corp."
	  s.GetValue("Employees[0].age")  // 29
*/
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

/*
GetValueOfType retrieves value located at 'path', converts the value to target type and returns it.

Notes:

	- Wrapped data must be either a struct, map, array or slice
	- map's keys must be strings
	- Nested structure is supported (e.g. array inside a map, inside a struct, inside a slice, etc)
	- Getting value of struct's unexported field is supported
	- If index is out-of-bound, (nil, nil) is returned

Example:

	data := map[string]interface{}{
		"Name": "Monster Corp.",
		"Year": 2003,
		"Employees": []map[string]interface{}{
			{
				"first_name": "Mike",
				"last_name" : "Wazowski",
				"email"     : "mike.wazowski@monster.com",
				"age"       : 29,
				"options"   : map[string]interface{}{
			  	"work_hours": []int{9, 10, 11, 12, 13, 14, 15, 16},
			  	"overtime"  : false,
			},
			},
			{
				"first_name": "Sulley",
				"last_name" : "Sullivan",
				"email"     : "sulley.sullivan@monster.com",
				"age"       : 30,
				"options"   : map[string]interface{}{
					"work_hours": []int{13, 14, 15, 16, 17, 18, 19, 20},
					"overtime"  :   true,
				},
			},
		},
	}
	s := NewSetima(data)
	var Name string = s.GetValueOfType("Name", reddo.TypeString).(string)            // "Monster Corp."
	var age int64   = s.GetValueOfType("Employees[0].age", reddo.TypeInt).(int64)    // 29
*/
func (s *Semita) GetValueOfType(path string, typ reflect.Type) (interface{}, error) {
	v, e := s.GetValue(path)
	if v == nil || e != nil {
		return v, e
	}
	return reddo.Convert(v, typ)
}

/*
GetTime returns a 'time.Time' located at 'path'.

Availability: This function is available since v0.1.2.

Notes:

	- Same rules/restrictions as of GetValue function.
	- If the value located at 'path' is 'time.Time': return it.
	- If the value is integer: depends on how big it is, treat the value as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
	- If the value is string and convertible to integer: depends on how big it is, treat the value as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
	- Otherwise, return error

Example:

	data := map[string]interface{}{
		"ValueInt"    : 1547549353,
		"ValueString" : "1547549353123",
		"ValueInvalid": -1,
	}
	s := NewSetima(data)
	s.GetTime("ValueInt")        returns Time(Tuesday, January 15, 2019 10:49:13.000 AM GMT), nil
	s.GetTime("ValueString")     returns Time(Tuesday, January 15, 2019 10:49:13.123 AM GMT), nil
	s.GetTime("ValueInvalid")    returns _, error
*/
func (s *Semita) GetTime(path string) (time.Time, error) {
	v, e := s.GetValue(path)
	if v == nil || e != nil {
		return time.Time{}, e
	}
	return reddo.ToTime(v)
}

/*
GetTimeWithLayout returns a 'time.Time' located at 'path'.

Availability: This function is available since v0.1.2.

Notes:

	- Same rules/restrictions as of GetTime function, plus:
	- If the value is string and NOT convertible to integer: 'layout' is used to convert the value to 'time.Time'. Error is returned if conversion fails.

Example:

	data := map[string]interface{}{
		"ValueInt"    : 1547549353,
		"ValueString" : "1547549353123",
		"ValueInvalid": -1,
		"ValueDateStr": "January 15, 2019 20:49:13.123",
	}
	s := NewSetima(data)
	s.GetTimeWithLayout("ValueInt", _)                                      returns Time(Tuesday, January 15, 2019 10:49:13.000 AM GMT), nil
	s.GetTimeWithLayout("ValueString", _)                                   returns Time(Tuesday, January 15, 2019 10:49:13.123 AM GMT), nil
	s.GetTimeWithLayout("ValueInvalid", _)                                  returns _, error
	s.GetTimeWithLayout("ValueDateStr", "January 02, 2006 15:04:05.000")    returns Time(Tuesday, January 15, 2019 08:49:13.123 PM GMT), nil
*/
func (s *Semita) GetTimeWithLayout(path, layout string) (time.Time, error) {
	zeroTime := time.Time{}
	v, e := s.GetValue(path)
	if e != nil {
		return zeroTime, e
	}
	t, e := reddo.ToTime(v)
	if e == nil {
		return t, nil
	}
	str, e := reddo.ToString(v)
	if e != nil {
		return time.Time{}, e
	}
	return time.Parse(layout, str)
}

/*
SetValue sets a value to position specified by 'path'.

Notes:

	- Wrapped data must be either a struct, map, array or slice
	- map's keys must be strings
	- If 'path' points to a map's key, the key must be exported
	- Nested structure is supported (e.g. array inside a map, inside a struct, inside a slice, etc)
	- If child nodes along the path does not exist, this function will create them
	- If index is out-of-bound, this function returns error

Example:

	data := map[string]interface{}{}
	s := NewSetima(data)
	s.SetValue("Name", "Monster Corp.") // data is now {"Name":"Monster Corp."}
	s.SetValue("Year", 2013)            // data is now {"Name":"Monster Corp.", "Year":2013}
	s.SetValue("employees[0].age", 29)  // data is now {"Name":"Monster Corp.", "Year":2013, "employees":[{"age":29}]}
*/
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
			_newNode, err = prevCursor.createChild(index, nextIndex)
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
