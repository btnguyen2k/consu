package main

import (
	"encoding/json"
	"fmt"

	"github.com/btnguyen2k/consu/semita"
)

func testReadMapsAndSlices(s *semita.Semita) {
	fmt.Println("-========== Semina demo: MapsAndSlices - READ ==========-")
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

	semita.PathSeparator = '.'
	path = "Employees[1].email"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	semita.PathSeparator = '/'
	path = "Employees/[0]/options/overtime"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	semita.PathSeparator = ':'
	path = "Employees[2]:age"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
}

func testWriteMapsAndSlices(s *semita.Semita) {
	fmt.Println("-========== Semina demo: MapsAndSlices - WRITE ==========-")
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

	// create a new node and set its value
	path = "Employees.[1].senior"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
	s.SetValue(path, true)
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
		fmt.Printf("\tValue at path '%v' (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
	s.SetValue("Employees[].name", "Mew Employee")
	v, e = s.GetValue(path)
	l = len(v.([]map[string]interface{}))
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tNew value at path '%v' (number of items: %v): %v\n", path, l, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}

	// create all nodes along the path
	path = "a.b[].c.d[]"
	s.SetValue(path, "value")
	j, _ = json.Marshal(s.Unwrap())
	fmt.Printf("\tNew Data: %v\n", string(j))

	// set new value to an exiting node (slice's entry)
	path = "Employees[0].options.work_hours"
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tValue at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
	s.SetValue(path+"[1]", 981)
	v, e = s.GetValue(path)
	if e == nil {
		j, _ := json.Marshal(v)
		fmt.Printf("\tNew value at path '%v': %v\n", path, string(j))
	} else {
		fmt.Printf("\tError while getting value at '%v': %e\n", path, e)
	}
}

func exampleMapsAndSlices() {
	data1 := sampleDataMapsAndSlices()
	s1 := semita.NewSemita(data1) // wrap around data
	semita.PathSeparator = '.'    // reset path separator
	testReadMapsAndSlices(s1)
	semita.PathSeparator = '.' // reset path separator
	testWriteMapsAndSlices(s1)

	data2 := sampleDataMapsAndSlices()
	s2 := semita.NewSemita(&data2) // wrap around a pointer to data
	semita.PathSeparator = '.'     // reset path separator
	testReadMapsAndSlices(s2)
	semita.PathSeparator = '.' // reset path separator
	testWriteMapsAndSlices(s2)
}
