package checksum

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/crc32"
	"reflect"
	"sort"
	"strings"
	"time"
	"unsafe"
)

// HashFunc defines a function that calculates hash value of a byte array.
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

func unwrap(v interface{}) (prv reflect.Value, rv reflect.Value) {
	rv = reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.Elem().Kind() == reflect.Struct {
			prv = rv
		}
		rv = rv.Elem()
	}
	return
}

/*
Checksum calculates checksum of an input using the provided hash function.

  - If v is a scalar type (bool, int*, uint*, float* or string) or pointer to scala type: checksum value is straightforward calculation.
  - If v is a slice or array: checksum value is combination of all elements' checksums, in order. If v is empty (has 0 elements), empty []byte is returned.
  - If v is a map: checksum value is combination of all entries' checksums, order-independent.
  - If v is a struct: if the struct has function `Checksum()` then use it to calculate checksum value; if v is time.Time then use its nanosecond to calculate checksum value; otherwise checksum value is combination of all fields' checksums, order-independent.

Note on special inputs:

  - Checksum of `nil` is a slice where all values are zeroes.
*/
func Checksum(hf HashFunc, v interface{}) []byte {
	return checksumSafe(hf, v, make(map[interface{}]struct{}))
}

func checksumSafe(hf HashFunc, v interface{}, visited map[interface{}]struct{}) []byte {
	if v == nil {
		result := hf(nil)
		for i := range result {
			result[i] = 0
		}
		return result
	}
	prv, rv := unwrap(v)
	switch rv.Kind() {
	case reflect.Map, reflect.Slice:
		ptr := rv.Pointer()
		if _, ok := visited[ptr]; ok {
			return nil
		}
		visited[ptr] = struct{}{}
		defer delete(visited, ptr)
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
			buf = hf(append(buf, checksumSafe(hf, rv.Index(i).Interface(), visited)...))
		}
		return buf
	case reflect.Map:
		temp := []string{"0x10"}
		for iter := rv.MapRange(); iter.Next(); {
			// field-name is taking into account
			fieldChecksum := checksumSafe(hf, []interface{}{iter.Key().Interface(), iter.Value().Interface()}, visited)
			temp = append(temp, fmt.Sprintf("%x", fieldChecksum))
		}
		sort.Strings(temp)
		return checksumSafe(hf, temp, visited)
	case reflect.Struct:
		m := rv.MethodByName("Checksum")
		if !m.IsValid() && prv.IsValid() {
			m = prv.MethodByName("Checksum")
		}
		if m.IsValid() && m.Type().NumIn() == 0 {
			result := m.Call(nil)
			arr := []interface{}{"0x11", rv.Type().String()}
			for _, v := range result {
				arr = append(arr, v.Interface())
			}
			if len(arr) > 2 {
				return Checksum(hf, arr)
			}
		}

		if rv.Type() == reflect.TypeOf(time.Time{}) {
			vtemp := []interface{}{"time.Time", rv.Interface().(time.Time).UnixNano()}
			return checksumSafe(hf, vtemp, visited)
		}

		temp := []string{"0x11", rv.Type().String()}
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
			fieldChecksum := checksumSafe(hf, []interface{}{fieldName, fieldValue.Interface()}, visited)
			temp = append(temp, fmt.Sprintf("%x", fieldChecksum))
		}
		sort.Strings(temp)
		return checksumSafe(hf, temp, visited)
	}
	return nil
}

// Crc32Checksum is shortcut of Checksum(Crc32HashFunc, v).
func Crc32Checksum(v interface{}) []byte {
	return Checksum(Crc32HashFunc, v)
}

// Md5Checksum is shortcut of Checksum(Md5HashFunc, v).
func Md5Checksum(v interface{}) []byte {
	return Checksum(Md5HashFunc, v)
}

// Sha1Checksum is shortcut of Checksum(Sha1HashFunc, v).
func Sha1Checksum(v interface{}) []byte {
	return Checksum(Sha1HashFunc, v)
}

// Sha256Checksum is shortcut of Checksum(Sha256HashFunc, v).
func Sha256Checksum(v interface{}) []byte {
	return Checksum(Sha256HashFunc, v)
}

// Sha512Checksum is shortcut of Checksum(Sha512HashFunc, v).
func Sha512Checksum(v interface{}) []byte {
	return Checksum(Sha512HashFunc, v)
}
