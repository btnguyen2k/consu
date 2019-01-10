// Package reddo provides utilities to convert values using Golang's reflect.
package reddo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

const (
	// Version defines version number of this package
	Version = "0.1.0"
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
//    ToPool(true)     returns true,  nil
//    ToPool(false)    returns false, nil
//    ToPool(0)        returns false, nil
//    ToPool(1)        returns true,  nil
//    ToPool(0.0)      returns false, nil
//    ToPool(1.2)      returns true,  nil
//    ToPool(1i)       returns true,  nil
//    ToPool(0i)       returns false, nil
//    ToPool("true")   returns true,  nil
//    ToPool("false")  returns false, nil
//    ToPool("blabla") returns _,     error
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
	return false, errors.New("Cannot convert value [" + vV.String() + "] to bool.")
}

// ToInt converts a value to int64. The output is guaranteed to ad-here to type assertion .(int64)
//
//   - If v is a number (integer or float): return its value as int64.
//   - If v is a bool: return 1 if its value is true, 0 otherwise.
//   - If v is a string: return result from strconv.ParseInt(string).
//   - Otherwise, return error
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
	return int64(0), errors.New("Cannot convert value [" + vV.String() + "] to int64.")
}

// ToUint converts a value to uint64. The output is guaranteed to ad-here to type assertion .(uint64)
//
//   - If v is a number (integer or float): return its value as uint64.
//   - If v is a bool: return 1 if its value is true, 0 otherwise.
//   - If v is a string: return result from strconv.ParseUint(string).
//   - Otherwise, return error
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
	return uint64(0), errors.New("Cannot convert value [" + vV.String() + "] to uint64.")
}

// ToFloat converts a value to float64. The output is guaranteed to ad-here to type assertion .(float64)
//
//   - If v is a number (integer or float): return its value as float64.
//   - If v is a bool: return 1.0 if its value is true, 0.0 otherwise.
//   - If v is a string: return result from strconv.ParseFloat(string).
//   - Otherwise, return error
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
	case reflect.Float32, reflect.Float64:
		return vV.Float(), nil
	case reflect.String:
		return strconv.ParseFloat(vV.String(), 64)
	}
	return float64(0), errors.New("Cannot convert value [" + vV.String() + "] to float64.")
}

// ToString converts a value to string. The output is guaranteed to ad-here to type assertion .(string)
//
//   - If v is a number (integer or float) or bool or string: return its value as string.
//   - Otherwise, return string representation of v (fmt.Sprint(v))
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

// ToStruct converts a value (v) to struct of type specified by (t). The output is guaranteed to have the same type as (t).
//
//   - If v is a struct:
//     - If v and t are same type, simply cast v to the specified type and return it
//     - Otherwise, loop through v's fields. If there is a field that is same type as t, return it
//   - Otherwise, return error
func ToStruct(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
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
	return nil, errors.New("Value of type [" + fmt.Sprint(vV.Type()) + "] cannot be converted to [" + fmt.Sprint(tV.Type()) + "]")
}

// ToSlice converts a value (v) to slice of type specified by (t) (t must be a slice or array, not an element of slice/array).
// The output is guaranteed to have the same type as (t).
//
//   - If v is an array or slice: convert each element of v to the correct type (specified by t), put them into a slice, and finally return it.
//   - Otherwise, return error
func ToSlice(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
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
	return nil, errors.New("Cannot convert [" + fmt.Sprint(v) + "] to [" + fmt.Sprint(tV.Type()) + "]!")
}

// ToMap converts a value (v) to map where types of key & value are specified by (t) (t must be a map).
// The output is guaranteed to have the same type as (t).
//
//   - If v is a map: convert each element {key:value} of v to the correct type (specified by t), put them into a map, and finally return it.
//   - Otherwise, return error
func ToMap(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
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
	return nil, nil
}

// ToPointer converts a value (v) to pointer of type specified by (t) (t must be a pointer).
// The output is guaranteed to have the same type as (t).
//
//   - If v is a map: convert each element {key:value} of v to the correct type (specified by t), put them into a map, and finally return it.
//   - Otherwise, return error
func ToPointer(v interface{}, t interface{}) (interface{}, error) {
	tV := reflect.ValueOf(t)
	vV := reflect.ValueOf(v)
	if vV.Kind() == reflect.Ptr {
		v, err := Convert(vV.Elem().Interface(), tV.Elem().Interface())
		x := reflect.ValueOf(v).Convert(tV.Elem().Type()).Interface()
		p := unsafe.Pointer(&x)
		z := reflect.NewAt(tV.Elem().Type(), p)
		return z.Interface(), err
	}
	return nil, nil
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
		return nil, nil
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
	return v, nil
}
