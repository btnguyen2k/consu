package main

import (
	"encoding/json"
	"fmt"
	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/consu/semita"
)

func exampleSimple() {
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
	yob, err = s.GetValueOfType("yob", reddo.ZeroUint) // yob should be uint64(1981)
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
