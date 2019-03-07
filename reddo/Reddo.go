/*
Package reddo provides utility functions to convert values using Golang's reflection.

Sample usage:

	package main

	import (
		"fmt"
		"reflect"
		"github.com/btnguyen2k/consu/reddo"
	)

	type Abc struct {
		A int
	}

	type Def struct {
		Abc
		D string
	}

	// convenient method to get value and discarding error
	func getValue(data map[string]interface{}, field string, typ reflect.Type) interface{} {
		v, err := reddo.Convert(data[field], typ)
		if err != nil {
			panic(err)
		}
		return v
	}

	func main() {
		// let's build a 'generic' key-value data store
		data := map[string]interface{}{}
		data["id"] = "1"
		data["name"] = "Thanh Nguyen"
		data["year"] = 2019
		data["abc"] = Abc{A: 103}
		data["def"] = Def{Abc: Abc{A: 1981}, D: "btnguyen2k"}

		// data["id"] and data["year"] both have type interface{}, we would want the correct type
		var id = getValue(data, "id", reddo.TypeString).(string)
		var year = getValue(data, "year", reddo.TypeInt).(int64)
		var yearUint = getValue(data, "year", reddo.TypeUint).(uint64)
		fmt.Printf("Id is %s, year is %d (%d)\n", id, year, yearUint)

		typeAbc := reflect.TypeOf(Abc{})
		typeDef := reflect.TypeOf(Def{})
		var abc = getValue(data, "abc", typeAbc).(Abc)
		var def = getValue(data, "def", typeDef).(Def)
		// special case: struct Def 'inherit' struct Abc, hence Def can be 'cast'-ed to Abc
		var abc2 = getValue(data, "def", typeAbc).(Abc)
		fmt.Println("data.abc       :", abc)  // data.abc       : {103}
		fmt.Println("data.def       :", def)  // data.def       : {{1981} btnguyen2k}
		fmt.Println("data.def as abc:", abc2) // data.def as abc: {1981}

		// special case: convert value to 'time.Time'
		v,_ := reddo.ToTime(1547549353)
		fmt.Println(v)                         // 2019-01-15 17:49:13 +0700 +07
		v,_ = reddo.ToTime("1547549353123")
		fmt.Println(v)                         // 2019-01-15 17:49:13.123 +0700 +07
	}
*/
package reddo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	// Version defines version number of this package
	Version = "0.1.3"
)

var (
	// TypeBool is reflection type of 'bool'
	TypeBool = reflect.TypeOf(true)

	// TypeInt is reflection type of 'int'
	TypeInt = reflect.TypeOf(int(0))

	// TypeUint is reflection type of 'uint'
	TypeUint = reflect.TypeOf(uint(0))

	// TypeFloat is reflection type of 'float64'
	TypeFloat = reflect.TypeOf(float64(0.0))

	// TypeString is reflection type of 'string'
	TypeString = reflect.TypeOf("")

	// TypeUintptr is reflection type of 'uintptr'
	TypeUintptr = reflect.TypeOf(uintptr(0))

	// TypeTime is reflection type of 'time.Time'
	TypeTime = reflect.TypeOf(time.Time{})
)

var zeroTime = time.Time{}

/*
ToBool converts a value to bool.

	- If v is indeed a bool: its value is returned.
	- If v is a number (integer, float or complex): return false if its value is 'zero', true otherwise.
	- If v is a pointer: return false if it is nil, true otherwise.
	- If v is a string: return result from strconv.ParseBool(string).
	- Otherwise, return error

Examples:

	ToBool(true)          returns true,  nil
	ToBool(false)         returns false, nil
	ToBool(0)             returns false, nil
	ToBool(1)             returns true,  nil
	ToBool(-1)            returns true,  nil
	ToBool(0.0)           returns false, nil
	ToBool(1.2)           returns true,  nil
	ToBool(-3.4)          returns true,  nil
	ToBool(1i)            returns true,  nil
	ToBool(-1i)           returns true,  nil
	ToBool(0i)            returns false, nil
	ToBool("true")        returns true,  nil
	ToBool("false")       returns false, nil
	ToBool("blabla")      returns _,     error
	ToBool(struct{}{})    returns _,     error
*/
func ToBool(v interface{}) (bool, error) {
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		return vV.Bool(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vV.Int() != 0, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return vV.Uint() != 0, nil
	case reflect.Float32, reflect.Float64:
		return vV.Float() != 0.0, nil
	case reflect.Complex64, reflect.Complex128:
		return vV.Complex() != 0i, nil
	case reflect.Ptr, reflect.UnsafePointer:
		return vV.Pointer() != 0, nil
	case reflect.String:
		return strconv.ParseBool(vV.String())
	}
	return false, errors.New("cannot convert value [" + vV.String() + "] to bool.")
}

/*
ToFloat converts a value to float64.

	- If v is a bool: return 1.0 if its value is true, 0.0 otherwise.
	- If v is a number (integer or float): return its value as float64.
	- If v is a string: return result from strconv.ParseFloat(string).
	- Otherwise, return error

Examples:

	ToFloat(true)          returns 1.0,  nil
	ToFloat(false)         returns 0.0,  nil
	ToFloat(0)             returns 0.0,  nil
	ToFloat(1.2)           returns 1.2,  nil
	ToFloat("-3.4")        returns -3.4, nil
	ToFloat("blabla")      returns _,     error
	ToFloat(struct{}{})    returns _,     error
*/
func ToFloat(v interface{}) (float64, error) {
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		if vV.Bool() {
			return float64(1.0), nil
		}
		return float64(0.0), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(vV.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(vV.Uint()), nil
	case reflect.Float32:
		// special case for float32
		return strconv.ParseFloat(fmt.Sprint(vV.Interface().(float32)), 64)
	case reflect.Float64:
		return vV.Float(), nil
	case reflect.String:
		return strconv.ParseFloat(vV.String(), 64)
	}
	return float64(0), errors.New("cannot convert value [" + vV.String() + "] to float64.")
}

/*
ToInt converts a value to int64.

	- If v is a number (integer or float): return its value as int64.
	- If v is a bool: return 1 if its value is true, 0 otherwise.
	- If v is a string: return result from strconv.ParseInt(string).
	- Otherwise, return error

Examples:

	ToInt(true)          returns 1,  nil
	ToInt(false)         returns 0,  nil
	ToInt(0)             returns 0,  nil
	ToInt(1.2)           returns 1,  nil
	ToInt("-3")          returns -3, nil
	ToInt("4.5")         returns _,  error
	ToInt("blabla")      returns _,  error
	ToInt(struct{}{})    returns _,  error
*/
func ToInt(v interface{}) (int64, error) {
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		if vV.Bool() {
			return int64(1), nil
		}
		return int64(0), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vV.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(vV.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int64(vV.Float()), nil
	case reflect.String:
		return strconv.ParseInt(vV.String(), 10, 64)
	}
	return int64(0), errors.New("cannot convert value [" + vV.String() + "] to int64.")
}

/*
ToUint converts a value to uint64.

	- If v is a number (integer or float): return its value as uint64.
	- If v is a bool: return 1 if its value is true, 0 otherwise.
	- If v is a string: return result from strconv.ParseUint(string).
	- Otherwise, return error

Examples:

	ToUint(true)          returns 1,  nil
	ToUint(false)         returns 0,  nil
	ToUint(0)             returns 0,  nil
	ToUint(1.2)           returns 1,  nil
	ToUint(-1)            returns 18446744073709551615,  nil // be caution with negative numbers!
	ToUint("-3")          returns _,  error
	ToUint("4.5")         returns _,  error
	ToUint("blabla")      returns _,  error
	ToUint(struct{}{})    returns _,  error
*/
func ToUint(v interface{}) (uint64, error) {
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		if vV.Bool() {
			return uint64(1), nil
		}
		return uint64(0), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(vV.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return vV.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return uint64(vV.Float()), nil
	case reflect.String:
		return strconv.ParseUint(vV.String(), 10, 64)
	}
	return uint64(0), errors.New("cannot convert value [" + vV.String() + "] to uint64.")
}

/*
ToString converts a value to string.

	- If v is a number (integer or float) or bool or string: return its value as string.
	- Otherwise, return string representation of v (fmt.Sprint(v))

Examples:

	ToString(true)          returns "true",   nil
	ToString(false)         returns "false",  nil
	ToString(0)             returns "0",      nil
	ToString(1.2)           returns "1.2",    nil
	ToString("blabla")      returns "blabla", nil
	ToString(struct{}{})    returns "{}",     nil
*/
func ToString(v interface{}) (string, error) {
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(vV.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(vV.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(vV.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(vV.Float(), 'G', -1, 64), nil
	case reflect.String:
		return vV.String(), nil
	}
	return fmt.Sprint(v), nil
}

/*
ToTime converts a value (v) to 'time.Time'.

	- If v is 'time.Time': return v
	- If v is integer: depends on how big v is, treat v as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
	- If v is string and convertible to integer: depends on how big v is, treat v as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
	- Otherwise, return error

Availability: This function is available since v0.1.0.

Examples:

	ToTime(1547549353)         returns Time(Tuesday, January 15, 2019 10:49:13.000 AM GMT), nil
	ToTime("1547549353123")    returns Time(Tuesday, January 15, 2019 10:49:13.123 AM GMT), nil
	ToTime(-1)                 returns _, error
*/
func ToTime(v interface{}) (time.Time, error) {
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Struct:
		if vV.Type().PkgPath() == "time" && vV.Type().Name() == "Time" {
			// same type, just cast it
			return vV.Interface().(time.Time), nil
		}
		return zeroTime, errors.New("value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [time.Time]")
	}
	v, e := ToInt(v)
	if e != nil {
		return zeroTime, errors.New("value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [time.Time]")
	}
	vi := v.(int64)
	if vi >= 0 && vi <= 99999999999 {
		// assume seconds
		return time.Unix(vi, 0), nil
	}
	if vi > 99999999999 && vi <= 99999999999999 {
		// assume milliseconds
		return time.Unix(vi/1000, (vi%1000)*1000000), nil
	}
	if vi > 99999999999999 && vi <= 99999999999999999 {
		// assume microseconds
		return time.Unix(vi/1000000, (vi%1000000)*1000), nil
	}
	if vi > 99999999999999999 {
		// assume nanoseconds
		return time.Unix(0, vi), nil
	}
	return zeroTime, errors.New("value of [" + fmt.Sprint(v) + "] cannot be converted to [time.Time]")
}

/*
ToTimeWithLayout converts a value (v) to 'time.Time'.

ToTimeWithLayout applies the same conversion rules as ToTime does, plus:

	- If v is string and NOT convertible to integer: 'layout' is used to convert the input to 'time.Time'. Error is returned if conversion fails.

Availability: This function is available since v0.1.3.

Examples:

	ToTimeWithLayout(1547549353, _)                                                       returns Time(Tuesday, January 15, 2019 10:49:13.000 AM GMT), nil
	ToTimeWithLayout("1547549353123", _)                                                  returns Time(Tuesday, January 15, 2019 10:49:13.123 AM GMT), nil
	ToTimeWithLayout(-1, _)                                                               returns _, error
	ToTimeWithLayout("January 15, 2019 20:49:13.123", "January 02, 2006 15:04:05.000")    returns Time(Tuesday, January 15, 2019 08:49:13.123 PM GMT), nil
*/
func ToTimeWithLayout(v interface{}, layout string) (time.Time, error) {
	t, e := ToTime(v)
	if e == nil {
		return t, nil
	}
	s, e := ToString(v)
	if e != nil {
		return zeroTime, e
	}
	return time.Parse(layout, s)
}

func isExportedField(fieldName string) bool {
	return len(fieldName) >= 0 && string(fieldName[0]) == strings.ToUpper(string(fieldName[0]))
}

/*
ToStruct converts a value (v) to struct of type specified by (t) (t must be a struct). The output is guaranteed to have the same type as (t).

	- If v is a struct:
		- If v and t are same type, simply cast v to the specified type and return it
		- Otherwise, loop through v's fields. If there is an exported field that is same type as t, return it
		- (since v0.1.1) special case: if t is 'time.Time', return result from ToTime(v)
	- Otherwise, return error

Examples:

	type Abc struct{ Key1 int }
	typeAbc := reflect.TypeOf(Abc{})
	type Def struct {
		Abc
		Key2 string
	}
	typeDef := reflect.TypeOf(Def{})

	ToStruct(Abc{Key1:1}, typeAbc)                      returns Abc{Key1:1},                   nil
	ToStruct(Abc{Key1:1}, typeDef)                      returns _,                             error
	ToStruct(Def{Abc:Abc{Key1:1},Key2:"a"}, typeDef)    returns Def{Abc:Abc{Key1:1},Key2:"a"}, nil
	ToStruct(Def{Abc:Abc{Key1:1},Key2:"a"}, typeAbc)    returns Abc{Key1:1},                   nil
	ToStruct(Abc{Key1:1}, reflect.TypeOf(""))           returns _,                             error
	ToStruct("", typeAbc)                               returns _,                             error
*/
func ToStruct(v interface{}, t reflect.Type) (interface{}, error) {
	if t == nil {
		return nil, errors.New("cannot detect type of target as it is [nil]")
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("target type must be a struct, but received [" + t.String() + "]")
	}
	if t.PkgPath() == "time" && t.Name() == "Time" {
		// special case: convert value to 'time.Time'
		return ToTime(v)
	}

	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Struct:
		if vV.Type().Name() == t.Name() {
			// same type, just cast it
			return vV.Interface(), nil
		}
		// difference type, look into fields
		for i, n := 0, vV.NumField(); i < n; i++ {
			f := vV.Field(i)
			fn := vV.Type().Field(i).Name
			if f.Kind() == reflect.Struct && isExportedField(fn) {
				if f.Type().Name() == t.Name() {
					return f.Interface(), nil
				}
			}
		}
	}
	return nil, errors.New("value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [" + t.String() + "]")
}

/*
ToSlice converts a value (v) to slice of type specified by (t) (t can be a slice or array, or an element of slice/array).
The output is guaranteed to have the same type as (t).

	- If v is an array or slice: convert each element of v to the correct type (specified by t), put them into a slice, and finally return it.
	- Otherwise, return error

Notes:

	- Array/slice is converted to slice
	- Element type can be converted too, for example: []int can be converted to []string

Examples:

	ToSlice([]bool{true,false}, reflect.TypeOf([0]int{}))    returns []int{1,0},               nil
	ToSlice([3]int{-1,0,1}, reflect.TypeOf([]string{""}))    returns []string{"-1","0","1"},   nil
	ToSlice([]bool{true,false}, TypeString)                  returns []string{"true","false"}, nil
*/
func ToSlice(v interface{}, t reflect.Type) (interface{}, error) {
	if t == nil {
		return nil, errors.New("cannot detect type of target as it is [nil]")
	}
	if t.Kind() != reflect.Array && t.Kind() != reflect.Slice {
		return ToSlice(v, reflect.SliceOf(t))
	}
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Array, reflect.Slice:
		elementType := t.Elem()                                        // type of slice element
		slice := reflect.MakeSlice(reflect.SliceOf(elementType), 0, 0) // create an empty slice
		for i, n := 0, vV.Len(); i < n; i++ {
			e, err := Convert(vV.Index(i).Interface(), elementType)
			if err == nil {
				slice = reflect.Append(slice, reflect.ValueOf(e).Convert(elementType))
			} else {
				return nil, err
			}
		}
		return slice.Interface(), nil
	}
	return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + t.String() + "]")
}

/*
ToMap converts a value (v) to map where types of key & value are specified by (t) (t must be a map).
The output is guaranteed to have the same type as (t).

	- If v is a map: convert each element {key:value} of v to the correct type (specified by t), put them into a map, and finally return it.
	- Otherwise, return error

Notes:

	- Element type can be converted too, for example: map[int]int can be converted to map[string]string

Examples:

	ToMap(map[string]bool{"a":true,"b":false}, reflect.TypeOf(map[string]int{}))    returns map[string]int{"a":1,"b":0"}, nil
*/
func ToMap(v interface{}, t reflect.Type) (interface{}, error) {
	if t == nil {
		return nil, errors.New("cannot detect type of target as it is [nil]")
	}
	if t.Kind() != reflect.Map {
		return nil, errors.New("target type must be a map, but received [" + t.String() + "]")
	}
	vV := reflect.ValueOf(v)
	if vV.Kind() == reflect.Map {
		keyType := t.Key()    // type of map's key
		valueType := t.Elem() // type of map's value
		m := reflect.MakeMap(reflect.MapOf(keyType, valueType))
		for _, k := range vV.MapKeys() {
			key, err := Convert(k.Interface(), keyType)
			if err != nil {
				return nil, err
			}
			value, err := Convert(vV.MapIndex(k).Interface(), valueType)
			if err != nil {
				return nil, err
			}
			m.SetMapIndex(reflect.ValueOf(key).Convert(keyType), reflect.ValueOf(value).Convert(valueType))
		}
		return m.Interface(), nil
	}
	return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + t.String() + "]")
}

/*
ToPointer converts a value (v) to pointer of type specified by (t) (t must be a pointer).
The output is guaranteed to have the same type as (t).

Example 1:

	a := float64(1.23)
	zero := int32(0)
	output, err := ToPointer(&a, reflect.TypeOf(&zero))
	// here err should be nil
	if err != nil {
		panic("Something wrong!")
	}
	// we now successfully converted *float644 to *int32
	i32 := *output.(*interface{}) // note: type of output is *interface{}
	fmt.Println(i32.(int32))      // i32 is safe to type asserted .(int32)

Example 2:

	type Abc struct {
		A int
	}
	type Def struct {
		Abc
		D string
	}
	z := Abc{}
	a := Def{Abc: Abc{1}, D: "2"}
	output, err := ToPointer(&a, reflect.TypeOf(&z))
	// here err should be nil
	if err != nil {
		panic("Something wrong!")
	}
	// we now successfully converted *Def to *Abc
	abc := *output.(*interface{}) // note: type of output is *interface{}
	fmt.Println(abc.(Abc))        // i32 is safe to type asserted .(Abc)
*/
func ToPointer(v interface{}, t reflect.Type) (interface{}, error) {
	if t == nil {
		return nil, errors.New("cannot detect type of target as it is [nil]")
	}
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("target type must be a pointer, but received [" + t.String() + "]")
	}
	vV := reflect.ValueOf(v)
	if vV.Kind() == reflect.Ptr {
		v, err := Convert(vV.Elem().Interface(), t.Elem())
		if err != nil {
			return nil, err
		}
		x := reflect.ValueOf(v).Convert(t.Elem()).Interface()
		return &x, nil
	}
	return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + t.String() + "]")
}

/*
Convert converts a value (v) to specified type (t):

	- If t is a bool: see ToBool(...)
	- If t is an integer (int, int8, int16, int32, int64): see ToInt(...)
	- If t is an unsigned-integer (uint, uint8, uint16, uint32, uint64, uintptr): see ToUint(...)
	- If t is a float (float32, float64): see ToFloat(...)
	- If t is a string: see ToString(...)
	- If t is a struct: see ToStruct(...)
	- If t is an array or a slice: see ToSlice(v...)
	- If t is a map: see ToMap(...)
	- If t is a pointer: see ToPointer(...)
	- (special case) If t is nil: this function returns (v, nil)
*/
func Convert(v interface{}, t reflect.Type) (interface{}, error) {
	if t == nil {
		return v, nil
	}
	if v == nil {
		return nil, errors.New("cannot convert: nil to " + t.String())
	}
	k := t.Kind()
	switch k {
	case reflect.Bool:
		return ToBool(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ToInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return ToUint(v)
	case reflect.Float32, reflect.Float64:
		return ToFloat(v)
	case reflect.String:
		return ToString(v)
	case reflect.Struct:
		return ToStruct(v, t)
	case reflect.Array, reflect.Slice:
		return ToSlice(v, t)
	case reflect.Map:
		return ToMap(v, t)
	case reflect.Ptr:
		return ToPointer(v, t)
	case reflect.Interface:
		// special case
		return v, nil
	}
	return v, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + t.String() + "]")
}
