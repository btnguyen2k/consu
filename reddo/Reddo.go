// Package reddo provides utility functions to convert values using Golang's reflect.
//
// Sample usage:
//
//   package main
//
//   import (
//     "fmt"
//     "github.com/btnguyen2k/consu/reddo"
//   )
//
//   type Abc struct {
//     A int
//   }
//
//   type Def struct {
//     Abc
//     D string
//   }
//
//   // convenient method to get value and discarding error
//   func getValue(data map[string]interface{}, field string, zero interface{}) interface{} {
//     v, err := reddo.Convert(data[field], zero)
//     if err != nil {
//       panic(err)
//     }
//     return v
//   }
//
//   func main() {
//     // let's build a 'generic' key-value data store
//     data := map[string]interface{}{}
//     data["id"] = "1"
//     data["name"] = "Thanh Nguyen"
//     data["year"] = 2019
//     data["abc"] = Abc{A: 103}
//     data["def"] = Def{Abc: Abc{A: 1981}, D: "btnguyen2k"}
//
//     // data["id"] and data["year"] both have type interface{}, we would want the correct type
//     var id = getValue(data, "id", reddo.ZeroString).(string)
//     var year = getValue(data, "year", reddo.ZeroInt).(int64)
//     var yearUint = getValue(data, "year", reddo.ZeroUint).(uint64)
//     fmt.Printf("Id is %s, year is %d (%d)\n", id, year, yearUint)
//
//     // we need a 'zero' value of the target type to retrieve the correct value & type from out data store
//     zeroAbc := Abc{}
//     zeroDef := Def{}
//     var abc = getValue(data, "abc", zeroAbc).(Abc)
//     var def = getValue(data, "def", zeroDef).(Def)
//     // special case: struct Def 'inherit' struct Abc, hence Def can be 'cast'-ed to Abc
//     var abc2 = getValue(data, "def", zeroAbc).(Abc)
//     fmt.Println("data.abc       :", abc)  // data.abc       : {103}
//     fmt.Println("data.def       :", def)  // data.def       : {{1981} btnguyen2k}
//     fmt.Println("data.def as abc:", abc2) // data.def as abc: {1981}
//
//     // Special case: convert value to 'time.Time'
//     v,_ := reddo.ToTime(1547549353)
//     fmt.Println(v)                         // 2019-01-15 17:49:13 +0700 +07
//     v,_ = reddo.ToTime("1547549353123")
//     fmt.Println(v)                         // 2019-01-15 17:49:13.123 +0700 +07
//   }
package reddo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

const (
	// Version defines version number of this package
	Version = "0.1.1"

	// ZeroBool defines 'zero' value of type bool
	ZeroBool = false

	// ZeroFloat defines 'zero' value of type float64
	ZeroFloat = float64(0.0)
	// ZeroFloat64 defines 'zero' value of type float64
	ZeroFloat64 = float64(0.0)

	// ZeroInt defines 'zero' value of type int64
	ZeroInt = int64(0)
	// ZeroInt64 defines 'zero' value of type int64
	ZeroInt64 = int64(0)

	// ZeroUint defines 'zero' value of type uint64
	ZeroUint = uint64(0)
	// ZeroUintptr defines 'zero' value of type uintptr
	ZeroUintptr = uintptr(0)

	// ZeroString defines 'zero' value of type string
	ZeroString = ""
)

var (
	// ZeroTime defines 'zero' value of type 'time.Time'
	ZeroTime = *new(time.Time)
)

// ToBool converts a value to bool. The output is guaranteed to ad-here to type assertion .(bool)
//
//   - If v is indeed a bool: its value is returned.
//   - If v is a number (integer, float or complex): return false if its value is 'zero', true otherwise.
//   - If v is a pointer: return false if it is nil, true otherwise.
//   - If v is a string: return result from strconv.ParseBool(string).
//   - Otherwise, return error
//
// Examples:
//
//    ToBool(true)       returns true,  nil
//    ToBool(false)      returns false, nil
//    ToBool(0)          returns false, nil
//    ToBool(1)          returns true,  nil
//    ToBool(-1)         returns true,  nil
//    ToBool(0.0)        returns false, nil
//    ToBool(1.2)        returns true,  nil
//    ToBool(-3.4)       returns true,  nil
//    ToBool(1i)         returns true,  nil
//    ToBool(-1i)        returns true,  nil
//    ToBool(0i)         returns false, nil
//    ToBool("true")     returns true,  nil
//    ToBool("false")    returns false, nil
//    ToBool("blabla")   returns _,     error
//    ToBool(struct{}{}) returns _,     error
func ToBool(v interface{}) (interface{}, error) {
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

// ToFloat converts a value to float64. The output is guaranteed to ad-here to type assertion .(float64)
//
//   - If v is a bool: return 1.0 if its value is true, 0.0 otherwise.
//   - If v is a number (integer or float): return its value as float64.
//   - If v is a string: return result from strconv.ParseFloat(string).
//   - Otherwise, return error
//
// Examples:
//
//   ToFloat(true)       returns 1.0,  nil
//   ToFloat(false)      returns 0.0,  nil
//   ToFloat(0)          returns 0.0,  nil
//   ToFloat(1.2)        returns 1.2,  nil
//   ToFloat("-3.4")     returns -3.4, nil
//   ToFloat("blabla")   returns _,     error
//   ToFloat(struct{}{}) returns _,     error
func ToFloat(v interface{}) (interface{}, error) {
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

// ToInt converts a value to int64. The output is guaranteed to ad-here to type assertion .(int64)
//
//   - If v is a number (integer or float): return its value as int64.
//   - If v is a bool: return 1 if its value is true, 0 otherwise.
//   - If v is a string: return result from strconv.ParseInt(string).
//   - Otherwise, return error
//
// Examples:
//
//   ToInt(true)       returns 1,  nil
//   ToInt(false)      returns 0,  nil
//   ToInt(0)          returns 0,  nil
//   ToInt(1.2)        returns 1,  nil
//   ToInt("-3")       returns -3, nil
//   ToInt("4.5")      returns _,  error
//   ToInt("blabla")   returns _,  error
//   ToInt(struct{}{}) returns _,  error
func ToInt(v interface{}) (interface{}, error) {
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

// ToUint converts a value to uint64. The output is guaranteed to ad-here to type assertion .(uint64)
//
//   - If v is a number (integer or float): return its value as uint64.
//   - If v is a bool: return 1 if its value is true, 0 otherwise.
//   - If v is a string: return result from strconv.ParseUint(string).
//   - Otherwise, return error
//
// Examples:
//
//   ToUint(true)       returns 1,  nil
//   ToUint(false)      returns 0,  nil
//   ToUint(0)          returns 0,  nil
//   ToUint(1.2)        returns 1,  nil
//   ToUint(-1)         returns 18446744073709551615,  nil /* be caution with negative numbers! */
//   ToUint("-3")       returns _,  error
//   ToUint("4.5")      returns _,  error
//   ToUint("blabla")   returns _,  error
//   ToUint(struct{}{}) returns _,  error
func ToUint(v interface{}) (interface{}, error) {
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

// ToString converts a value to string. The output is guaranteed to ad-here to type assertion .(string)
//
//   - If v is a number (integer or float) or bool or string: return its value as string.
//   - Otherwise, return string representation of v (fmt.Sprint(v))
//
// Examples:
//
//   ToString(true)       returns "true",   nil
//   ToString(false)      returns "false",  nil
//   ToString(0)          returns "0",      nil
//   ToString(1.2)        returns "1.2",    nil
//   ToString("blabla")   returns "blabla", nil
//   ToString(struct{}{}) returns "{}",     nil
func ToString(v interface{}) (interface{}, error) {
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

// ToTime converts a value (v) to 'time.Time'.
//
//   - If v is 'time.Time': return v
//   - If v is integer: depends on how big is v, treat v as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
//   - If v is string and convertable to integer: depends on how big is v, treat v as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
//   - Otherwise, return error
//
// Availability: This function is available since v0.1.0.
//
// Examples:
//
//   ToTime(1547549353)      returns Time(Tuesday, January 15, 2019 10:49:13.000 AM GMT), nil
//   ToTime("1547549353123") returns Time(Tuesday, January 15, 2019 10:49:13.123 AM GMT), nil
//   ToTime(-1)              returns _, error
func ToTime(v interface{}) (time.Time, error) {
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Struct:
		if vV.Type().PkgPath() == "time" && vV.Type().Name() == "Time" {
			// same type, just cast it
			return vV.Interface().(time.Time), nil
		}
		return ZeroTime, errors.New("value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [time.Time]")
	}
	v, e := ToInt(v)
	if e != nil {
		return ZeroTime, errors.New("value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [time.Time]")
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
	return ZeroTime, errors.New("value of [" + fmt.Sprint(v) + "] cannot be converted to [time.Time]")
}

// ToStruct converts a value (v) to struct of type specified by (t) (t must be a struct). The output is guaranteed to have the same type as (t).
//
//   - If v is a struct:
//     - If v and t are same type, simply cast v to the specified type and return it
//     - Otherwise, loop through v's fields. If there is a field that is same type as t, return it
//     - (since v0.1.1) special case: if t is 'time.Time', return result from ToTime(v)
//   - Otherwise, return error
//
// Examples:
//
//   type Abc struct{ Key1 int }
//   type Dev struct {
//     Abc
//     Key2 string
//   }
//   ToStruct(Abc{Key1:1}, Abc{})                   returns Abc{Key1:1},                   nil
//   ToStruct(Abc{Key1:1}, Def{})                   returns _,                             error
//   ToStruct(Def{Abc:Abc{Key1:1},Key2:"a"}, Def{}) returns Def{Abc:Abc{Key1:1},Key2:"a"}, nil
//   ToStruct(Def{Abc:Abc{Key1:1},Key2:"a"}, Abc{}) returns Abc{Key1:1},                   nil
//   ToStruct(Abc{Key1:1}, "")                      returns _,                             error
//   ToStruct("", Abc{})                            returns _,                             error
func ToStruct(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
	if tV.Kind() != reflect.Struct {
		return nil, errors.New("target type must be a struct, but received [" + fmt.Sprint(tV.Type()) + "]")
	}

	if tV.Type().PkgPath() == "time" && tV.Type().Name() == "Time" {
		// special case: convert value to 'time.Time'
		return ToTime(v)
	}

	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Struct:
		if vV.Type().Name() == tV.Type().Name() {
			// same type, just cast it
			return vV.Interface(), nil
		}
		// difference type, look into fields
		for i, n := 0, vV.NumField(); i < n; i++ {
			f := vV.Field(i)
			if f.Kind() == reflect.Struct {
				if f.Type().Name() == tV.Type().Name() {
					return f.Interface(), nil
				}
			}
		}
	}
	return nil, errors.New("value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [" + fmt.Sprint(tV.Type()) + "]")
}

// ToSlice converts a value (v) to slice of type specified by (t) (t must be a slice or array, not an element of slice/array).
// The output is guaranteed to have the same type as (t).
//
//   - If v is an array or slice: convert each element of v to the correct type (specified by t), put them into a slice, and finally return it.
//   - Otherwise, return error
//
// Notes:
//   - Array/slice is converted to slice
//   - Element type can be converted too, for example: []int can be converted to []string
//
// Examples:
//
//   ToSlice([]bool{true,false}, [0]int{})    returns []int{1,0},             nil
//   ToSlice([3]int{-1,0,1}, []string{""})    returns []string{"-1","0","1"}, nil
func ToSlice(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
	if tV.Kind() != reflect.Array && tV.Kind() != reflect.Slice {
		return nil, errors.New("target type must be an array or slice, but received [" + fmt.Sprint(tV.Type()) + "]")
	}
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Array, reflect.Slice:
		elementType := tV.Type().Elem()                                // type of slice element
		zero := reflect.Zero(elementType)                              // create a 'zero' value of type "elementType"
		slice := reflect.MakeSlice(reflect.SliceOf(zero.Type()), 0, 0) // create an empty slice
		for i, n := 0, vV.Len(); i < n; i++ {
			e, err := Convert(vV.Index(i).Interface(), zero.Interface())
			if err == nil {
				slice = reflect.Append(slice, reflect.ValueOf(e).Convert(elementType))
			} else {
				return nil, err
			}
		}
		return slice.Interface(), nil
	}
	return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + fmt.Sprint(tV.Type()) + "]")
}

// ToMap converts a value (v) to map where types of key & value are specified by (t) (t must be a map).
// The output is guaranteed to have the same type as (t).
//
//   - If v is a map: convert each element {key:value} of v to the correct type (specified by t), put them into a map, and finally return it.
//   - Otherwise, return error
//
// Notes:
//   - Element type can be converted too, for example: map[int]int can be converted to map[string]string
//
// Examples:
//
//   ToMap(map[string]bool{"a":true,"b":false}, map[string]int{})    returns map[string]int{"a":1,"b":0"}, nil
func ToMap(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
	if tV.Kind() != reflect.Map {
		return nil, errors.New("target type must be a map, but received [" + fmt.Sprint(tV.Type()) + "]")
	}
	vV := reflect.ValueOf(v)
	if vV.Kind() == reflect.Map {
		keyType := tV.Type().Key()           // type of map's key
		zeroKey := reflect.Zero(keyType)     // create a 'zero' value of type "keyType"
		valueType := tV.Type().Elem()        // type of map's value
		zeroValue := reflect.Zero(valueType) // create a 'zero' value of type "keyType"
		m := reflect.MakeMap(reflect.MapOf(keyType, valueType))
		for _, k := range vV.MapKeys() {
			key, err := Convert(k.Interface(), zeroKey.Interface())
			if err != nil {
				return nil, err
			}
			value, err := Convert(vV.MapIndex(k).Interface(), zeroValue.Interface())
			if err != nil {
				return nil, err
			}
			m.SetMapIndex(reflect.ValueOf(key).Convert(keyType), reflect.ValueOf(value).Convert(valueType))
		}
		return m.Interface(), nil
	}
	return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + fmt.Sprint(tV.Type()) + "]")
}

// ToPointer converts a value (v) to pointer of type specified by (t) (t must be a pointer).
// The output is guaranteed to have the same type as (t).
//
// Example 1:
//
//   a := float64(1.23)
//   zero := int32(0)
//   output, err := ToPointer(&a, &zero)
//   /* here err should be nil */
//   if err != nil {
//     panic("Something wrong!")
//   }
//   /* we now successfully converted *float644 to *int32 */
//   i32 := *output.(*interface{}) // note: type of output is *interface{}
//   fmt.Println(i32.(int32))      // i32 is safe to type asserted .(int32)
//
// Example 2:
//
//   type Abc struct {
//     A int
//   }
//   type Def struct {
//     Abc
//     D string
//   }
//   a := Def{Abc: Abc{1}, D: "2"}
//   output, err := ToPointer(&a, &Abc{})
//   /* here err should be nil */
//   if err != nil {
//     panic("Something wrong!")
//   }
//   /* we now successfully converted *Def to *Abc */
//   abc := *output.(*interface{}) // note: type of output is *interface{}
//   fmt.Println(abc.(Abc))        // i32 is safe to type asserted .(Abc)
func ToPointer(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
	if tV.Kind() != reflect.Ptr {
		return nil, errors.New("target type must be a pointer, but received [" + fmt.Sprint(tV.Type()) + "]")
	}
	vV := reflect.ValueOf(v)
	if vV.Kind() == reflect.Ptr {
		v, err := Convert(vV.Elem().Interface(), tV.Elem().Interface())
		if err != nil {
			return nil, err
		}
		x := reflect.ValueOf(v).Convert(tV.Elem().Type()).Interface()
		return &x, nil
	}
	return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + fmt.Sprint(tV.Type()) + "]")
}

// Convert converts a value (v) to specified type (t):
//
//   - If t is a bool: see ToBool(interface{}) (interface{}, error)
//   - If t is an integer (int, int8, int16, int32, int64): see ToInt(interface{}) (interface{}, error)
//   - If t is an unsigned-integer (uint, uint8, uint16, uint32, uint64, uintptr): see ToUint(interface{}) (interface{}, error)
//   - If t is a float (float32, float64): see ToFloat(interface{}) (interface{}, error)
//   - If t is a string: see ToString(interface{}) (interface{}, error)
//   - If t is a struct: see ToStruct(interface{}, interface{}) (interface{}, error)
//   - If t is an array or a slice: see ToSlice(v interface{}, interface{}) (interface{}, error)
//   - If t is a map: see ToMap(interface{}, interface{}) (interface{}, error)
//   - If t is a pointer: see ToPointer(interface{}, interface{}) (interface{}, error)
func Convert(v interface{}, t interface{}) (interface{}, error) {
	if v == nil || t == nil {
		return nil, errors.New("cannot convert: both (v) and (t) must not be nil")
	}
	k := reflect.TypeOf(t).Kind()
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
	}
	return v, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + fmt.Sprint(reflect.TypeOf(t)) + "]")
}
