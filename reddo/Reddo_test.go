package reddo

import (
	"testing"
)

func ifFailed(t *testing.T, f string, e error) {
	if e != nil {
		t.Errorf("%s failed: %e", f, e)
	}
}

/*----------------------------------------------------------------------*/

func testToBool(t *testing.T, input interface{}, expected bool) {
	v, e := ToBool(input)
	ifFailed(t, "TestToBool", e)
	if v != expected {
		t.Errorf("TestToBool failed: expected [%v] but received [%v]", expected, v)
	}
}

// TestToBool tests if values are converted correctly to bool
func TestToBool(t *testing.T) {
	var inputList = []interface{}{false, true}
	var expectedList = []bool{false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []bool{false, true, true, false, true, true, false, true, true, false, true, true, false, true, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []bool{false, true, false, true, false, true, false, true, false, true, false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []bool{false, true, true, true, true, false, true, true, true, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{0 + 0i, 0 - 0i, 0 + 2i, 0 - 3i, -1 + 0i, 2 + 0i, 3 - 2i, 3 + 3i, -4 + 5i, -5 + 6i}
	expectedList = []bool{false, false, true, true, true, true, true, true, true, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i])
	}

	var i = 0
	var p1 *int
	var p2 = &i
	inputList = []interface{}{p1, p2}
	expectedList = []bool{false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{"false", "true", "False", "True", "FALSE", "TRUE"}
	expectedList = []bool{false, true, false, true, false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i])
	}

	{
		input := "blabla"
		_, e := ToBool(input)
		if e == nil {
			t.Errorf("TestToBool failed: [%v] should not be convertable to bool!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToBool(input)
		if e == nil {
			t.Errorf("TestToBool failed: [%v] should not be convertable to bool!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

// const epsilon = 1E-9

func testToFloat(t *testing.T, input interface{}, expected float64) {
	v, e := ToFloat(input)
	ifFailed(t, "TestToFloat", e)
	if v != expected {
		t.Errorf("TestToFloat failed: expected [%f] but received [%f]", expected, v)
	}
	// delta := v.(float64) - expected
	// if math.Abs(delta) > epsilon {
	// 	t.Errorf("TestToFloat failed: expected [%f] but received [%f]", expected, v)
	// }
}

// TestToFloat tests if values are converted correctly to float
func TestToFloat(t *testing.T) {
	var inputList = []interface{}{false, true}
	var expectedList = []float64{0.0, 1.0}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []float64{0.0, -1.0, 2.0, 0.0, -2.0, 3.0, 0.0, -3.0, 4.0, 0.0, -4.0, 5.0, 0.0, -5.0, 6.0}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []float64{0.0, 1.0, 0.0, 2.0, 0.0, 3.0, 0.0, 4.0, 0.0, 5.0, 0.0, 6.0}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []float64{0.0, 0.001, -0.001, -1.2, 3.4, 0.0, 0.001, -0.001, -5.6, 7.8}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{"0", "0.0", "0.001", "-0.001", "-1.2", "3.4", "-1E9", "1e9", "-1e-9", "1E-9"}
	expectedList = []float64{0.0, 0.0, 0.001, -0.001, -1.2, 3.4, -1e9, 1E9, -1E-9, 1e-9}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i])
	}

	{
		input := "blabla"
		_, e := ToFloat(input)
		if e == nil {
			t.Errorf("TestToFloat failed: [%v] should not be convertable to float!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToFloat(input)
		if e == nil {
			t.Errorf("TestToFloat failed: [%v] should not be convertable to float!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

func testToInt(t *testing.T, input interface{}, expected int64) {
	v, e := ToInt(input)
	ifFailed(t, "TestToInt", e)
	if v != expected {
		t.Errorf("TestToInt failed: expected [%d] but received [%d]", expected, v)
	}
}

// TestToInt tests if values are converted correctly to int
func TestToInt(t *testing.T) {
	var inputList = []interface{}{false, true}
	var expectedList = []int64{0, 1}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []int64{0, -1, 2, 0, -2, 3, 0, -3, 4, 0, -4, 5, 0, -5, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []int64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []int64{0, 0, -0, -1, 3, 0, 0, -0, -5, 7}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{"0", "-1", "2", "-3", "4"}
	expectedList = []int64{0, -1, 2, -3, 4}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i])
	}

	{
		input := "blabla"
		_, e := ToInt(input)
		if e == nil {
			t.Errorf("TestToInt failed: [%v] should not be convertable to int!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToInt(input)
		if e == nil {
			t.Errorf("TestToInt failed: [%v] should not be convertable to int!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

const (
	MaxUint = ^uint64(0)
	// MinUint = 0
	// MaxInt  = int64(^uint64(0) >> 1)
	// MinInt  = -MaxInt - 1
)

func testToUint(t *testing.T, input interface{}, expected uint64) {
	v, e := ToUint(input)
	ifFailed(t, "TestToUint", e)
	if v != expected {
		t.Errorf("TestToUint failed: expected [%d] but received [%d]", expected, v)
	}
}

// TestToUint tests if values are converted correctly to uint
func TestToUint(t *testing.T) {
	var inputList = []interface{}{false, true}
	var expectedList = []uint64{0, 1}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []uint64{0, MaxUint, 2, 0, MaxUint - 1, 3, 0, MaxUint - 2, 4, 0, MaxUint - 3, 5, 0, MaxUint - 4, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []uint64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []uint64{0, 0, 0, MaxUint, 3, 0, 0, 0, MaxUint - 4, 7}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{"0", "1", "2", "3", "4"}
	expectedList = []uint64{0, 1, 2, 3, 4}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i])
	}

	{
		input := "blabla"
		_, e := ToUint(input)
		if e == nil {
			t.Errorf("TestToUint failed: [%v] should not be convertable to uint!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToUint(input)
		if e == nil {
			t.Errorf("TestToUint failed: [%v] should not be convertable to uint!", input)
		}
	}
}
