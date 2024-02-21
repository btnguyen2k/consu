/*
Package reddo provides utility functions to convert values using Golang's reflection.

Sample usage:

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
		fmt.Printf("Id is %s, year is %d (%d)\n", id, year, yearUint)

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
		fmt.Println(v)                         // 2019-01-15 17:49:13 +0700 +07
		v,_ = reddo.ToTime("1547549353123")
		fmt.Println(v)                         // 2019-01-15 17:49:13.123 +0700 +07
	}
*/
package reddo

const (
	// Version defines version number of this package
	Version = "0.1.8"
)

// This file contains module's metadata only, which is package level documentation and module Version string.
// Module's code should go into other files.
