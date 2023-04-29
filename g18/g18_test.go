package g18

import (
	"reflect"
	"testing"
)

func TestDeduplicate_string(t *testing.T) {
	testName := "TestDeduplicate_string"
	testData := []struct {
		input    []string
		expected []string
	}{
		{
			nil, []string{},
		},
		{
			[]string{}, []string{},
		},
		{
			[]string{"1", "3", "2"}, []string{"1", "2", "3"},
		},
		{
			[]string{"3", "1", "3", "2"}, []string{"1", "2", "3"},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicate_int(t *testing.T) {
	testName := "TestDeduplicate_int"
	testData := []struct {
		input    []int
		expected []int
	}{
		{
			nil, []int{},
		},
		{
			[]int{}, []int{},
		},
		{
			[]int{1, 3, 2}, []int{1, 2, 3},
		},
		{
			[]int{3, 1, 3, 2}, []int{1, 2, 3},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicate_uint(t *testing.T) {
	testName := "TestDeduplicate_uint"
	testData := []struct {
		input    []uint
		expected []uint
	}{
		{
			nil, []uint{},
		},
		{
			[]uint{}, []uint{},
		},
		{
			[]uint{1, 3, 2}, []uint{1, 2, 3},
		},
		{
			[]uint{3, 1, 3, 2}, []uint{1, 2, 3},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestFindInSlice_bool(t *testing.T) {
	testName := "TestFindInSlice_bool"
	haystack := []bool{true}
	testData := map[bool]int{
		false: -1,
		true:  0,
	}
	for needle, expected := range testData {
		if v := FindInSlice(needle, haystack); v != expected {
			t.Fatalf("%s failed: needle: %#v / expected %#v but received %#v", testName, needle, expected, v)
		}
	}
}

func TestFindInSlice_string(t *testing.T) {
	testName := "TestFindInSlice_string"
	haystack := []string{"1", "2", "3", "4", "5"}
	testData := map[string]int{
		"0": -1,
		"1": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": -1,
	}
	for needle, expected := range testData {
		if v := FindInSlice(needle, haystack); v != expected {
			t.Fatalf("%s failed: needle: %#v / expected %#v but received %#v", testName, needle, expected, v)
		}
	}
}

func TestFindInSlice_int(t *testing.T) {
	testName := "TestFindInSlice_int"
	haystack := []int{1, 2, 3, 4, 5}
	testData := map[int]int{
		0: -1,
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
		6: -1,
	}
	for needle, expected := range testData {
		if v := FindInSlice(needle, haystack); v != expected {
			t.Fatalf("%s failed: needle: %#v / expected %#v but received %#v", testName, needle, expected, v)
		}
	}
}

func TestFindInSlice_uint(t *testing.T) {
	testName := "TestFindInSlice_uint"
	haystack := []uint{1, 2, 3, 4, 5}
	testData := map[uint]int{
		0: -1,
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
		6: -1,
	}
	for needle, expected := range testData {
		if v := FindInSlice(needle, haystack); v != expected {
			t.Fatalf("%s failed: needle: %#v / expected %#v but received %#v", testName, needle, expected, v)
		}
	}
}
