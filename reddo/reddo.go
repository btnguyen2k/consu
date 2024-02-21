package reddo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
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

	// ZeroMode controls how nil is converted.
	// When set to false, converting nil to primitive data types (number, string) results in error.
	// When set to true, nil is converted to 'zero' value in primitive data types (0 for number, "" for string).
	//
	// Available since v0.1.7
	ZeroMode = true
)

var (
	// ErrorNilToBool is returned by ToBool when ZeroMode=false
	ErrorNilToBool = errors.New("cannot convert nil to bool")

	// ErrorNilToFloat is returned by ToFloat when ZeroMode=false
	ErrorNilToFloat = errors.New("cannot convert nil to float")

	// ErrorNilToInt is returned by ToInt when ZeroMode=false
	ErrorNilToInt = errors.New("cannot convert nil to int")

	// ErrorNilToUint is returned by ToUint when ZeroMode=false
	ErrorNilToUint = errors.New("cannot convert nil to uint")

	// ErrorNilToString is returned by ToString when ZeroMode=false
	ErrorNilToString = errors.New("cannot convert nil to string")

	// ErrorNilToTime is returned by ToTime when ZeroMode=false
	ErrorNilToTime = errors.New("cannot convert nil to time.Time")

	// ErrorNilToStruct is returned by ToStruct when ZeroMode=false
	ErrorNilToStruct = errors.New("cannot convert nil to struct")
)

var zeroTime = time.Time{}

/*
ToBool converts a value to bool.

  - If v is nil: return false if ZeroMode is true, error otherwise.
  - If v is indeed a bool: its value is returned.
  - If v is a number (integer, float or complex): return false if its value is 'zero', true otherwise.
  - If v is a pointer: return false if it is nil, true otherwise.
  - If v is a string: return result from strconv.ParseBool(string).
  - Otherwise, return error

Examples:

	ToBool(nil)           return true,  nil (if ZeroMode = true)
	ToBool(nil)           return _,     error (if ZeroMode = false)
	ToBool(true)          return true,  nil
	ToBool(false)         return false, nil
	ToBool(0)             return false, nil
	ToBool(1)             return true,  nil
	ToBool(-1)            return true,  nil
	ToBool(0.0)           return false, nil
	ToBool(1.2)           return true,  nil
	ToBool(-3.4)          return true,  nil
	ToBool(1i)            return true,  nil
	ToBool(-1i)           return true,  nil
	ToBool(0i)            return false, nil
	ToBool("true")        return true,  nil
	ToBool("false")       return false, nil
	ToBool("blabla")      return _,     error
	ToBool(struct{}{})    return _,     error
*/
func ToBool(v interface{}) (bool, error) {
	if v == nil {
		if ZeroMode {
			return false, nil
		}
		return false, ErrorNilToBool
	}
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
	default:
		return false, fmt.Errorf("cannot convert value [%s] to bool", vV.String())
	}
}

/*
ToFloat converts a value to float64.

  - If v is nil: return 0.0 if ZeroMode is true, error otherwise.
  - If v is a bool: return 1.0 if its value is true, 0.0 otherwise.
  - If v is a number (integer or float): return its value as float64.
  - If v is a string: return result from strconv.ParseFloat(string).
  - Otherwise, return error

Examples:

	ToFloat(nil)           return 0.0,  nil (if ZeroMode = true)
	ToFloat(nil)           return _,    error (if ZeroMode = false)
	ToFloat(true)          return 1.0,  nil
	ToFloat(false)         return 0.0,  nil
	ToFloat(0)             return 0.0,  nil
	ToFloat(1.2)           return 1.2,  nil
	ToFloat("-3.4")        return -3.4, nil
	ToFloat("blabla")      return _,     error
	ToFloat(struct{}{})    return _,     error
*/
func ToFloat(v interface{}) (float64, error) {
	if v == nil {
		if ZeroMode {
			return 0.0, nil
		}
		return 0.0, ErrorNilToFloat
	}
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		if vV.Bool() {
			return 1.0, nil
		}
		return 0.0, nil
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
	default:
		return 0.0, fmt.Errorf("cannot convert value [%s] to float64", vV.String())
	}
}

/*
ToInt converts a value to int64.

  - If v is nil: return 0 if ZeroMode is true, error otherwise.
  - If v is a number (integer or float): return its value as int64.
  - If v is a bool: return 1 if its value is true, 0 otherwise.
  - If v is a string: return result from strconv.ParseInt(string).
  - Otherwise, return error

Examples:

	ToInt(nil)           return 0,  nil (if ZeroMode = true)
	ToInt(nil)           return _,  error (if ZeroMode = false)
	ToInt(true)          return 1,  nil
	ToInt(false)         return 0,  nil
	ToInt(0)             return 0,  nil
	ToInt(1.2)           return 1,  nil
	ToInt("-3")          return -3, nil
	ToInt("4.5")         return _,  error
	ToInt("blabla")      return _,  error
	ToInt(struct{}{})    return _,  error
*/
func ToInt(v interface{}) (int64, error) {
	if v == nil {
		if ZeroMode {
			return 0, nil
		}
		return 0, ErrorNilToInt
	}
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		if vV.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vV.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(vV.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int64(vV.Float()), nil
	case reflect.String:
		return strconv.ParseInt(vV.String(), 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert value [%s] to int64", vV.String())
	}
}

/*
ToUint converts a value to uint64.

  - If v is nil: return 0 if ZeroMode is true, error otherwise.
  - If v is a number (integer or float): return its value as uint64.
  - If v is a bool: return 1 if its value is true, 0 otherwise.
  - If v is a string: return result from strconv.ParseUint(string).
  - Otherwise, return error

Examples:

	ToUint(nil)           return 0,  nil (if ZeroMode = true)
	ToUint(nil)           return _,  error (if ZeroMode = false)
	ToUint(true)          return 1,  nil
	ToUint(false)         return 0,  nil
	ToUint(0)             return 0,  nil
	ToUint(1.2)           return 1,  nil
	ToUint(-1)            return 18446744073709551615,  nil // be caution with negative numbers!
	ToUint("-3")          return _,  error
	ToUint("4.5")         return _,  error
	ToUint("blabla")      return _,  error
	ToUint(struct{}{})    return _,  error
*/
func ToUint(v interface{}) (uint64, error) {
	if v == nil {
		if ZeroMode {
			return 0, nil
		}
		return 0, ErrorNilToUint
	}
	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Bool:
		if vV.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(vV.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return vV.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return uint64(vV.Float()), nil
	case reflect.String:
		return strconv.ParseUint(vV.String(), 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert value [%s] to uint64", vV.String())
	}
}

/*
ToString converts a value to string.

  - If v is nil: return "" if ZeroMode is true, error otherwise.
  - If v is a number (integer or float) or bool or string: return its value as string.
  - Otherwise, return string representation of v (fmt.Sprint(v))

Examples:

	ToString(nil)           return "",       nil (if ZeroMode = true)
	ToString(nil)           return _,        error (if ZeroMode = false)
	ToString(true)          return "true",   nil
	ToString(false)         return "false",  nil
	ToString(0)             return "0",      nil
	ToString(1.2)           return "1.2",    nil
	ToString("blabla")      return "blabla", nil
	ToString(struct{}{})    return "{}",     nil
*/
func ToString(v interface{}) (string, error) {
	if v == nil {
		if ZeroMode {
			return "", nil
		}
		return "", ErrorNilToString
	}
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
	case reflect.Slice, reflect.Array:
		// since v0.1.4.1: (special case) convert []byte to string
		if vV.Type().Elem().Kind() == reflect.Uint8 {
			return string(vV.Interface().([]byte)), nil
		}
	default:
	}
	return fmt.Sprint(v), nil
}

/*
ToTime converts a value (v) to 'time.Time'.

  - If v is nil: return zero-time if ZeroMode is true, error otherwise.
  - If v is 'time.Time': return v.
  - If v is integer: depends on how big v is, treat v as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
  - If v is string and convertible to integer: depends on how big v is, treat v as UNIX timestamp in seconds, milliseconds, microseconds or nanoseconds, convert to 'time.Time' and return the result.
  - Otherwise, return error

Availability: This function is available since v0.1.0.

Examples:

	ToTime(nil)                return zero-time, nil (if ZeroMode = true)
	ToTime(nil)                return _,         error (if ZeroMode = false)
	ToTime(1547549353)         return Time(Tuesday, January 15, 2019 10:49:13.000 AM GMT), nil
	ToTime("1547549353123")    return Time(Tuesday, January 15, 2019 10:49:13.123 AM GMT), nil
	ToTime(-1)                 return _, error
*/
func ToTime(v interface{}) (time.Time, error) {
	if v == nil {
		if ZeroMode {
			return zeroTime, nil
		}
		return zeroTime, ErrorNilToTime
	}
	vTyp := reflect.ValueOf(v).Type()
	if vTyp == TypeTime {
		// same type, just cast it
		return v.(time.Time), nil
	}

	v, e := ToInt(v)
	if e != nil {
		return zeroTime, errors.New("value of type [" + fmt.Sprint(vTyp) + "] cannot be converted to [time.Time]")
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

	ToTimeWithLayout(nil, _)                                                              return zero-time, nil (if ZeroMode = true)
	ToTimeWithLayout(nil, _)                                                              return _,         error (if ZeroMode = false)
	ToTimeWithLayout(1547549353, _)                                                       return Time(Tuesday, January 15, 2019 10:49:13.000 AM GMT), nil
	ToTimeWithLayout("1547549353123", _)                                                  return Time(Tuesday, January 15, 2019 10:49:13.123 AM GMT), nil
	ToTimeWithLayout(-1, _)                                                               return _, error
	ToTimeWithLayout("January 15, 2019 20:49:13.123", "January 02, 2006 15:04:05.000")    return Time(Tuesday, January 15, 2019 08:49:13.123 PM GMT), nil
*/
func ToTimeWithLayout(v interface{}, layout string) (time.Time, error) {
	if v == nil {
		if ZeroMode {
			return zeroTime, nil
		}
		return zeroTime, ErrorNilToTime
	}
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
	return len(fieldName) > 0 && string(fieldName[0]) == strings.ToUpper(string(fieldName[0]))
}

/*
ToStruct converts a value (v) to struct of type specified by (t) (t must be a struct). The output is guaranteed to have the same type as (t).

  - If v is nil: return zero-value of type (t) if ZeroMode=true, error otherwise.
  - If v is a struct:
    a) If v and t are same type, simply cast v to the specified type and return it
    b) Otherwise, loop through v's fields. If there is an exported field that is same type as t, return it
    c) (since v0.1.1) special case: if t is 'time.Time', return result from ToTime(v)
  - Otherwise, return error

Examples:

	type Abc struct{ Key1 int }
	typeAbc := reflect.TypeOf(Abc{})
	type Def struct {
		Abc
		Key2 string
	}
	typeDef := reflect.TypeOf(Def{})

	ToStruct(nil, typeAbc)                              return Abc{},                         nil
	ToStruct(nil, typeDef)                              return Def{},                         nil
	ToStruct(Abc{Key1:1}, typeAbc)                      return Abc{Key1:1},                   nil (rule a)
	ToStruct(Abc{Key1:1}, typeDef)                      return _,                             error
	ToStruct(Def{Abc:Abc{Key1:1},Key2:"a"}, typeDef)    return Def{Abc:Abc{Key1:1},Key2:"a"}, nil (rule a)
	ToStruct(Def{Abc:Abc{Key1:1},Key2:"a"}, typeAbc)    return Abc{Key1:1},                   nil (rule b)
	ToStruct(Abc{Key1:1}, reflect.TypeOf(""))           return _,                             error
	ToStruct("", typeAbc)                               return _,                             error
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

	if v == nil {
		if ZeroMode {
			return reflect.New(t).Elem().Interface(), nil
		}
		return nil, ErrorNilToStruct
	}

	vV := reflect.ValueOf(v)
	switch vV.Kind() {
	case reflect.Struct:
		if vV.Type().Name() == t.Name() {
			// same type, just cast it
			return vV.Interface(), nil
		}
		// different types, look into fields
		for i, n := 0, vV.NumField(); i < n; i++ {
			f := vV.Field(i)
			fn := vV.Type().Field(i).Name
			if f.Kind() == reflect.Struct && isExportedField(fn) {
				if f.Type().Name() == t.Name() {
					return f.Interface(), nil
				}
			}
		}
	default:
	}
	return nil, errors.New("value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [" + t.String() + "]")
}

/*
ToSlice converts a value (v) to slice of type specified by (typ) (typ can be a slice or array, or an element of slice/array).
The output is guaranteed to have the same type as (typ).

  - If v is nil: return nil regardless ZeroMode.
  - If v is an array or slice: convert each element of v to the correct type (specified by typ), put them into a slice, and finally return it.
  - Otherwise, return error

Notes:

  - Array/slice is converted to slice.
  - Element type is converted as well, for example: []int can be converted to []string

Examples:

	ToSlice(nil, _)                                          return nil,                      nil
	ToSlice([]bool{true,false}, reflect.TypeOf([0]int{}))    return []int{1,0},               nil
	ToSlice([3]int{-1,0,1}, reflect.TypeOf([]string{""}))    return []string{"-1","0","1"},   nil
	ToSlice([]bool{true,false}, TypeString)                  return []string{"true","false"}, nil
	ToSlice(_, nil)                                          return _,                        error
*/
func ToSlice(v interface{}, typ reflect.Type) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	if typ == nil {
		return nil, errors.New("cannot detect type of target as it is [nil]")
	}
	if typ.Kind() != reflect.Array && typ.Kind() != reflect.Slice {
		return ToSlice(v, reflect.SliceOf(typ))
	}
	vV := reflect.ValueOf(v)
	if typ.Elem().Kind() == reflect.Uint8 && vV.Kind() == reflect.String {
		// since v0.1.4.1: (special case) converting string to []byte
		return []byte(v.(string)), nil
	}
	switch vV.Kind() {
	case reflect.Array, reflect.Slice:
		elementType := typ.Elem()                                      // type of slice element
		slice := reflect.MakeSlice(reflect.SliceOf(elementType), 0, 0) // create an empty slice
		for i, n := 0, vV.Len(); i < n; i++ {
			e, err := Convert(vV.Index(i).Interface(), elementType)
			if err != nil {
				return nil, err
			}
			if e == nil {
				slice = reflect.Append(slice, reflect.New(elementType).Elem())
			} else {
				slice = reflect.Append(slice, reflect.ValueOf(e).Convert(elementType))
			}
		}
		return slice.Interface(), nil
	default:
		return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + typ.String() + "]")
	}
}

/*
ToMap converts a value (v) to map where types of key & value are specified by (typ) (typ must be a map).
The output is guaranteed to have the same type as (typ).

  - If v is nil: return nil regardless ZeroMode.
  - If v is a map: convert each element {key:value} of v to the correct type (specified by typ), put them into a map, and finally return it.
  - Otherwise, return error

Notes:

  - Element type is converted as well, for example: map[int]int can be converted to map[string]string

Examples:

	ToMap(nil, _)                                                                   return nil,                          nil
	ToMap(map[string]bool{"a":true,"b":false}, reflect.TypeOf(map[string]int{}))    return map[string]int{"a":1,"b":0"}, nil
	ToMap(_, nil)                                                                   return _,                            error
*/
func ToMap(v interface{}, typ reflect.Type) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	if typ == nil {
		return nil, errors.New("cannot detect type of target as it is [nil]")
	}
	if typ.Kind() != reflect.Map {
		return nil, errors.New("target type must be a map, but received [" + typ.String() + "]")
	}

	vV := reflect.ValueOf(v)
	if vV.Kind() == reflect.Map {
		keyType := typ.Key()    // type of map's key
		valueType := typ.Elem() // type of map's value
		m := reflect.MakeMap(reflect.MapOf(keyType, valueType))
		for _, k := range vV.MapKeys() {
			key, err := Convert(k.Interface(), keyType)
			if err != nil {
				return nil, err
			}
			vkey := reflect.ValueOf(key).Convert(keyType)

			vvalue := vV.MapIndex(k)
			value, err := Convert(vvalue.Interface(), valueType)
			if err != nil {
				return nil, err
			}
			if value == nil {
				vvalue = reflect.New(valueType).Elem()
			} else {
				vvalue = reflect.ValueOf(value).Convert(valueType)
			}

			m.SetMapIndex(vkey, vvalue)
		}
		return m.Interface(), nil
	}
	return nil, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + typ.String() + "]")
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
	// we now successfully converted *float64 to *int32
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
	if v == nil {
		return nil, nil
	}
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
  - (special case) If t is nil: this function return (v, nil)
*/
func Convert(v interface{}, t reflect.Type) (interface{}, error) {
	if t == nil {
		return v, nil
	}
	// if v == nil {
	// 	return nil, errors.New("cannot convert: nil to " + t.String())
	// }
	switch t.Kind() {
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
	default:
		return v, errors.New("cannot convert [" + fmt.Sprint(v) + "] to [" + t.String() + "]")
	}
}
