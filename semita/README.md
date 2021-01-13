# consu/semita

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/semita)](https://pkg.go.dev/github.com/btnguyen2k/consu/semita)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/semita/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/semita/graph/badge.svg?token=PWSL21DE1D)](https://app.codecov.io/gh/btnguyen2k/consu/branch/semita)

Package `semita` provides utility functions to access data from a hierarchy structure.

## Installation

```shell
$ go get -u github.com/btnguyen2k/consu/semita
```


## Usage

A 'path' is used to specify the location of item in the hierarchy data. Sample of a path `Employees.[1].first_name`, where:
- `.` (the dot character): path separator
- `Name`: access attribute of a map/struct specified by 'Name'
- `[i]`: access i'th element of a slice/array (0-based)
- The dot right before `[]` can be omitted: `Employees[1].first_name` is equivalent to `Employees.[1].first_name`.

Notes:
- Supported nested arrays, slices, maps and structs.
- Struct's un-exported fields can be read, but not written.
- Unaddressable structs and arrays are read-only.

Example:

(more examples on [project repository](https://github.com/btnguyen2k/consu/tree/master/semita/examples)).

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/btnguyen2k/consu/reddo"
    "github.com/btnguyen2k/consu/semita"
)

func main() {
	fmt.Println("-========== Semina demo ==========-")
	data := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Thanh",
			"last":  "ngn",
		},
		"yob":   1981,
		"alias": []string{"btnguyen2k", "thanhnb"},
	}
	s := semita.NewSemita(data)
	var err error
	var v interface{}

	// current data tree: {"alias":["btnguyen2k","thanhnb"],"name":{"first":"Thanh","last":"ngn"},"yob":1981}
	tree := s.Unwrap()
	js, _ := json.Marshal(tree)
	fmt.Println("Data tree:", string(js))

	// get nested value
	v, err = s.GetValue("name.first") // v should be "Thanh"
	if err == nil {
		fmt.Println("Firstname:", v)
	} else {
		fmt.Println("Error:", err)
	}

	// set nested value
	err = s.SetValue("name.last", "Nguyen") // v should be "Nguyen" (instead of "ngn")
	if err == nil {
		v, err = s.GetValue("name.last")
		if err == nil {
			fmt.Println("Lastname:", v)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Error:", err)
	}

	// get a value and its type
	yob, err := s.GetValue("yob") // yob should be int(1981)
	if err == nil {
		fmt.Println("YOB:", yob.(int))
	} else {
		fmt.Println("Error:", err)
	}

	// get a value and type
	yob, err = s.GetValueOfType("yob", reddo.TypeUint) // yob should be uint64(1981)
	if err == nil {
		fmt.Println("YOB:", yob.(uint64)) // all uint types are returned as uint64
	} else {
		fmt.Println("Error:", err)
	}

	// append new item to end of slice
	err = s.SetValue("alias[]", "another")
	if err == nil {
		// either alias[2] or alias.[2] is accepted
		alias, err := s.GetValue("alias.[2]") // alias should be "another"
		if err == nil {
			fmt.Println("New Alias:", alias)
		} else {
			fmt.Println("Error:", err)
		}

		allAlias, err := s.GetValue("alias") // allAlias should be ["btnguyen2k","thanhnb","another"]
		if err == nil {
			fmt.Println("All Alias:", allAlias)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Error:", err)
	}

	// create missing nodes along the path
	err = s.SetValue("a.b[].c.d", true)
	if err == nil {
		// missing nodes should be created
		// data tree should be: {"a":{"b":[{"c":{"d":true}}]},"alias":["btnguyen2k","thanhnb","another"],"name":{"first":"Thanh","last":"Nguyen"},"yob":1981}
		tree := s.Unwrap()
		js, _ := json.Marshal(tree)
		fmt.Println("Data tree:", string(js))
	} else {
		fmt.Println("Error:", err)
	}
}
```


## History

### 2019-04-12 - v0.1.4.1

- Upgrade to `consu/reddo-v0.1.6`:
  - Return `zero` value when input is `nil`.


### 2019-04-04 - v0.1.4

- Migrate to Go modular design.


### 2019-03-07 - v0.1.2

- Upgrade to `consu/reddo-v0.1.3`:
  - New functions `GetTime(path string) (time.Time, error)` and `GetTimeWithLayout(path, layout string) (time.Time, error)`


### 2019-03-05 - v0.1.1

- Compatible with `consu/reddo-v0.1.2`

### 2019-02-22 - v0.1.0

First release:
- Supported nested arrays, slices, maps and structs.
- Struct's un-exported fields can be read, but not written.
- Unaddressable structs and arrays are read-only.
