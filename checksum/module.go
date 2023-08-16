/*
Package checksum provides utility functions to calculate checksum.

  - A value of type integer will have the same checksum regardless its type (int, int8, int16, int32, int64, uint, uint8, uint16, uint32 or uint64). E.g. checksum(int(103)) == checksum(uint64(103))
  - A value of type float will have the same checksum regardless its type (float32 or float64). E.g. checksum(float32(10.3)) == checksum(float64(10.3))
  - Pointer to a value will have the same checksum as the value itself. E.g. checksum(myInt) == checksum(&myInt)
  - Slice and Array: will have the same checksum. E.g. checksum([]int{1,2,3}) == checksum([3]int{1,2,3})
  - Map and Struct: order of fields does not affect checksum, but field names do! E.g. checksum(map[string]int{"one":1,"two":2}) == checksum(map[string]int{"two":2,"one":1}), but checksum(map[string]int{"a":1,"b":2}) != checksum(map[string]int{"x":1,"y":2})
  - Struct: be able to calculate checksum of unexported fields.

Note on special inputs:

  - Checksum of `nil` is a slice where all values are zeroes.

Sample usage:

	package main

	import (
		"fmt"

		"github.com/btnguyen2k/consu/checksum"
	)

	func main() {
		myValue := "any thing"

		// calculate checksum using MD5 hash
		checksum1 := checksum.Checksum(checksum.Md5HashFunc, myValue)
		fmt.Printf("%x\n", checksum1)

		// shortcut to calculate checksum using MD5 hash
		checksum2 := checksum.Md5Checksum(myValue)
		fmt.Printf("%x\n", checksum2)
	}
*/
package checksum

const (
	// Version defines version number of this package
	Version = "1.0.0"
)
