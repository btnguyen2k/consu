# consu/semver

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/semver)](https://pkg.go.dev/github.com/btnguyen2k/consu/semver)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/semver/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/semver/graph/badge.svg)](https://app.codecov.io/gh/btnguyen2k/consu/tree/semver/semver)

Package `semver` provides utility functions to work with semantic versioning.

## Installation

```shell
$ go get -u github.com/btnguyen2k/consu/semver
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/btnguyen2k/consu/semver"
)

func main() {
	input := "2.0.0-rc.1+build.123"
	myVer := semver.ParseSemver(input)
    fmt.Printf("Version: %v\n", myVer)

	otherVer := semver.ParseSemver("2.0.1")
	fmt.Printf("My version vs other version: %v\n", myVer.Compare(otherVer))
}
```

## Features

⭐ Parse semantic versioning string, following [Semantic Versioning 2.0.0](Semantic Versioning 2.0.0) spec.

⭐ Helper method to compare two semantic versions.

## License

MIT - see [LICENSE](LICENSE.md).
