# Olaf

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/consu)](https://goreportcard.com/report/github.com/btnguyen2k/consu)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/btnguyen2k/consu/olaf)](https://pkg.go.dev/github.com/btnguyen2k/consu/olaf)
[![Actions Status](https://github.com/btnguyen2k/consu/workflows/olaf/badge.svg)](https://github.com/btnguyen2k/consu/actions)
[![codecov](https://codecov.io/gh/btnguyen2k/consu/branch/olaf/graph/badge.svg)](https://app.codecov.io/gh/btnguyen2k/consu/tree/olaf/olaf)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

Go implementation of Twitter Snowflake.

## Getting Started

### Install Package

```
go get -u github.com/btnguyen2k/olaf
```

### Usage

```go
package main

import (
    "fmt"

    "github.com/btnguyen2k/consu/olaf"
)

func main() {
    // use default epoch
    o := olaf.NewOlaf(1981)

    // use custom epoch (note: epoch is in milliseconds)
    // o := olaf.NewOlafWithEpoch(103, 1546543604123)

    id64 := o.Id64()
    id64Hex := o.Id64Hex()
    id64Ascii := o.Id64Ascii()
    fmt.Println("ID 64-bit (int)   : ", id64, " / Timestamp: ", o.ExtractTime64(id64))
    fmt.Println("ID 64-bit (hex)   : ", id64Hex, " / Timestamp: ", o.ExtractTime64Hex(id64Hex))
    fmt.Println("ID 64-bit (ascii) : ", id64Ascii, " / Timestamp: ", o.ExtractTime64Ascii(id64Ascii))

    id128 := o.Id128()
    id128Hex := o.Id128Hex()
    id128Ascii := o.Id128Ascii()
    fmt.Println("ID 128-bit (int)  : ", id128.String(), " / Timestamp: ", o.ExtractTime128(id128))
    fmt.Println("ID 128-bit (hex)  : ", id128Hex, " / Timestamp: ", o.ExtractTime128Hex(id128Hex))
    fmt.Println("ID 128-bit (ascii): ", id128Ascii, " / Timestamp: ", o.ExtractTime128Ascii(id128Ascii))
}
```

## License

MIT - see [LICENSE.md](LICENSE.md).
