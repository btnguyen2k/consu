/*
Package semver provides utility functions to work with semantic versioning.

Sample usage:

	package main

	import "fmt"
	import "github.com/btnguyen2k/consu/semver"

	func main() {
		v := semver.ParseSemver("1.0.0")
		fmt.Printf("Version: %#v\n",v)
	}
*/
package semver

const (
	// Version defines version number of this package
	Version = "0.1.1"
)
