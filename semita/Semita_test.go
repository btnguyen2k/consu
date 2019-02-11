package semita

import (
	"github.com/btnguyen2k/consu/reddo"
	"reflect"
	"testing"
	"time"
)

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

// TestNewSemita test if Semita instance can be created correctly.
func TestNewSemita(t *testing.T) {
	// only Array, Slice, Map and Struct can be wrapped
	{
		data := [3]int{1, 2, 3}
		s1 := NewSemita(data)
		s2 := NewSemita(&data)
		if s1 == nil || s2 == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := []string{"a", "b", "c"}
		s1 := NewSemita(data)
		s2 := NewSemita(&data)
		if s1 == nil || s2 == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := map[string]interface{}{}
		s1 := NewSemita(data)
		s2 := NewSemita(&data)
		if s1 == nil || s2 == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := struct {
			a int
			b string
			c bool
		}{a: 1, b: "2", c: true}
		s1 := NewSemita(data)
		s2 := NewSemita(&data)
		if s1 == nil || s2 == nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}

	{
		data := 1
		s1 := NewSemita(data)
		s2 := NewSemita(&data)
		if s1 != nil || s2 != nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := "string"
		s1 := NewSemita(data)
		s2 := NewSemita(&data)
		if s1 != nil || s2 != nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
	{
		data := false
		s1 := NewSemita(data)
		s2 := NewSemita(&data)
		if s1 != nil || s2 != nil {
			t.Errorf("TestNewSemita failed for data %#v", data)
		}
	}
}

func TestSemita_Unwrap(t *testing.T) {
	data := map[string]interface{}{}

	s1 := NewSemita(data)
	d1 := s1.Unwrap().(map[string]interface{})
	if !reflect.DeepEqual(data, d1) {
		t.Errorf("TestSemita_Unwrap failed for data %#v", data)
	}

	s2 := NewSemita(&data)
	d2 := s2.Unwrap().(map[string]interface{})
	if !reflect.DeepEqual(data, d2) {
		t.Errorf("TestSemita_Unwrap failed for data %#v", data)
	}
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
			t.Errorf("TestSemita_GetValueInvalid getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		data := struct {
			a string
			b int
			c bool
		}{
			a: "string",
			b: 1,
			c: true,
		}
		s := NewSemita(data)
		p := "[1]"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueInvalid getting value at [%#v] for data %#v", p, data)
		}
	}

	{
		data := [3]int{1, 2, 3}
		s := NewSemita(data)
		p := "1"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueInvalid getting value at [%#v] for data %#v", p, data)
		}
	}
	{
		data := []string{"1", "2", "3"}
		s := NewSemita(data)
		p := "1"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueInvalid getting value at [%#v] for data %#v", p, data)
		}
	}
}

func TestSemita_GetValueArray(t *testing.T) {
	v := genDataArray()
	s1 := NewSemita(v)
	s2 := NewSemita(&v)
	var p string
	var err error
	var n interface{}

	p = "abc"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}

	p = "[999]"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}

	p = "[]"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueArray failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"[4].[0]", "[5][1]", "[6].z", "[7].A.[0]", "[7].B[1]", "[7].M.z", "[7].S.s"} {
		n, err = s1.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueArray failed with data %#v at path {%#v}", v, p)
		}
		n, err = s2.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueArray failed with data %#v at path {%#v}", v, p)
		}
	}
}

func TestSemita_GetValueSlice(t *testing.T) {
	v := genDataSlice()
	s1 := NewSemita(v)
	s2 := NewSemita(&v)
	var p string
	var err error
	var n interface{}

	p = "abc"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[999]"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[]"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestSemita_GetValueSlice failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"[4].[0]", "[5][1]", "[6].z", "[7].A.[0]", "[7].B[1]", "[7].M.z", "[7].S.s"} {
		n, err = s1.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueSlice failed with data %#v at path {%#v}", v, p)
		}
		n, err = s2.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueSlice failed with data %#v at path {%#v}", v, p)
		}
	}
}

func TestSemita_GetValueMap(t *testing.T) {
	v := genDataMap()
	s1 := NewSemita(v)
	s2 := NewSemita(&v)
	var p string
	var err error
	var n interface{}

	p = "[-1]"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}

	p = "[999]"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}

	p = "[]"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}

	p = "not exist"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// non-exists entry
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// non-exists entry
		t.Errorf("TestSemita_GetValueMap failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"a.[0]", "b[1]", "m.z", "s.A.[0]", "s.B[1]", "s.M.z", "s.S.s"} {
		n, err = s1.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueMap failed with data %#v at path {%#v}", v, p)
		}
		n, err = s2.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueMap failed with data %#v at path {%#v}", v, p)
		}
	}
}

func TestSemita_GetValueStruct(t *testing.T) {
	v := genDataOuter()
	s1 := NewSemita(v)
	s2 := NewSemita(&v)
	var p string
	var err error
	var n interface{}

	p = "[-1]"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "[999]"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "[]"
	n, err = s1.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err == nil {
		// invalid type
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "not exist"
	n, err = s1.GetValue(p)
	if n != nil || err != nil {
		// non-exists entry
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}
	n, err = s2.GetValue(p)
	if n != nil || err != nil {
		// non-exists entry
		t.Errorf("TestSemita_GetValueStruct failed with data %#v at index {%#v}", v, p)
	}

	for _, p := range []string{"A.[0]", "B[1]", "M.z", "S.s", "private"} {
		n, err = s1.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueStruct failed with data %#v at path {%#v}", v, p)
		}
		n, err = s2.GetValue(p)
		if n == nil || err != nil {
			t.Errorf("TestSemita_GetValueStruct failed with data %#v at path {%#v}", v, p)
		}
	}
}

/*----------------------------------------------------------------------*/

func ifFailed(t *testing.T, f string, e error) {
	if e != nil {
		t.Errorf("%s failed: %e", f, e)
		t.FailNow()
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

type (
	Options struct {
		WorkHours []int
		Overtime  bool
	}
	Employee struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
		Options   Options
		JoinDate  time.Time
	}
	Company struct {
		Name      string
		Year      int
		Employees []Employee
	}

	OptionsMixed struct {
		WorkHours []int
		Overtime   bool
	}
	CompanyMixed struct {
		Name      string
		Year      int
		Employees []map[string]interface{}
	}
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
	d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
	d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)
	return Company{
		Name: companyName,
		Year: companyYear,
		Employees: []Employee{
			{
				FirstName: employee0FirstName,
				LastName:  employee0LastName,
				Email:     employee0Email,
				Age:       employee0Age,
				Options: Options{
					WorkHours: employee0WorkHours,
					Overtime:  employee0Overtime,
				},
				JoinDate: d0,
			},
			{
				FirstName: employee1FirstName,
				LastName:  employee1LastName,
				Email:     employee1Email,
				Age:       employee1Age,
				Options: Options{
					WorkHours: employee1WorkHours,
					Overtime:  employee1Overtime,
				},
				JoinDate: d1,
			},
		},
	}
}

func generateDataMixed() interface{} {
	d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
	d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)
	return CompanyMixed{
		Name: companyName,
		Year: companyYear,
		Employees: []map[string]interface{}{
			{
				"first_name": employee0FirstName,
				"last_name":  employee0LastName,
				"email":      employee0Email,
				"age":        employee0Age,
				"options": OptionsMixed{
					WorkHours: employee0WorkHours,
					Overtime:   employee0Overtime,
				},
				"join_date": d0,
			},
			{
				"first_name": employee1FirstName,
				"last_name":  employee1LastName,
				"email":      employee1Email,
				"age":        employee1Age,
				"options": OptionsMixed{
					WorkHours: employee1WorkHours,
					Overtime:   employee1Overtime,
				},
				"join_date": d1,
			},
		},
	}
}

/*----------------------------------------------------------------------*/

func TestSemita_GetValueOfType_MultiLevelMap(t *testing.T) {
	name := "TestSemita_GetValueOfType_MultiLevelMap"
	data := generateDataMap()
	s1 := NewSemita(data)
	d := data.(map[string]interface{})
	s2 := NewSemita(&d)

	{
		p := "Name"
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != companyName {
			t.Errorf("%s getting value at {%#v} for data {%#v}", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != companyName {
			t.Errorf("%s getting value at {%#v} for data {%#v}", name, p, data)
		}
	}
	{
		p := "Year"
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].age"
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].email"
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].options.work_hours"
		v, e := s1.GetValueOfType(p, []int{})
		ifFailed(t, name, e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
		v, e = s2.GetValueOfType(p, []int{})
		ifFailed(t, name, e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
	}
	{
		p := "Employees.[1].options.overtime"
		v, e := s1.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].join_date"
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].join_date"
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
}

func TestSemita_GetValueOfType_MultiLevelStruct(t *testing.T) {
	name := "TestSemita_GetValueOfType_MultiLevelStruct"
	data := generateDataStruct()
	s1 := NewSemita(data)
	d := data.(Company)
	s2 := NewSemita(&d)

	{
		p := "Name"
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Year"
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].Age"
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].Email"
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].Options.WorkHours"
		v, e := s1.GetValueOfType(p, []int{})
		ifFailed(t, name, e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
		v, e = s2.GetValueOfType(p, []int{})
		ifFailed(t, name, e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
	}
	{
		p := "Employees.[1].Options.Overtime"
		v, e := s1.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].JoinDate"
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].JoinDate"
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
}

func TestSemita_GetValueOfType_MultiLevelMixed(t *testing.T) {
	name := "TestSemita_GetValueOfType_MultiLevelMixed"
	data := generateDataMixed()
	s1 := NewSemita(data)
	d := data.(CompanyMixed)
	s2 := NewSemita(&d)

	{
		p := "Name"
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Year"
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].age"
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].email"
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].options.WorkHours"
		v, e := s1.GetValueOfType(p, []int{})
		ifFailed(t, name, e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
		v, e = s2.GetValueOfType(p, []int{})
		ifFailed(t, name, e)
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
	}
	{
		p := "Employees.[1].options.Overtime"
		v, e := s1.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].join_date"
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].join_date"
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
}

/*----------------------------------------------------------------------*/

func TestSemita_SetValue_MultiLevelMap(t *testing.T) {
	name := "TestSemita_SetValue_MultiLevelMap"
	data := generateDataMap()
	s1 := NewSemita(data)
	d := data.(map[string]interface{})
	s2 := NewSemita(&d)

	{
		p := "Name"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroUint)
		ifFailed(t, name, e)
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Year"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].age"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroUint)
		ifFailed(t, name, e)
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].email"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].options.work_hours.[0]"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroUint)
		ifFailed(t, name, e)
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[1].options.overtime"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].join_date"
		d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
		d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)

		vSet1 := d1
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := d0
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
}

func TestSemita_SetValue_MultiLevelStruct(t *testing.T) {
	name := "TestSemita_SetValue_MultiLevelStruct"
	data := generateDataStruct()
	s1 := NewSemita(data)
	d := data.(Company)
	s2 := NewSemita(&d)

	{
		p := "Name"

		// vSet1 := "1"
		// e := s1.SetValue(p, vSet1)
		// ifFailed(t, name, e)
		// v, e := s1.GetValueOfType(p, reddo.ZeroString)
		// ifFailed(t, name, e)
		// if v.(string) != vSet1 {
		// 	t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		// }

		vSet2 := "2"
		e := s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e := s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Year"

		// vSet1 := "1"
		// e := s1.SetValue(p, vSet1)
		// ifFailed(t, name, e)
		// v, e := s1.GetValueOfType(p, reddo.ZeroString)
		// ifFailed(t, name, e)
		// if v.(string) != vSet1 {
		// 	t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		// }

		vSet2 := 2
		e := s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e := s2.GetValueOfType(p, reddo.ZeroUint)
		ifFailed(t, name, e)
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].Age"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroUint)
		ifFailed(t, name, e)
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].Email"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroString)
		ifFailed(t, name, e)
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].Options.WorkHours.[0]"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroInt)
		ifFailed(t, name, e)
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroUint)
		ifFailed(t, name, e)
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[1].Options.Overtime"

		vSet1 := !employee1Overtime
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := !vSet1
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroBool)
		ifFailed(t, name, e)
		if v.(bool) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].JoinDate"
		d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
		d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)

		vSet1 := d1
		e := s1.SetValue(p, vSet1)
		ifFailed(t, name, e)
		v, e := s1.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := d0
		e = s2.SetValue(p, vSet2)
		ifFailed(t, name, e)
		v, e = s2.GetValueOfType(p, reddo.ZeroTime)
		ifFailed(t, name, e)
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
}
