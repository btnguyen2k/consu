/*
Package olaf provides methods to generate unique ID using Twitter Snowflake algorithm.

Sample usage:

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
		fmt.Println("ID 64-bit (int)   : ", id64, " / timestamp: ", o.ExtractTime64(id64))
		fmt.Println("ID 64-bit (hex)   : ", id64Hex, " / timestamp: ", o.ExtractTime64Hex(id64Hex))
		fmt.Println("ID 64-bit (ascii) : ", id64Ascii, " / timestamp: ", o.ExtractTime64Ascii(id64Ascii))

		id128 := o.Id128()
		id128Hex := o.Id128Hex()
		id128Ascii := o.Id128Ascii()
		fmt.Println("ID 128-bit (int)  : ", id128.String(), " / timestamp: ", o.ExtractTime128(id128))
		fmt.Println("ID 128-bit (hex)  : ", id128Hex, " / timestamp: ", o.ExtractTime128Hex(id128Hex))
		fmt.Println("ID 128-bit (ascii): ", id128Ascii, " / timestamp: ", o.ExtractTime128Ascii(id128Ascii))
	}
*/
package olaf

const (
	// Version defines version number of this package
	Version = "0.1.3"
)
