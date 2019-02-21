package main

import "time"

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
	// Options is struct used by examples
	Options struct {
		WorkHours []int
		Overtime  bool
	}
	// Employee is struct used by examples
	Employee struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
		Options   Options
		JoinDate  time.Time
	}
	// Company is struct used by examples
	Company struct {
		privateName string
		Name        string
		Year        int
		Employees   []Employee
	}

	// OptionsMixed is struct used by examples
	OptionsMixed struct {
		WorkHours []int
		Overtime  bool
	}
	// CompanyMixed is struct used by examples
	CompanyMixed struct {
		Name      string
		Year      int
		Employees []map[string]interface{}
	}
)

// generate sample data where root node is a map, with nested maps and slices
func sampleDataMapsAndSlices() map[string]interface{} {
	d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
	d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)
	_employee0WorkHours := make([]int, len(employee0WorkHours))
	copy(_employee0WorkHours, employee0WorkHours)
	_employee1WorkHours := make([]int, len(employee1WorkHours))
	copy(_employee1WorkHours, employee1WorkHours)
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
					"work_hours": _employee0WorkHours,
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
					"work_hours": _employee1WorkHours,
					"overtime":   employee1Overtime,
				},
				"join_date": d1,
			},
		},
	}
}

// generate sample data where root node is a struct, with nested structs and slices
func sampleDataStructs() Company {
	d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
	d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)
	return Company{
		privateName: "private-" + companyName,
		Name:        companyName,
		Year:        companyYear,
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

// generate sample data where root node is a struct, with nested structs, maps and slices
func sampleDataMixed() interface{} {
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
					Overtime:  employee0Overtime,
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
					Overtime:  employee1Overtime,
				},
				"join_date": d1,
			},
		},
	}
}
