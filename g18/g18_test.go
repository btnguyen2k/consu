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

func TestPointerOf_bool(t *testing.T) {
	testName := "TestPointerOf_bool"
	v1, expected := PointerOf(false), false
	if *v1 != expected {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v1)
	}

	v2, expected := PointerOf(true), true
	if *v2 != expected {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v2)
	}
}

func TestPointerOf_string(t *testing.T) {
	testName := "TestPointerOf_string"
	v, expected := PointerOf("a string"), "a string"
	if *v != expected {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v)
	}
}

func TestPointerOf_int(t *testing.T) {
	testName := "TestPointerOf_int"
	v, expected := PointerOf(int(-123)), int(-123)
	if *v != expected {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v)
	}
}

func TestPointerOf_uint(t *testing.T) {
	testName := "TestPointerOf_uint"
	v, expected := PointerOf(uint(123)), uint(123)
	if *v != expected {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v)
	}
}

func TestPointerOf_float(t *testing.T) {
	testName := "TestPointerOf_float"
	v1, expected1 := PointerOf(float32(-12.3)), float32(-12.3)
	if *v1 != expected1 {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected1, v1)
	}

	v2, expected2 := PointerOf(float64(4.56)), float64(4.56)
	if *v2 != expected2 {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected2, v2)
	}
}

func TestPointerOf_struct(t *testing.T) {
	testName := "TestPointerOf_struct"
	type mystruct struct {
		S  string
		i  int
		B  bool
		ui uint
		F1 float64
		f2 float32
	}
	v, expected := PointerOf(mystruct{
		S:  "a string",
		i:  -12,
		B:  true,
		ui: 34,
		F1: 5.6,
		f2: -0.78,
	}), mystruct{
		S:  "a string",
		i:  -12,
		B:  true,
		ui: 34,
		F1: 5.6,
		f2: -0.78,
	}
	if *v != expected {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v)
	}
}

func TestPointerOf_map(t *testing.T) {
	testName := "TestPointerOf_map"
	v, expected := PointerOf(map[string]interface{}{
		"S":  "a string",
		"i":  int(-12),
		"B":  true,
		"ui": uint(34),
		"F1": float64(5.6),
		"f2": float32(-0.78),
	}), map[string]interface{}{
		"S":  "a string",
		"i":  int(-12),
		"B":  true,
		"ui": uint(34),
		"F1": float64(5.6),
		"f2": float32(-0.78),
	}
	if !reflect.DeepEqual(*v, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v)
	}
}

func TestPointerOf_slice(t *testing.T) {
	testName := "TestPointerOf_slice"
	v, expected := PointerOf([]interface{}{
		"a string",
		int(-12),
		true,
		uint(34),
		float64(5.6),
		float32(-0.78),
	}), []interface{}{
		"a string",
		int(-12),
		true,
		uint(34),
		float64(5.6),
		float32(-0.78),
	}
	if !reflect.DeepEqual(*v, expected) {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, expected, v)
	}
}
