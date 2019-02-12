# consu/reddo

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![GoDoc](https://godoc.org/github.com/btnguyen2k/consu/reddo?status.svg)](https://godoc.org/github.com/btnguyen2k/consu/reddo)

[GoCover](https://gocover.io/github.com/btnguyen2k/consu/reddo)

Package reddo provides utilities to convert values using Golang's reflect.

## Installation

```shell
$ go get -u github.com/btnguyen2k/consu/reddo
```


## Usage

```go
package main

import (
	"fmt"
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
func getValue(data map[string]interface{}, field string, zero interface{}) interface{} {
	v, err := reddo.Convert(data[field], zero)
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
	var id = getValue(data, "id", reddo.ZeroString).(string)
	var year = getValue(data, "year", reddo.ZeroInt).(int64)
	var yearUint = getValue(data, "year", reddo.ZeroUint).(uint64)
	fmt.Printf("Id is %s, year is %d (%d)\n", id, year, yearUint) // Id is 1, year is 2019 (2019) 

	// we need a 'zero' value of the target type to retrieve the correct value & type from out data store
	zeroAbc := Abc{}
	zeroDef := Def{}
	var abc = getValue(data, "abc", zeroAbc).(Abc)
	var def = getValue(data, "def", zeroDef).(Def)
	// special case: struct Def 'inherit' struct Abc, hence Def can be 'cast'-ed to Abc
	var abc2 = getValue(data, "def", zeroAbc).(Abc)
	fmt.Println("data.abc       :", abc)  // data.abc       : {103}
	fmt.Println("data.def       :", def)  // data.def       : {{1981} btnguyen2k}
	fmt.Println("data.def as abc:", abc2) // data.def as abc: {1981}
	
	// Special case: convert value to 'time.Time'
	v,_ := reddo.ToTime(1547549353)
	fmt.Println(v) // 2019-01-15 17:49:13 +0700 +07
	v,_ = reddo.ToTime("1547549353123")
	fmt.Println(v) // 2019-01-15 17:49:13.123 +0700 +07
}
```


## Documentation

See [GoDoc](https://godoc.org/github.com/btnguyen2k/consu/reddo).


## History

### 2019-02-12 - v0.1.1.2

- New (semi)constants `ZeroMap` and `ZeroSlice`
- Fix: to solve the case "convert to `interface{}`"
  - Function `Convert(v interface{}, t interface{}) (interface{}, error)` returns `(v, nil)` if `t` is `nil`


### 2019-02-11 - v0.1.1.1

- New constant `ZeroUint64`


### 2019-01-15 - v0.1.1

- `ToStruct(interface{}, interface{}) (interface{}, error)` & new function `ToTime(interface{}) (time.Time, error)`:
  - Add special case when converting to `time.Time`
  - Add global value `ZeroTime`
  - Fix a bug when converting a unexported field


### 2019-01-12 - v0.1.0

First release:
- Convert primitive types (`bool`, `float*`, `int*`, `uint*`, `string`)
- Convert `struct`, `array/slice` and `map`
- Convert pointer
