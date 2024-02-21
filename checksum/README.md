# consu/checksum

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/checksum)](https://pkg.go.dev/github.com/btnguyen2k/consu/checksum)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/checksum/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/checksum/graph/badge.svg)](https://app.codecov.io/gh/btnguyen2k/consu/tree/checksum/checksum)

Package `checksum` provides utility functions to calculate checksum.

## Installation

```shell
$ go get -u github.com/btnguyen2k/consu/checksum
```

## Usage

```go
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
```

## Features:

⭐ Calculate checksum of scalar types (`bool`, `int*`, `uint*`, `float*`, `string`) as well as `slice/array` and `map/struct`.

⭐ `Struct`:
  - If `time.Time`, its nanosecond is used to calculate checksum (since `v0.1.2`).
  - Be able to calculate checksum of unexported fields.
  - If the struct has function `Checksum()`, use it instead of reflecting through struct fields.

⭐ Supported hash functions: `CRC32`, `MD5`, `SHA1`, `SHA256`, `SHA512`.

⭐ A value of type integer will have the same checksum regardless its type (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32` or `uint64`).
E.g. `Checksum(int(103)) == Checksum(uint64(103))`

⭐ A value of type float will have the same checksum regardless its type (`float32` or `float64`).
E.g. `Checksum(float32(10.3)) == Checksum(float64(10.3))`

⭐ Pointer to a value will have the same checksum as the value itself.
E.g. `Checksum(myInt) == Checksum(&myInt)`

⭐ `Slice` and `Array` will have the same checksum.
E.g. `Checksum([]int{1,2,3}) == Checksum([3]int{1,2,3})`

⭐ `Map` and `Struct`: order of fields does not affect checksum, but field names do!
E.g. `Checksum(map[string]int{"one":1,"two":2}) == Checksum(map[string]int{"two":2,"one":1})`,
but `Checksum(map[string]int{"a":1,"b":2}) != Checksum(map[string]int{"x":1,"y":2})`

⭐ `Map` and `Struct` have different checksums even if they have the same fields and values.

⭐ Two different `Struct`s have different checksums even if they have the same fields and values.

Note on special inputs:

⭐ `Checksum(nil)` returns a slice where all values are zero.

⭐ All empty maps have the same checksum, e.g. `Checksum(map[string]int{}) == Checksum(map[int]string{})`. 

⭐ All empty slices/arrays have the same checksum, e.g. `Checksum([]int{}) == Checksum([0]int{}) == Checksum([]string{}) == Checksum([0]string{})`.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Support and Contribution

Feel free to create [pull requests](https://github.com/btnguyen2k/consu/pulls) or [issues](https://github.com/btnguyen2k/consu/issues) to report bugs or suggest new features.
Please search the existing issues before filing new issues to avoid duplicates. For new issues, file your bug or feature request as a new issue.

If you find this project useful, please star it.
