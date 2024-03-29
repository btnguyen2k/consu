package g18

import (
	"reflect"
	"testing"
	"time"
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

func TestDeduplicate_float32(t *testing.T) {
	testName := "TestDeduplicate_float32"
	testData := []struct {
		input    []float32
		expected []float32
	}{
		{
			nil, []float32{},
		},
		{
			[]float32{}, []float32{},
		},
		{
			[]float32{1.2, 3.4, 2.3}, []float32{1.2, 2.3, 3.4},
		},
		{
			[]float32{3.4, 1.2, 3.4, 2.3}, []float32{1.2, 2.3, 3.4},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicate_float64(t *testing.T) {
	testName := "TestDeduplicate_float64"
	testData := []struct {
		input    []float64
		expected []float64
	}{
		{
			nil, []float64{},
		},
		{
			[]float64{}, []float64{},
		},
		{
			[]float64{1.2, 3.4, 2.3}, []float64{1.2, 2.3, 3.4},
		},
		{
			[]float64{3.4, 1.2, 3.4, 2.3}, []float64{1.2, 2.3, 3.4},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicate_byte(t *testing.T) {
	testName := "TestDeduplicate_byte"
	testData := []struct {
		input    []byte
		expected []byte
	}{
		{
			nil, []byte{},
		},
		{
			[]byte{}, []byte{},
		},
		{
			[]byte{1, 3, 2}, []byte{1, 2, 3},
		},
		{
			[]byte{3, 1, 3, 2}, []byte{1, 2, 3},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicate_rune(t *testing.T) {
	testName := "TestDeduplicate_rune"
	testData := []struct {
		input    []rune
		expected []rune
	}{
		{
			nil, []rune{},
		},
		{
			[]rune{}, []rune{},
		},
		{
			[]rune{'1', '3', '2'}, []rune{'1', '2', '3'},
		},
		{
			[]rune{'3', '1', '3', '2'}, []rune{'1', '2', '3'},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicate_timeDuration(t *testing.T) {
	testName := "TestDeduplicate_timeDuration"
	testData := []struct {
		input    []time.Duration
		expected []time.Duration
	}{
		{
			nil, []time.Duration{},
		},
		{
			[]time.Duration{}, []time.Duration{},
		},
		{
			[]time.Duration{1, 3, 2}, []time.Duration{1, 2, 3},
		},
		{
			[]time.Duration{3, 1, 3, 2}, []time.Duration{1, 2, 3},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicate_uintptr(t *testing.T) {
	testName := "TestDeduplicate_uintptr"
	testData := []struct {
		input    []uintptr
		expected []uintptr
	}{
		{
			nil, []uintptr{},
		},
		{
			[]uintptr{}, []uintptr{},
		},
		{
			[]uintptr{1, 3, 2}, []uintptr{1, 2, 3},
		},
		{
			[]uintptr{3, 1, 3, 2}, []uintptr{1, 2, 3},
		},
	}
	for _, td := range testData {
		result := Deduplicate(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_bool(t *testing.T) {
	testName := "TestDeduplicateStable_bool"
	testData := []struct {
		input    []bool
		expected []bool
	}{
		{
			nil, []bool{},
		},
		{
			[]bool{}, []bool{},
		},
		{
			[]bool{true, false}, []bool{true, false},
		},
		{
			[]bool{false, true}, []bool{false, true},
		},
		{
			[]bool{true, true, false, false, true}, []bool{true, false},
		},
		{
			[]bool{false, false, true, true, false}, []bool{false, true},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_string(t *testing.T) {
	testName := "TestDeduplicateStable_string"
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
			[]string{"1", "3", "2"}, []string{"1", "3", "2"},
		},
		{
			[]string{"3", "1", "3", "2"}, []string{"3", "1", "2"},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_int(t *testing.T) {
	testName := "TestDeduplicateStable_int"
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
			[]int{1, 3, 2}, []int{1, 3, 2},
		},
		{
			[]int{3, 1, 3, 2}, []int{3, 1, 2},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_uint(t *testing.T) {
	testName := "TestDeduplicateStable_uint"
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
			[]uint{1, 3, 2}, []uint{1, 3, 2},
		},
		{
			[]uint{3, 1, 3, 2}, []uint{3, 1, 2},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_float32(t *testing.T) {
	testName := "TestDeduplicateStable_float32"
	testData := []struct {
		input    []float32
		expected []float32
	}{
		{
			nil, []float32{},
		},
		{
			[]float32{}, []float32{},
		},
		{
			[]float32{1.2, 3.4, 2.3}, []float32{1.2, 3.4, 2.3},
		},
		{
			[]float32{3.4, 1.2, 3.4, 2.3}, []float32{3.4, 1.2, 2.3},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_float64(t *testing.T) {
	testName := "TestDeduplicateStable_float64"
	testData := []struct {
		input    []float64
		expected []float64
	}{
		{
			nil, []float64{},
		},
		{
			[]float64{}, []float64{},
		},
		{
			[]float64{1.2, 3.4, 2.3}, []float64{1.2, 3.4, 2.3},
		},
		{
			[]float64{3.4, 1.2, 3.4, 2.3}, []float64{3.4, 1.2, 2.3},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_complex64(t *testing.T) {
	testName := "TestDeduplicateStable_complex64"
	testData := []struct {
		input    []complex64
		expected []complex64
	}{
		{
			nil, []complex64{},
		},
		{
			[]complex64{}, []complex64{},
		},
		{
			[]complex64{1.2i, 3.4i, 2.3i}, []complex64{1.2i, 3.4i, 2.3i},
		},
		{
			[]complex64{3.4i, 1.2i, 3.4i, 2.3i}, []complex64{3.4i, 1.2i, 2.3i},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_complex128(t *testing.T) {
	testName := "TestDeduplicateStable_complex128"
	testData := []struct {
		input    []complex128
		expected []complex128
	}{
		{
			nil, []complex128{},
		},
		{
			[]complex128{}, []complex128{},
		},
		{
			[]complex128{1.2i, 3.4i, 2.3i}, []complex128{1.2i, 3.4i, 2.3i},
		},
		{
			[]complex128{3.4i, 1.2i, 3.4i, 2.3i}, []complex128{3.4i, 1.2i, 2.3i},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_byte(t *testing.T) {
	testName := "TestDeduplicateStable_byte"
	testData := []struct {
		input    []byte
		expected []byte
	}{
		{
			nil, []byte{},
		},
		{
			[]byte{}, []byte{},
		},
		{
			[]byte{1, 3, 2}, []byte{1, 3, 2},
		},
		{
			[]byte{3, 1, 3, 2}, []byte{3, 1, 2},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_rune(t *testing.T) {
	testName := "TestDeduplicateStable_rune"
	testData := []struct {
		input    []rune
		expected []rune
	}{
		{
			nil, []rune{},
		},
		{
			[]rune{}, []rune{},
		},
		{
			[]rune{'1', '3', '2'}, []rune{'1', '3', '2'},
		},
		{
			[]rune{'3', '1', '3', '2'}, []rune{'3', '1', '2'},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_timeDuration(t *testing.T) {
	testName := "TestDeduplicateStable_timeDuration"
	testData := []struct {
		input    []time.Duration
		expected []time.Duration
	}{
		{
			nil, []time.Duration{},
		},
		{
			[]time.Duration{}, []time.Duration{},
		},
		{
			[]time.Duration{1, 3, 2}, []time.Duration{1, 3, 2},
		},
		{
			[]time.Duration{3, 1, 3, 2}, []time.Duration{3, 1, 2},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

func TestDeduplicateStable_uintptr(t *testing.T) {
	testName := "TestDeduplicateStable_uintptr"
	testData := []struct {
		input    []uintptr
		expected []uintptr
	}{
		{
			nil, []uintptr{},
		},
		{
			[]uintptr{}, []uintptr{},
		},
		{
			[]uintptr{1, 3, 2}, []uintptr{1, 3, 2},
		},
		{
			[]uintptr{3, 1, 3, 2}, []uintptr{3, 1, 2},
		},
	}
	for _, td := range testData {
		result := DeduplicateStable(td.input)
		if !reflect.DeepEqual(result, td.expected) {
			t.Fatalf("%s failed: {test data: %#v / expected: %#v / received: %#v}", testName, td.input, td.expected, result)
		}
	}
}

/*----------------------------------------------------------------------*/

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

/*----------------------------------------------------------------------*/

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

/*----------------------------------------------------------------------*/
func TestMinMax_string(t *testing.T) {
	testName := "TestMinMax_string"
	testData := []struct {
		input    []string
		min, max string
	}{
		{
			input: []string{"1", "3", "2"}, min: "1", max: "3",
		},
		{
			input: []string{"3", "1", "3", "2"}, min: "1", max: "3",
		},
	}
	for _, td := range testData {
		mi := Min(td.input...)
		ma := Max(td.input...)
		if mi != td.min || ma != td.max {
			t.Fatalf("%s failed: {test data: %#v / expected: (mi: %#v, ma: %#v) / received: (mi: %#v, ma: %#v)}", testName, td.input, td.min, td.max, mi, ma)
		}
	}
}

func TestMinMax_int(t *testing.T) {
	testName := "TestMinMax_int"
	testData := []struct {
		input    []int
		min, max int
	}{
		{
			input: []int{1, 3, 2}, min: 1, max: 3,
		},
		{
			input: []int{3, 1, 3, 2}, min: 1, max: 3,
		},
	}
	for _, td := range testData {
		mi := Min(td.input...)
		ma := Max(td.input...)
		if mi != td.min || ma != td.max {
			t.Fatalf("%s failed: {test data: %#v / expected: (mi: %#v, ma: %#v) / received: (mi: %#v, ma: %#v)}", testName, td.input, td.min, td.max, mi, ma)
		}
	}
}

func TestMinMax_uint(t *testing.T) {
	testName := "TestMinMax_uint"
	testData := []struct {
		input    []uint
		min, max uint
	}{
		{
			input: []uint{1, 3, 2}, min: 1, max: 3,
		},
		{
			input: []uint{3, 1, 3, 2}, min: 1, max: 3,
		},
	}
	for _, td := range testData {
		mi := Min(td.input...)
		ma := Max(td.input...)
		if mi != td.min || ma != td.max {
			t.Fatalf("%s failed: {test data: %#v / expected: (mi: %#v, ma: %#v) / received: (mi: %#v, ma: %#v)}", testName, td.input, td.min, td.max, mi, ma)
		}
	}
}

func TestMinMax_float32(t *testing.T) {
	testName := "TestMinMax_float32"
	testData := []struct {
		input    []float32
		min, max float32
	}{
		{
			input: []float32{1.2, 3.4, 2.3}, min: 1.2, max: 3.4,
		},
		{
			input: []float32{3.4, 1.2, 3.4, 2.3}, min: 1.2, max: 3.4,
		},
	}
	for _, td := range testData {
		mi := Min(td.input...)
		ma := Max(td.input...)
		if mi != td.min || ma != td.max {
			t.Fatalf("%s failed: {test data: %#v / expected: (mi: %#v, ma: %#v) / received: (mi: %#v, ma: %#v)}", testName, td.input, td.min, td.max, mi, ma)
		}
	}
}

func TestMinMax_float64(t *testing.T) {
	testName := "TestMinMax_float64"
	testData := []struct {
		input    []float64
		min, max float64
	}{
		{
			input: []float64{1.2, 3.4, 2.3}, min: 1.2, max: 3.4,
		},
		{
			input: []float64{3.4, 1.2, 3.4, 2.3}, min: 1.2, max: 3.4,
		},
	}
	for _, td := range testData {
		mi := Min(td.input...)
		ma := Max(td.input...)
		if mi != td.min || ma != td.max {
			t.Fatalf("%s failed: {test data: %#v / expected: (mi: %#v, ma: %#v) / received: (mi: %#v, ma: %#v)}", testName, td.input, td.min, td.max, mi, ma)
		}
	}
}
