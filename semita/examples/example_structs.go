package main

import (
	"encoding/json"
	"fmt"

	"github.com/btnguyen2k/consu/semita"
)

func testReadStructs(s *semita.Semita) {
	fmt.Println("-========== Semina demo: Structs - READ ==========-")
	var path string
	var v interface{}
	var e error

	j, _ := json.Marshal(s.Unwrap())
	fmt.Printf("Data: %v\n", string(j))

	path = "privateName"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

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

	semita.PathSeparator = '.'
	path = "Employees[1].Email"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	semita.PathSeparator = '/'
	path = "Employees/[0]/Options/Overtime"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	semita.PathSeparator = ':'
	path = "Employees[2]:Age"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
}

func testWriteStructs(s *semita.Semita) {
	fmt.Println("-========== Semina demo: Structs - WRITE ==========-")
	var path string
	var v interface{}
	var e error

	j, _ := json.Marshal(s.Unwrap())
	fmt.Printf("Data: %v\n", string(j))

	// set new value to an exiting node (map's entry)
	path = "Employees[0].Age"
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
	l := len(v.([]Employee))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}
	s.SetValue("Employees[].Name", "New Employee") // does not work if the wrapped struct is not addressable
	v, e = s.GetValue(path)
	l = len(v.([]Employee))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tNew value at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}

	// append new item to slice
	path = "Employees[1].Options.WorkHours"
	v, e = s.GetValue(path)
	l = len(v.([]int))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}
	s.SetValue(path+".[]", 999) // this works for nested slice event if the wrapped struct is not addressable
	v, e = s.GetValue(path)
	l = len(v.([]int))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tNew value at path %v (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at %v: %e\n", path, e)
	}
}

func exampleStructs() {
	data1 := sampleDataStructs()
	s1 := semita.NewSemita(data1) // wrap around data
	semita.PathSeparator = '.'    // reset path separator
	testReadStructs(s1)
	semita.PathSeparator = '.' // reset path separator
	testWriteStructs(s1)

	data2 := sampleDataStructs()
	s2 := semita.NewSemita(&data2) // wrap around a pointer to data
	semita.PathSeparator = '.'     // reset path separator
	testReadStructs(s2)
	semita.PathSeparator = '.' // reset path separator
	testWriteStructs(s2)
}
