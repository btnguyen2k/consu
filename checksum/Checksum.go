/*
Package checksum provides utility functions to calculate checksum.

	- A value of type integer will have the same checksum regardless it is int, int8, int16, int32, int64, uint, uint8, uint16, uint32 or uint64. E.g. checksum(int(103)) == checksum(uint64(103))
	- A value of type float will have the same checksum regardless it is float32 or float64. E.g. checksum(float32(10.3)) == checksum(float64(10.3))
	- Pointer to a value will have the same checksum as the value itself. E.g. checksum(myInt) == checksum(&myInt)
	- Slice and Array: will have the same checksum. E.g. checksum([]int{1,2,3}) == checksum([3]int{1,2,3})
	- Map and Struct: order of fields does not affect checksum, but field names do! E.g. checksum(map[string]int{"one":1,"two":2}) == checksum(map[string]int{"two":2,"one":1}), but checksum(map[string]int{"a":1,"b":2}) != checksum(map[string]int{"x":1,"y":2})
	- Struct: be able to calculate checksum of unexported fields.

Sample usage:

	package main

	import (
		"fmt"
		"github.com/btnguyen2k/consu/checksum"
	)

	func main() {
		myValue := "any thing"

		// calculate checksum using MD5 hash
		checksum1 := Checksum(Md5HashFunc, myValue)
		fmt.Printf("%x\n", checksum1)

		// shortcut to calculate checksum using MD5 hash
		checksum2 := Md5Checksum(myValue)
		fmt.Printf("%x\n", checksum2)
	}
*/
package checksum

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash"
	"hash/crc32"
	"reflect"
	"strings"
	"unsafe"
)

const (
	// Version defines version number of this package
	Version = "0.1.0"
)

/*
HashFunc defines a function that calculates hash value of a byte array.
*/
type HashFunc func(input []byte) []byte

func hashFunc(hf hash.Hash, input []byte) []byte {
	hf.Write(input)
	return hf.Sum(nil)
}

// Crc32HashFunc is a HashFunc that calculates hash value using CRC32.
var Crc32HashFunc HashFunc = func(input []byte) []byte {
	return hashFunc(crc32.NewIEEE(), input)
}

// Md5HashFunc is a HashFunc that calculates hash value using MD5.
var Md5HashFunc HashFunc = func(input []byte) []byte {
	return hashFunc(md5.New(), input)
}

// Sha1HashFunc is a HashFunc that calculates hash value using SHA1.
var Sha1HashFunc HashFunc = func(input []byte) []byte {
	return hashFunc(sha1.New(), input)
}

// Sha256HashFunc is a HashFunc that calculates hash value using SHA256.
var Sha256HashFunc HashFunc = func(input []byte) []byte {
	return hashFunc(sha256.New(), input)
}

// Sha512HashFunc is a HashFunc that calculates hash value using SHA512.
var Sha512HashFunc HashFunc = func(input []byte) []byte {
	return hashFunc(sha512.New(), input)
}

func boolToBytes(v bool) []byte {
	if v {
		return []byte{1}
	}
	return []byte{0}
}

func intToBytes(v int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func uintToBytes(v uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func floatToBytes(v float64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func isExportedField(fieldName string) bool {
	return len(fieldName) >= 0 && string(fieldName[0]) == strings.ToUpper(string(fieldName[0]))
}

/*
Checksum calculates checksum of an input using the provided hash function.

	- If v is a scalar type (bool, int*, uint*, float* or string) or pointer to scala type: checksum value is straightforward calculation.
	- If v is a slice or array: checksum value is combination of all elements' checksums, in order. If v is empty (has 0 elements), empty []byte is returned.
	- If v is a map: checksum value is combination of all entries' checksums, order-independent.
	- If v is a struct: checksum value is combination of all fields' checksums, order-independent.
*/
func Checksum(hf HashFunc, v interface{}) []byte {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Bool:
		return hf(boolToBytes(rv.Bool()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return hf(intToBytes(rv.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return hf(uintToBytes(rv.Uint()))
	case reflect.Float32, reflect.Float64:
		return hf(floatToBytes(rv.Float()))
	case reflect.String:
		return hf([]byte(rv.String()))
	case reflect.Array, reflect.Slice:
		buf := make([]byte, 0)
		for i, n := 0, rv.Len(); i < n; i++ {
			buf = hf(append(buf, Checksum(hf, rv.Index(i).Interface())...))
		}
		return buf
	case reflect.Map:
		buf := hf([]byte{})
		for iter := rv.MapRange(); iter.Next(); {
			// field-name is taking into account
			temp := Checksum(hf, []interface{}{iter.Key().Interface(), iter.Value().Interface()})
			for i, n := 0, len(buf); i < n; i++ {
				buf[i] ^= temp[i]
			}
		}
		return buf
	case reflect.Struct:
		buf := hf([]byte{})
		for i, n := 0, rv.NumField(); i < n; i++ {
			// field-name is taking into account
			fieldName := rv.Type().Field(i).Name
			fieldValue := rv.Field(i)
			if !isExportedField(fieldName) {
				// handle unexported field
				rv2 := reflect.New(rv.Type()).Elem()
				rv2.Set(rv)
				fieldValue = rv2.Field(i)
				fieldValue = reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
			}
			temp := Checksum(hf, []interface{}{fieldName, fieldValue.Interface()})
			for i, n := 0, len(buf); i < n; i++ {
				buf[i] ^= temp[i]
			}
		}
		return buf
	}
	return nil
}

/*
Crc32Checksum is shortcut of Checksum(Crc32HashFunc, v).
*/
func Crc32Checksum(v interface{}) []byte {
	return Checksum(Crc32HashFunc, v)
}

/*
Md5Checksum is shortcut of Checksum(Md5HashFunc, v).
*/
func Md5Checksum(v interface{}) []byte {
	return Checksum(Md5HashFunc, v)
}

/*
Sha1Checksum is shortcut of Checksum(Sha1HashFunc, v).
*/
func Sha1Checksum(v interface{}) []byte {
	return Checksum(Sha1HashFunc, v)
}

/*
Sha256Checksum is shortcut of Checksum(Sha256HashFunc, v).
*/
func Sha256Checksum(v interface{}) []byte {
	return Checksum(Sha256HashFunc, v)
}

/*
Sha512Checksum is shortcut of Checksum(Sha512HashFunc, v).
*/
func Sha512Checksum(v interface{}) []byte {
	return Checksum(Sha512HashFunc, v)
}
