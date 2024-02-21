# consu/gjrc

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/gjrc)](https://pkg.go.dev/github.com/btnguyen2k/consu/gjrc)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/gjrc/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/gjrc/graph/badge.svg)](https://app.codecov.io/gh/btnguyen2k/consu/tree/gjrc/gjrc)

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
	"time"

	"github.com/btnguyen2k/consu/gjrc"
	"github.com/btnguyen2k/consu/reddo"
)

func main() {
	// // pre-build a http.Client
	// httpClient := &http.Client{}
	// client := NewGjrc(httpClient, 0)

	// or, a new http.Client is created with 10 seconds timeout
	client := gjrc.NewGjrc(nil, 10*time.Second)

	url := "https://httpbin.org/post"
	resp := client.PostJson(url, map[string]interface{}{"key1": "value1", "key2": 2, "key3": true})

	val1, err := resp.GetValueAsType("json.key1", reddo.TypeString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", val1) // output: "value1"

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

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Support and Contribution

Feel free to create [pull requests](https://github.com/btnguyen2k/consu/pulls) or [issues](https://github.com/btnguyen2k/consu/issues) to report bugs or suggest new features.
Please search the existing issues before filing new issues to avoid duplicates. For new issues, file your bug or feature request as a new issue.

If you find this project useful, please star it.
