# consu/reddo

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/reddo)](https://pkg.go.dev/github.com/btnguyen2k/consu/reddo)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/reddo/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/reddo/graph/badge.svg?token=PWSL21DE1D)](https://app.codecov.io/gh/btnguyen2k/consu/branch/reddo)

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


## History

### 2021-01-12 - v0.1.7

- Bug fix: ToMap causes panic if map key or value is nil.
- Bug fix: ToSlice causes panic if an element is nil.
- Add `ZeroMode`:
  - If `reddo.ZeroMode=true`: `zero` value is returned when input is `nil`.
  - If `reddo.ZeroMode=false`: error is returned by `ToBool`, `ToFloat`, `ToInt`, `ToUint`, `ToString`, `ToTime`/`ToTimeWithLayout` and `ToStruct`


### 2019-04-12 - v0.1.6

- Return `zero` value when input is `nil`:
  - `reddo.ToBool(...)` returns `false`
  - `reddo.ToFloat(...)`, `reddo.ToInt(...)` and `reddo.ToUint(...)` returns `0`
  - `reddo.ToStrirng(...)` returns `""`
  - `reddo.ToTime(...)` and `reddo.ToTimeWithLayout(...)` returns `time.Time{}`
  - `reddo.ToSlice(...)` returns `nil`
  - `reddo.ToMap(...)` returns `nil`
  - `reddo.ToPointer(...)` returns `nil`


### 2019-04-02 - v0.1.5

- `reddo.ToString(...)`: handle the case converting `[]byte` to `string`.
- `reddo.ToSlice(...)`: handle the case converting `string` to `[]byte`.


### 2019-04-01 - v0.1.4

- Migrate to Go modular design.


### 2019-03-07 - v0.1.3

- New function `ToTimeWithLayout(v interface{}, layout string) (time.Time, error)`


### 2019-03-05 - v0.1.2

- Refactoring:
  - `ToBool(...)` now returns `(bool, error)`
  - `ToFloat(...)` now returns `(float64, error)`
  - `ToInt(...)` now returns `(int64, error)`
  - `ToUint(...)` now returns `(uint64, error)`
  - `ToString(...)` now returns `(string, error)`
  - `ToStruct(...)` changes its parameters to `ToStruct(v interface{}, t reflect.Type)`. Supplied target type can be slice, array or an element or array/slice.
  - `ToMap(...)` changes its parameters to `ToMap(v interface{}, t reflect.Type)`.
  - `Convert(...)` changes its parameters to `Convert(v interface{}, t reflect.Type)`.
- Remove `Zero...`, add `Type...`
- Other fixes and enhancements

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
