# consu/reddo

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/reddo)](https://pkg.go.dev/github.com/btnguyen2k/consu/reddo)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/reddo/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/reddo/graph/badge.svg)](https://app.codecov.io/gh/btnguyen2k/consu/tree/reddo/reddo)

Package `reddo` provides utility functions to convert values using Golang's reflection.

## Installation

```shell
$ go get -u github.com/btnguyen2k/consu/reddo
```

## Usage

```go
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
	fmt.Printf("Id is %s, year is %d (%d)\n", id, year, yearUint) // Id is 1, year is 2019 (2019) 

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
	fmt.Println(v) // 2019-01-15 17:49:13 +0700 +07
	v,_ = reddo.ToTime("1547549353123")
	fmt.Println(v) // 2019-01-15 17:49:13.123 +0700 +07
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Support and Contribution

Feel free to create [pull requests](https://github.com/btnguyen2k/consu/pulls) or [issues](https://github.com/btnguyen2k/consu/issues) to report bugs or suggest new features.
Please search the existing issues before filing new issues to avoid duplicates. For new issues, file your bug or feature request as a new issue.

If you find this project useful, please star it.
