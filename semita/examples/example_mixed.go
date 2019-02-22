package main

import (
	"encoding/json"
	"fmt"
	"github.com/btnguyen2k/consu/semita"
)

func testReadMixed(s *semita.Semita) {
	fmt.Println("-========== Semina demo: Mixed - READ ==========-")
	var path string
	var v interface{}
	var e error

	j, _ := json.Marshal(s.Unwrap())
	fmt.Printf("Data: %v\n", string(j))

	path = "Employees"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	path = "Employees[0]"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	path = "Employees[1].email"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	path = "Employees.[0].options.Overtime"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	path = "Employees[2].age"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
}

func testWriteMixed(s *semita.Semita) {
	fmt.Println("-========== Semina demo: Mixed - WRITE ==========-")
	var path string
	var v interface{}
	var e error

	j, _ := json.Marshal(s.Unwrap())
	fmt.Printf("Data: %v\n", string(j))

	// set new value to an exiting node (map's entry)
	path = "Employees[0].age"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
	s.SetValue(path, 123)
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tNew value at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	// append new item to slice
	path = "Employees"
	v, e = s.GetValue(path)
	l := len(v.([]map[string]interface{}))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}
	s.SetValue("Employees[].name", "New Employee") // does not work if the wrapped struct is not addressable
	v, e = s.GetValue(path)
	l = len(v.([]map[string]interface{}))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tNew value at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}

	// append new item to slice
	path = "Employees[1].options.WorkHours"
	v, e = s.GetValue(path)
	l = len(v.([]int))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}
	s.SetValue(path+".[]", 999) // this does not work because nested struct is not addressable
	v, e = s.GetValue(path)
	l = len(v.([]int))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tNew value at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}
}

func exampleMixed() {
	data1 := sampleDataMixed()
	s1 := semita.NewSemita(data1) // wrap around data
	// testReadMixed(s1)
	testWriteMixed(s1)

	data2 := sampleDataMixed()
	s2 := semita.NewSemita(&data2) // wrap around a pointer to data
	// testReadMixed(s2)
	testWriteMixed(s2)
}
