# consu/gjrc

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/gjrc)](https://pkg.go.dev/github.com/btnguyen2k/consu/gjrc)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/gjrc/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![GoCover](https://img.shields.io/badge/coverage-GoCover-YellowGreen.svg)](https://gocover.io/github.com/btnguyen2k/consu/gjrc)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/gjrc/graph/badge.svg?token=PWSL21DE1D)](https://app.codecov.io/gh/btnguyen2k/consu/branch/gjrc)

Package `gjrc` offers generic utilities to work with JSON-based RESTful API.

## Installation

```shell
$ go get -u github.com/btnguyen2k/consu/gjrc
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/btnguyen2k/consu/gjrc"
	"github.com/btnguyen2k/consu/reddo"
)

func main() {
	// // pre-build a http.Client
	// httpClient := &http.Client{}
	// client := NewGjrc(httpClient, 0)

	// or, a new http.Client is created with 10 seconds timeout
	client := NewGjrc(nil, 10*time.Second)

	url := "https://httpbin.org/post"
	resp := client.PostJson(url, map[string]interface{}{"key1": "value", "key2": 1, "key3": true})

	val1, err := resp.GetValueAsType("json.key1", reddo.TypeString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", val1) // output: "value"

	val2, err := resp.GetValueAsType("json.key2", reddo.TypeInt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", val2) // output: 2

	val3, err := resp.GetValueAsType("json.key3", reddo.TypeBool)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", val3) // output: true
}
```

## History

### 2020-11-01 - v0.1.1

`go.mod` fixed.


### 2020-11-01 - v0.1.0

First release.
