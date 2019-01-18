package semita

import (
	"github.com/btnguyen2k/consu/reddo"
	"testing"
	"time"
)

/*----------------------------------------------------------------------*/

// TestNewSemita test if Semita instance can be created correctly.
func TestNewSemita(t *testing.T) {
	// only Array, Slice, Map and Struct can be wrapped
	{
		data := struct {
			a int
			b string
			c bool
		}{a: 1, b: "2", c: true}
		s := NewSemita(data)
		if s == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := map[string]interface{}{}
		s := NewSemita(data)
		if s == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := [3]int{1, 2, 3}
		s := NewSemita(data)
		if s == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := []string{"a", "b", "c"}
		s := NewSemita(data)
		if s == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}

	{
		data := 1
		s := NewSemita(data)
		if s != nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := "string"
		s := NewSemita(data)
		if s != nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := false
		s := NewSemita(data)
		if s != nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
}

/*----------------------------------------------------------------------*/

func testSplitPath(t *testing.T, path string, expected []string) {
	tokens := SplitPath(path)
	if len(tokens) != len(expected) {
		t.Errorf("TestSplitPath failed for data [%s], expected %#v but received %#v.", path, expected, tokens)
	}
}

// TestSplitPath tests if a path is correctly split into components
func TestSplitPath(t *testing.T) {
	testSplitPath(t, "a.b.c.[i].d", []string{"a", "b", "c", "[i]", "d"})
	testSplitPath(t, "a.b.c[i].d", []string{"a", "b", "c", "[i]", "d"})
	testSplitPath(t, "a.b.c.[i].[j].d", []string{"a", "b", "c", "[i]", "[j]", "d"})
	testSplitPath(t, "a.b.c[i].[j].d", []string{"a", "b", "c", "[i]", "[j]", "d"})
	testSplitPath(t, "a.b.c[i][j].d", []string{"a", "b", "c", "[i]", "[j]", "d"})
	testSplitPath(t, "a.b.c.[i][j].d", []string{"a", "b", "c", "[i]", "[j]", "d"})
}

/*----------------------------------------------------------------------*/

func TestSemita_GetValueInvalid(t *testing.T) {
	{
		data := map[string]interface{}{
			"a": "string",
			"b": 1,
			"c": true,
		}
		s := NewSemita(data)
		p := "[1]"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueArray getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		data := [3]int{1, 2, 3}
		s := NewSemita(data)
		p := "1"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueArray getting value at [%#v] for data %#v", p, data)
		}
	}
}

func TestSemita_GetValueArray(t *testing.T) {
	data := [3]int{1, 2, 3}
	s := NewSemita(data)

	{
		// index out-of-bound
		p := "[-1]"
		v, e := s.GetValue(p)
		if e == nil && v != nil {
			t.Errorf("TestSemita_GetValueArray getting value at path {%#v} for data {%#v}", p, data)
		}
	}
	{
		// index out-of-bound
		p := "[3]"
		v, e := s.GetValue(p)
		if e == nil && v != nil {
			t.Errorf("TestSemita_GetValueArray getting value at path {%#v} for data {%#v}", p, data)
		}
	}

	{
		p := "[0]"
		v, e := s.GetValue(p)
		if e != nil || v != data[0] {
			t.Errorf("TestSemita_GetValueArray getting value at path {%#v} for data {%#v}", p, data)
		}
	}
	{
		p := "[a]"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueArray getting value at path {%#v} for data {%#v}", p, data)
		}
	}
}

func TestSemita_GetValueSlice(t *testing.T) {
	data := []string{"1", "2", "3"}
	s := NewSemita(data)

	{
		// index out-of-bound
		p := "[-1]"
		v, e := s.GetValue(p)
		if e == nil && v != nil {
			t.Errorf("TestSemita_GetValueSlice getting value at path {%#v} for data {%#v}", p, data)
		}
	}
	{
		// index out-of-bound
		p := "[3]"
		v, e := s.GetValue(p)
		if e == nil && v != nil {
			t.Errorf("TestSemita_GetValueSlice getting value at path {%#v} for data {%#v}", p, data)
		}
	}

	{
		p := "[0]"
		v, e := s.GetValue(p)
		if e != nil || v != data[0] {
			t.Errorf("TestSemita_GetValueSlice getting value at path {%#v} for data {%#v}", p, data)
		}
	}
	{
		p := "[a]"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueSlice getting value at path {%#v} for data {%#v}", p, data)
		}
	}
}

func TestSemita_GetValueMap(t *testing.T) {
	data := map[string]interface{}{
		"a": "string",
		"b": 1,
		"c": true,
	}
	s := NewSemita(data)

	{
		p := "a"
		v, e := s.GetValue(p)
		if e != nil || v != data[p] {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "b"
		v, e := s.GetValue(p)
		if e != nil || v != data[p] {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "b"
		v, e := s.GetValue(p)
		if e != nil || v != data[p] {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for data %#v", p, data)
		}
	}

	{
		p := "z"
		v, e := s.GetValue(p)
		if v != nil && e != nil {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for data %#v", p, data)
		}
	}
}

func TestSemita_GetValueStruct(t *testing.T) {
	type MyStruct struct {
		A string
		B int
		C bool
		x string // un-exported field
	}
	data := MyStruct{
		A: "string",
		B: 1,
		C: true,
		x: "another string",
	}
	s := NewSemita(data)

	{
		p := "A"
		v, e := s.GetValue(p)
		if e != nil || v != data.A {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "B"
		v, e := s.GetValue(p)
		if e != nil || v != data.B {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "C"
		v, e := s.GetValue(p)
		if e != nil || v != data.C {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "x"
		v, e := s.GetValue(p)
		if e != nil || v != data.x {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for data %#v", p, data)
		}
	}

	{
		p := "z"
		v, e := s.GetValue(p)
		if v != nil && e != nil {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for data %#v", p, data)
		}
	}
}

/*----------------------------------------------------------------------*/

func ifFailed(t *testing.T, f string, e error) {
	if e != nil {
		t.Errorf("%s failed: %e", f, e)
	}
}

var (
	companyName = "Monster Corp."
	companyYear = 2003

	employee0FirstName      = "Mike"
	employee0LastName       = "Wazowski"
	employee0Email          = "mike.wazowski@monster.com"
	employee0Age            = 29
	employee0WorkHours      = []int{9, 10, 11, 12, 13, 14, 15, 16}
	employee0Overtime       = false
	employee0JoinDate       = "Apr 29, 2011"
	employee0JoinDateFormat = "Jan 02, 2006"

	employee1FirstName      = "Sulley"
	employee1LastName       = "Sullivan"
	employee1Email          = "sulley.sullivan@monster.com"
	employee1Age            = 30
	employee1WorkHours      = []int{13, 14, 15, 16, 17, 18, 19, 20}
	employee1Overtime       = true
	employee1JoinDate       = "2012-03-01 01:30:15 PM"
	employee1JoinDateFormat = "2006-01-02 03:04:05 PM"
)

func generateDataMap() interface{} {
	d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
	d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)
	return map[string]interface{}{
		"Name": companyName,
		"Year": companyYear,
		"Employees": []map[string]interface{}{
			{
				"first_name": employee0FirstName,
				"last_name":  employee0LastName,
				"email":      employee0Email,
				"age":        employee0Age,
				"options": map[string]interface{}{
					"work_hours": employee0WorkHours,
					"overtime":   employee0Overtime,
				},
				"join_date": d0,
			},
			{
				"first_name": employee1FirstName,
				"last_name":  employee1LastName,
				"email":      employee1Email,
				"age":        employee1Age,
				"options": map[string]interface{}{
					"work_hours": employee1WorkHours,
					"overtime":   employee1Overtime,
				},
				"join_date": d1,
			},
		},
	}
}

func generateDataStruct() interface{} {
	type Options struct {
		work_hours []int
		overtime   bool
	}

	type Employee struct {
		first_name string
		last_name  string
		email      string
		age        int
		options    Options
		join_date  time.Time
	}

	type Company struct {
		Name      string
		Year      int
		Employees []Employee
	}
	d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
	d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)
	return Company{
		Name: companyName,
		Year: companyYear,
		Employees: []Employee{
			{
				first_name: employee0FirstName,
				last_name:  employee0LastName,
				email:      employee0Email,
				age:        employee0Age,
				options: Options{
					work_hours: employee0WorkHours,
					overtime:   employee0Overtime,
				},
				join_date: d0,
			},
			{
				first_name: employee1FirstName,
				last_name:  employee1LastName,
				email:      employee1Email,
				age:        employee1Age,
				options: Options{
					work_hours: employee1WorkHours,
					overtime:   employee1Overtime,
				},
				join_date: d1,
			},
		},
	}
}

/*----------------------------------------------------------------------*/

func TestSemita_GetValueOfTypeMultiLevelMap(t *testing.T) {
	data := generateDataMap()
	s := NewSemita(data)

	{
		p := "Name"
		v, e := s.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(string) != companyName {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Year"
		v, e := s.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees.[0].age"
		v, e := s.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees[1].email"
		v, e := s.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(string) != employee1Email {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees[0].options.work_hours"
		v, e := s.GetValueOfType(p, []int{})
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
			}
		}
	}
	{
		p := "Employees.[1].options.overtime"
		v, e := s.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(bool) != employee1Overtime {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees.[0].join_date"
		v, e := s.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees[1].join_date"
		v, e := s.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
}

func TestSemita_GetValueOfTypeMultiLevelStruct(t *testing.T) {
	data := generateDataStruct()
	s := NewSemita(data)

	{
		p := "Name"
		v, e := s.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(string) != companyName {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Year"
		v, e := s.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees.[0].age"
		v, e := s.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees[1].email"
		v, e := s.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(string) != employee1Email {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees[0].options.work_hours"
		v, e := s.GetValueOfType(p, []int{})
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
			}
		}
	}
	{
		p := "Employees.[1].options.overtime"
		v, e := s.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(bool) != employee1Overtime {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees.[0].join_date"
		v, e := s.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "Employees[1].join_date"
		v, e := s.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
}

func generateDataMixed() interface{} {
	type Options struct {
		work_hours []int
		overtime   bool
	}

	type Company struct {
		name      string
		year      int
		employees []map[string]interface{}
	}
	d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
	d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)
	return Company{
		name: companyName,
		year: companyYear,
		employees: []map[string]interface{}{
			{
				"first_name": employee0FirstName,
				"last_name":  employee0LastName,
				"email":      employee0Email,
				"age":        employee0Age,
				"options": Options{
					work_hours: employee0WorkHours,
					overtime:   employee0Overtime,
				},
				"join_date": d0,
			},
			{
				"first_name": employee1FirstName,
				"last_name":  employee1LastName,
				"email":      employee1Email,
				"age":        employee1Age,
				"options": Options{
					work_hours: employee1WorkHours,
					overtime:   employee1Overtime,
				},
				"join_date": d1,
			},
		},
	}
}

func TestSemita_GetValueOfTypeMultiLevelMixed(t *testing.T) {
	data := generateDataMixed()
	s := NewSemita(data)

	{
		p := "name"
		v, e := s.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(string) != companyName {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "year"
		v, e := s.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "employees.[0].age"
		v, e := s.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "employees[1].email"
		v, e := s.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(string) != employee1Email {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "employees[0].options.work_hours"
		v, e := s.GetValueOfType(p, []int{})
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
			}
		}
	}
	{
		p := "employees.[1].options.overtime"
		v, e := s.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(bool) != employee1Overtime {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "employees.[0].join_date"
		v, e := s.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		p := "employees[1].join_date"
		v, e := s.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, "TestSemita_GetValueOfTypeMultiLevelMap", e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("TestSemita_GetValueOfTypeMultiLevelMap getting value at [%#v] for data %#v", p, data)
		}
	}
}
