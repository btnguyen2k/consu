package reddo

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func testToBool(t *testing.T, input interface{}, expected bool) {
	name := "TestToBool"
	{
		v, e := ToBool(input)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != expected {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
	}
	if input != nil {
		v, e := Convert(input, TypeBool)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(bool) != expected {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
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
			t.Fatalf("TestToBool failed: [%#v] should not be convertible to bool!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeBool)
		if e == nil {
			t.Fatalf("TestToBool failed: [%#v] should not be convertible to bool!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToBool(input)
		if e == nil {
			t.Fatalf("TestToBool failed: [%#v] should not be convertible to bool!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeBool)
		if e == nil {
			t.Fatalf("TestToBool failed: [%#v] should not be convertible to bool!", input)
		}
	}

	{
		ZeroMode = true
		v, e := ToBool(nil)
		if v != false || e != nil {
			t.Fatalf("TestToBool failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = ToBool(nil)
		if e == nil {
			t.Fatalf("TestToBool failed: [nil] should not be convertible to bool when ZeroMode=false!")
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, TypeBool)
		if v != false || e != nil {
			t.Fatalf("TestToBool failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = Convert(nil, TypeBool)
		if e == nil {
			t.Fatalf("TestToBool failed: [nil] should not be convertible to bool when ZeroMode=false!")
		}
	}
}

/*----------------------------------------------------------------------*/

func testToFloat(t *testing.T, input interface{}, expected float64) {
	name := "TestToFloat"
	{
		v, e := ToFloat(input)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != expected {
			t.Fatalf("%s failed: expected [%f] but received [%f]", name, expected, v)
		}
	}
	if input != nil {
		v, e := Convert(input, TypeFloat)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(float64) != expected {
			t.Fatalf("%s failed: expected [%f] but received [%f]", name, expected, v)
		}
	}
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
	expectedList = []float64{0.0, 0.0, 0.001, -0.001, -1.2, 3.4, -1e9, 1e9, -1e-9, 1e-9}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i])
	}

	{
		input := "blabla"
		_, e := ToFloat(input)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertible to float!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeFloat)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertible to float!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToFloat(input)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertible to float!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeFloat)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertible to float!", input)
		}
	}

	{
		ZeroMode = true
		v, e := ToFloat(nil)
		if v != 0.0 || e != nil {
			t.Fatalf("TestToFloat failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = ToFloat(nil)
		if e == nil {
			t.Fatalf("TestToFloat failed: [nil] should not be convertible to float when ZeroMode=false!")
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, TypeFloat)
		if v != 0.0 || e != nil {
			t.Fatalf("TestToFloat failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = Convert(nil, TypeFloat)
		if e == nil {
			t.Fatalf("TestToFloat failed: [nil] should not be convertible to float when ZeroMode=false!")
		}
	}
}

/*----------------------------------------------------------------------*/

func testToInt(t *testing.T, input interface{}, expected int64) {
	name := "TestToInt"
	{
		v, e := ToInt(input)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != expected {
			t.Fatalf("%s failed: expected [%d] but received [%d]", name, expected, v)
		}
	}
	if input != nil {
		v, e := Convert(input, TypeInt)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(int64) != expected {
			t.Fatalf("%s failed: expected [%d] but received [%d]", name, expected, v)
		}
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
		input := "-1.2"
		_, e := ToInt(input)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertible to int!", input)
		}
	}
	{
		input := "3.4"
		_, e := Convert(input, TypeInt)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertible to int!", input)
		}
	}

	{
		input := "blabla"
		_, e := ToInt(input)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertible to int!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeInt)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertible to int!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToInt(input)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertible to int!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeInt)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertible to int!", input)
		}
	}

	{
		ZeroMode = true
		v, e := ToInt(nil)
		if v != 0 || e != nil {
			t.Fatalf("TestToInt failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = ToInt(nil)
		if e == nil {
			t.Fatalf("TestToInt failed: [nil] should not be convertible to int when ZeroMode=false!")
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, TypeInt)
		if v != int64(0) || e != nil {
			t.Fatalf("TestToInt failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = Convert(nil, TypeInt)
		if e == nil {
			t.Fatalf("TestToInt failed: [nil] should not be convertible to int when ZeroMode=false!")
		}
	}
}

/*----------------------------------------------------------------------*/

const (
	MaxUint = ^uint64(0)
)

func testToUint(t *testing.T, input interface{}, expected uint64) {
	name := "TestToUint"
	{
		v, e := ToUint(input)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != expected {
			t.Fatalf("%s failed: expected [%d] but received [%d]", name, expected, v)
		}
	}
	if input != nil {
		v, e := Convert(input, TypeUint)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(uint64) != expected {
			t.Fatalf("%s failed: expected [%d] but received [%d]", name, expected, v)
		}
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
		input := "-1"
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertible to uint!", input)
		}
	}
	{
		input := "-1.2"
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertible to uint!", input)
		}
	}
	{
		input := "3.4"
		_, e := Convert(input, TypeUint)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertible to uint!", input)
		}
	}

	{
		input := "blabla"
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertible to uint!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeUint)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertible to uint!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertible to uint!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeUint)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertible to uint!", input)
		}
	}

	{
		ZeroMode = true
		v, e := ToUint(nil)
		if v != uint64(0) || e != nil {
			t.Fatalf("TestToUint failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = ToUint(nil)
		if e == nil {
			t.Fatalf("TestToUint failed: [nil] should not be convertible to uint when ZeroMode=false!")
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, TypeUint)
		if v != uint64(0) || e != nil {
			t.Fatalf("TestToUint failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = Convert(nil, TypeUint)
		if e == nil {
			t.Fatalf("TestToUint failed: [nil] should not be convertible to uint when ZeroMode=false!")
		}
	}
}

/*----------------------------------------------------------------------*/

func testToString(t *testing.T, input interface{}, expected string) {
	name := "TestToString"
	{
		v, e := ToString(input)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != expected {
			t.Fatalf("%s failed: expected [%s] but received [%s]", name, expected, v)
		}
	}
	if input != nil {
		v, e := Convert(input, TypeString)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(string) != expected {
			t.Fatalf("%s failed: expected [%s] but received [%s]", name, expected, v)
		}
	}
}

// TestToString tests if values are converted correctly to string
func TestToString(t *testing.T) {
	var inputList = []interface{}{false, true}
	var expectedList = []string{"false", "true"}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []string{"0", "-1", "2", "0", "-2", "3", "0", "-3", "4", "0", "-4", "5", "0", "-5", "6"}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []string{"0", "1", "0", "2", "0", "3", "0", "4", "0", "5", "0", "6"}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i])
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []string{"0", "0.001", "-0.001", "-1.2", "3.4", "0", "0.001", "-0.001", "-5.6", "7.8"}
	for i, n := 0, len(inputList); i < n; i++ {
		var expected string
		v := reflect.ValueOf(inputList[i])
		if v.Kind() == reflect.Float32 {
			expected = strconv.FormatFloat(v.Float(), 'g', -1, 64)
		} else {
			expected = expectedList[i]
		}
		testToString(t, inputList[i], expected)
	}

	inputList = []interface{}{[]byte("a string"), "0", "-1", "2", "-3", "4", "a", "b", "c", ""}
	expectedList = []string{"a string", "0", "-1", "2", "-3", "4", "a", "b", "c", ""}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i])
	}

	{
		input := struct {
		}{}
		testToString(t, input, fmt.Sprint(input))
	}

	{
		ZeroMode = true
		v, e := ToString(nil)
		if v != "" || e != nil {
			t.Fatalf("TestToString failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = ToString(nil)
		if e == nil {
			t.Fatalf("TestToString failed: [nil] should not be convertible to string when ZeroMode=false!")
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, TypeString)
		if v != "" || e != nil {
			t.Fatalf("TestToString failed: %#v / %s", v, e)
		}
		ZeroMode = false
		_, e = Convert(nil, TypeString)
		if e == nil {
			t.Fatalf("TestToString failed: [nil] should not be convertible to string when ZeroMode=false!")
		}
	}
}

/*----------------------------------------------------------------------*/
// TestToTimeError tests if values are converted correctly to time.Time
func TestToTimeError(t *testing.T) {
	{
		input := -1
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertible to time.Time!", input)
		}
	}

	{
		input := "-1"
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertible to time.Time!", input)
		}
	}

	{
		input := "-1.abc"
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertible to time.Time!", input)
		}
	}

	{
		input := struct {
		}{}
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertible to time.Time!", input)
		}
	}
}

// TestToTimeStruct tests if time.Time are converted correctly to time.Time
func TestToTimeStruct(t *testing.T) {
	name := "TestToTimeStruct"

	{
		// convert 'time.Time' to 'time.Time'
		now := time.Now()
		input := now
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).UnixNano() != now.UnixNano() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
}

// TestToTimeInteger tests if integers are converted correctly to time.Time
func TestToTimeInteger(t *testing.T) {
	name := "TestToTimeInteger"

	{
		// convert 'long(seconds)' to 'time.Time'
		now := time.Now()
		input := now.Unix()
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(milliseconds)' to 'time.Time'
		now := time.Now()
		input := now.UnixNano() / 1000000
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(microseconds)' to 'time.Time'
		now := time.Now()
		input := now.UnixNano() / 1000
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(nanoseconds)' to 'time.Time'
		now := time.Now()
		input := now.UnixNano()
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
}

// TestToTimeString tests if strings are converted correctly to time.Time
func TestToTimeString(t *testing.T) {
	name := "TestToTimeString"

	{
		// convert 'long(seconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.Unix(), 10)
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(milliseconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.UnixNano()/1000000, 10)
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(microseconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.UnixNano()/1000, 10)
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(nanoseconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.UnixNano(), 10)
		v, e := ToStruct(input, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
}

// TestToTimeNil tests if nil is converted correctly to time.Time
func TestToTimeNil(t *testing.T) {
	name := "TestToTimeNil"

	{
		ZeroMode = true
		v, e := ToTime(nil)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != zeroTime.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, zeroTime, v)
		}

		ZeroMode = false
		_, e = ToTime(nil)
		if e == nil {
			t.Fatalf("%s failed: [nil] should not be convertible to time.Time when ZeroMode=false!", name)
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, TypeTime)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.(time.Time).Unix() != zeroTime.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, zeroTime, v)
		}

		ZeroMode = false
		_, e = Convert(nil, TypeTime)
		if e == nil {
			t.Fatalf("%s failed: [nil] should not be convertible to time.Time when ZeroMode=false!", name)
		}
	}
}

// TestToTimeWithLayout tests if strings are converted correctly to time.Time using layout
func TestToTimeWithLayout(t *testing.T) {
	name := "TestToTimeWithLayout"

	{
		ZeroMode = true
		v, e := ToTimeWithLayout(nil, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != zeroTime.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, zeroTime, v)
		}

		ZeroMode = false
		_, e = ToTimeWithLayout(nil, "")
		if e == nil {
			t.Fatalf("%s failed: [nil] should not be convertible to time.Time when ZeroMode=false!", name)
		}
	}

	{
		// convert 'long(seconds)' to 'time.Time'
		now := time.Now()
		input := now.Unix()
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
	{
		// convert 'long(seconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.Unix(), 10)
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(milliseconds)' to 'time.Time'
		now := time.Now()
		input := now.UnixNano() / 1000000
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
	{
		// convert 'long(milliseconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.UnixNano()/1000000, 10)
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(microseconds)' to 'time.Time'
		now := time.Now()
		input := now.UnixNano() / 1000
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
	{
		// convert 'long(microseconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.UnixNano()/1000, 10)
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// convert 'long(nanoseconds)' to 'time.Time'
		now := time.Now()
		input := now.UnixNano()
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
	{
		// convert 'long(nanoseconds)' to 'time.Time'
		now := time.Now()
		input := strconv.FormatInt(now.UnixNano(), 10)
		v, e := ToTimeWithLayout(input, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != now.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		// invalid input
		input := "abc"
		layout := "Jan"
		_, e := ToTimeWithLayout(input, layout)
		if e == nil {
			t.Fatalf("%s failed: %e", name, e)
		}
	}
	{
		// invalid layout
		input := "2019-01-01"
		layout := "month"
		_, e := ToTimeWithLayout(input, layout)
		if e == nil {
			t.Fatalf("%s failed: %e", name, e)
		}
	}
	{
		input := "2019-04-29T20:59:10"
		layout := "2006-01-02T15:04:05"
		expected := time.Date(2019, 04, 29, 20, 59, 10, 0, time.UTC)
		v, e := ToTimeWithLayout(input, layout)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != expected.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
	}
	{
		loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		input := "2019-04-29 20:59:10.165067 +0700 +07"
		layout := "2006-01-02 15:04:05.999999 -0700 -07"
		expected := time.Date(2019, 04, 29, 20, 59, 10, 0, loc)
		v, e := ToTimeWithLayout(input, layout)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != expected.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
	}
}

// TestToStruct tests if values are converted correctly to struct
func TestToStruct(t *testing.T) {
	name := "TestToStruct"

	{
		v, e := ToStruct(nil, nil)
		if v != nil || e == nil {
			t.Fatalf("%s failed: %v - %e", name, v, e)
		}
	}

	type Abc struct{ Key1 int }
	typeAbc := reflect.TypeOf(Abc{})
	zeroAbc := Abc{}

	type Def struct {
		Abc
		Key2 string
	}
	typeDef := reflect.TypeOf(Def{})
	zeroDef := Def{}

	{
		ZeroMode = true
		v, e := ToStruct(nil, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != zeroAbc {
			t.Fatalf("%s failed: expected %#v but received %#v", name, zeroAbc, v)
		}
		v, e = ToStruct(nil, typeDef)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != zeroDef {
			t.Fatalf("%s failed: expected %#v but received %#v", name, zeroDef, v)
		}

		ZeroMode = false
		_, e = ToStruct(nil, typeAbc)
		if e == nil {
			t.Fatalf("%s failed: [nil] should not be convertible to struct Abc when ZeroMode=false!", name)
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != zeroAbc {
			t.Fatalf("%s failed: expected %#v but received %#v", name, zeroAbc, v)
		}
		v, e = Convert(nil, typeDef)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != zeroDef {
			t.Fatalf("%s failed: expected %#v but received %#v", name, zeroDef, v)
		}

		ZeroMode = false
		_, e = Convert(nil, typeAbc)
		if e == nil {
			t.Fatalf("%s failed: [nil] should not be convertible to struct Abc when ZeroMode=false!", name)
		}
	}

	{
		// Abc is convertible to Abc
		input := Abc{}
		v, e := ToStruct(input, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}
	{
		// Abc is convertible to Abc
		input := Abc{}
		v, e := Convert(input, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}

	{
		// Abc is NOT convertible to Def
		input := Abc{}
		_, e := ToStruct(input, typeDef)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to struct Def!", name, input)
		}
	}
	{
		// Abc is NOT convertible to Def
		input := Abc{}
		_, e := Convert(input, typeDef)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to struct Def!", name, input)
		}
	}

	{
		// Def is convertible to Def
		input := Def{}
		v, e := ToStruct(input, typeDef)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}
	{
		// Def is convertible to Def
		input := Def{}
		v, e := Convert(input, typeDef)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}

	{
		// Def is convertible to Abc
		input := Def{}
		v, e := ToStruct(input, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input.Abc {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}
	{
		// Def is convertible to Abc
		input := Def{}
		v, e := Convert(input, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input.Abc {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}

	{
		input := Abc{}
		_, e := ToStruct(input, TypeString)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to string!", name, input)
		}
	}
	{
		input := "a string"
		_, e := ToStruct(input, typeAbc)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to struct Abc!", name, input)
		}
	}
	{
		input := "another string"
		_, e := Convert(input, typeAbc)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to struct Abc!", name, input)
		}
	}
}

/*----------------------------------------------------------------------*/
func testToSlice(t *testing.T, input interface{}, expected interface{}, typ reflect.Type) {
	name := "TestToSlice"

	v, e := ToSlice(input, typ)
	if input == nil {
		if v != nil || e != nil {
			t.Fatalf("%s failed: %v - %e", name, v, e)
		}
		return
	}
	if typ == nil {
		if v != nil || e == nil {
			t.Fatalf("%s failed: %v - %e", name, v, e)
		}
		return
	}

	if e != nil {
		t.Fatalf("%s failed: %e", name, e)
	} else {
		from := reflect.ValueOf(v)
		to := reflect.ValueOf(expected)
		if from.Len() != to.Len() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
		for i, n := 0, from.Len(); i < n; i++ {
			if from.Index(i).Interface() != to.Index(i).Interface() {
				t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
				break
			}
		}
	}

	if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		v, e = Convert(input, typ)
	} else {
		v, e = Convert(input, reflect.SliceOf(typ))
	}
	if e != nil {
		t.Fatalf("%s failed: %e", name, e)
	} else {
		from := reflect.ValueOf(v)
		to := reflect.ValueOf(expected)
		if from.Len() != to.Len() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
		for i, n := 0, from.Len(); i < n; i++ {
			if from.Index(i).Interface() != to.Index(i).Interface() {
				t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
				break
			}
		}
	}
}

// TestToSlice tests if values are converted correctly to slice
func TestToSlice(t *testing.T) {
	name := "TestToSlice"

	{
		testToSlice(t, nil, nil, reflect.TypeOf([0]int{}))
		testToSlice(t, []bool{true, false}, []int{1, 0}, nil)
	}

	{
		input := "a string"
		testToSlice(t, input, []byte(input), reflect.TypeOf([0]byte{}))
	}

	{
		ZeroMode = true
		input := []interface{}{1, "2", 3.4, false, nil}
		testToSlice(t, input, []int{1, 2, 3, 0, 0}, reflect.TypeOf([]int{}))
		testToSlice(t, input, []float64{1.0, 2.0, 3.4, 0.0, 0.0}, reflect.TypeOf([]float64{}))
		now := time.Now()
		input = []interface{}{1, "2", 3.4, false, now, nil}
		testToSlice(t, input, []interface{}{1, "2", 3.4, false, now, nil}, reflect.TypeOf([]interface{}{}))

		ZeroMode = false
		input = []interface{}{1, "2", 3.4, false, nil}
		_, e := ToSlice(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: input %#v should result error", name, input)
		}
		// nil value can be converted to interface{}(nil)!
		input = []interface{}{1, "2", 3.4, false, now, nil}
		testToSlice(t, input, []interface{}{1, "2", 3.4, false, now, nil}, reflect.TypeOf([]interface{}{}))
	}

	{
		input := []bool{true, false}
		testToSlice(t, input, []int{1, 0}, reflect.TypeOf([0]int{}))
	}
	{
		input := [5]int{-2, 1, 0, 1, 2}
		testToSlice(t, input, []string{"-2", "1", "0", "1", "2"}, reflect.TypeOf([]string{}))
	}
	{
		input := []bool{true, false}
		testToSlice(t, input, []string{"true", "false"}, TypeString)
	}

	{
		input := ""
		_, e := ToSlice(input, reflect.TypeOf([0]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", name, input)
		}
	}
	{
		input := ""
		_, e := Convert(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", name, input)
		}
	}

	{
		input := []string{"a", "b", "c"}
		_, e := ToSlice(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", name, input)
		}
	}
	{
		input := []string{"a", "b", "c"}
		_, e := Convert(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", name, input)
		}
	}
}

/*----------------------------------------------------------------------*/
func testToMap(t *testing.T, input interface{}, expected interface{}, typ reflect.Type) {
	name := "TestToMap"

	v, e := ToMap(input, typ)
	if input == nil {
		if v != nil || e != nil {
			t.Fatalf("%s failed: %v - %e", name, v, e)
		}
		return
	}
	if typ == nil {
		if v != nil || e == nil {
			t.Fatalf("%s failed: %v - %e", name, v, e)
		}
		return
	}

	if e != nil {
		t.Fatalf("%s failed: %e / ZeroMode: %v / Input: %#v", name, e, ZeroMode, input)
	} else {
		if !reflect.DeepEqual(expected, v) {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
		// from := reflect.ValueOf(v)
		// to := reflect.ValueOf(expected)
		// if from.Len() != to.Len() {
		// 	t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		// }
		// for _, k := range from.MapKeys() {
		// 	if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
		// 		t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		// 	}
		// }
		// for _, k := range to.MapKeys() {
		// 	if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
		// 		t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		// 	}
		// }
	}

	v, e = Convert(input, typ)
	if e != nil {
		t.Fatalf("%s failed: %e", name, e)
	} else {
		if !reflect.DeepEqual(expected, v) {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
		// from := reflect.ValueOf(v)
		// to := reflect.ValueOf(expected)
		// if from.Len() != to.Len() {
		// 	t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		// }
		// for _, k := range from.MapKeys() {
		// 	if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
		// 		t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		// 	}
		// }
		// for _, k := range to.MapKeys() {
		// 	if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
		// 		t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		// 	}
		// }
	}
}

// TestToMap tests if values are converted correctly to map
func TestToMap(t *testing.T) {
	name := "TestToMap"

	{
		typeMap := reflect.TypeOf(map[string]interface{}{})

		ZeroMode = true
		testToMap(t, nil, nil, nil)
		testToMap(t, map[string]bool{"1": true, "0": false}, map[int]string{0: "-false-", 1: "-true-"}, nil)
		testToMap(t, nil, nil, typeMap)
		testToMap(t, map[int]bool{1: true, 0: false}, map[string]interface{}{"1": true, "0": false}, typeMap)
		testToMap(t, map[int]interface{}{0: false, 1: "one", 2: 2}, map[string]interface{}{"0": false, "1": "one", "2": 2}, typeMap)

		ZeroMode = false
		testToMap(t, nil, nil, nil)
		testToMap(t, map[string]bool{"1": true, "0": false}, map[int]string{0: "-false-", 1: "-true-"}, nil)
		testToMap(t, nil, nil, typeMap)
		testToMap(t, map[int]bool{1: true, 0: false}, map[string]interface{}{"1": true, "0": false}, typeMap)
		testToMap(t, map[int]interface{}{0: false, 1: "one", 2: 2}, map[string]interface{}{"0": false, "1": "one", "2": 2}, typeMap)
	}
	{
		ZeroMode = true
		typeMap := reflect.TypeOf(map[string]interface{}{})
		now := time.Now()
		testToMap(t,
			map[interface{}]interface{}{0: false, "1": "one", 2: 2, "3": now, nil: "nil key",
				"map": map[string]interface{}{"true": true, "false": "false", "nil": nil}, "mapnil": nil,
				"list": []interface{}{1, "2", true, nil}, "listnil": nil},
			map[string]interface{}{"0": false, "1": "one", "2": 2, "3": now, "": "nil key",
				"map": map[string]interface{}{"true": true, "false": "false", "nil": nil}, "mapnil": nil,
				"list": []interface{}{1, "2", true, nil}, "listnil": nil},
			typeMap)

		ZeroMode = false
		input := map[interface{}]interface{}{0: false, "1": "one", 2: 2, "3": now, nil: "nil key",
			"map": map[string]interface{}{"true": true, "false": "false", "nil": nil}, "mapnil": nil,
			"list": []interface{}{1, "2", true, nil}, "listnil": nil}
		_, err := ToMap(input, typeMap)
		if err == nil {
			t.Fatalf("%s failed: input %#v should result error", name, input)
		}
	}

	{
		input := map[string]bool{"1": true, "0": false}
		testToMap(t, input, map[int]string{0: "false", 1: "true"}, reflect.TypeOf(map[int]string{}))
	}

	{
		input := map[string]bool{"one": true, "0": false}
		_, e := ToMap(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map[int]string!", name, input)
		}
	}

	{
		input := map[bool]string{true: "1", false: "zero"}
		_, e := ToMap(input, reflect.TypeOf(map[bool]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map[bool]int!", name, input)
		}
	}

	{
		input := ""
		_, e := ToMap(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map!", name, input)
		}
	}
	{
		input := ""
		_, e := Convert(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map!", name, input)
		}
	}
	{
		input := map[string]bool{"1": true, "0": false}
		_, e := ToMap(input, TypeString)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to string!", name, input)
		}
	}
}

/*----------------------------------------------------------------------*/

// TestToPointer tests if values are converted correctly to pointer
func TestToPointer(t *testing.T) {
	name := "TestToPointer"

	{
		output, e := ToPointer(nil, nil)
		if output != nil || e != nil {
			t.Fatalf("%s failed: %v - %e", name, output, e)
		}
	}
	{
		input := 1
		output, e := ToPointer(&input, nil)
		if output != nil || e == nil {
			t.Fatalf("%s failed: %v - %e", name, output, e)
		}
	}
	{
		zero := int32(0)
		output, e := ToPointer(nil, reflect.TypeOf(&zero))
		if output != nil || e != nil {
			t.Fatalf("%s failed: %v - %e", name, output, e)
		}
	}

	{
		a := float64(1.23)
		zero := int32(0)
		output, e := ToPointer(&a, reflect.TypeOf(&zero))
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else {
			i32 := *output.(*interface{})
			if i32.(int32) != 1 {
				t.Fatalf("%s failed: received [%#v]", name, output)
			}
		}
	}
	{
		a := float64(1.23)
		zero := int32(0)
		output, e := Convert(&a, reflect.TypeOf(&zero))
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else {
			i32 := *output.(*interface{})
			if i32.(int32) != 1 {
				t.Fatalf("%s failed: received [%#v]", name, output)
			}
		}
	}

	{
		a := string("1.23")
		zero := float64(0)
		output, e := ToPointer(&a, reflect.TypeOf(&zero))
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else {
			f64 := *output.(*interface{})
			if f64.(float64) != 1.23 {
				t.Fatalf("%s failed: received [%#v]", name, output)
			}
		}
	}
	{
		a := string("1.23")
		zero := float64(0)
		output, e := Convert(&a, reflect.TypeOf(&zero))
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else {
			f64 := *output.(*interface{})
			if f64.(float64) != 1.23 {
				t.Fatalf("%s failed: received [%#v]", name, output)
			}
		}
	}

	{
		a := string("blabla")
		zero := float64(0)
		_, e := ToPointer(&a, reflect.TypeOf(&zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to [%#v]!", name, &a, &zero)
		}
	}

	{
		a := ""
		zero := int64(0)
		_, e := ToPointer(a, reflect.TypeOf(&zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to [%#v]!", name, a, &zero)
		}
	}

	{
		a := ""
		zero := int64(0)
		_, e := ToPointer(&a, reflect.TypeOf(&zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to [%#v]!", name, &a, zero)
		}
	}

	{
		type Abc struct {
			A int
		}
		type Def struct {
			Abc
			D string
		}
		a := Def{Abc: Abc{1}, D: "2"}
		output, e := ToPointer(&a, reflect.TypeOf(&Abc{}))
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else {
			abc := *output.(*interface{})
			if abc.(Abc).A != 1 {
				t.Fatalf("%s failed: received [%#v]", name, output)
			}
		}
	}
	{
		type Abc struct {
			A int
		}
		type Def struct {
			Abc
			D string
		}
		a := Def{Abc: Abc{1}, D: "2"}
		output, e := Convert(&a, reflect.TypeOf(&Abc{}))
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else {
			abc := *output.(*interface{})
			if abc.(Abc).A != 1 {
				t.Fatalf("%s failed: received [%#v]", name, output)
			}
		}
	}
}

/*----------------------------------------------------------------------*/

func TestConvert(t *testing.T) {
	name := "TestConvert"

	{
		ZeroMode = true
		v, e := Convert("", nil)
		if e != nil || v != "" {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, "", v)
		}
		ZeroMode = false
		v, e = Convert("", nil)
		if e != nil || v != "" {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, "", v)
		}
	}
	{
		ZeroMode = true
		v, e := Convert(nil, TypeString)
		if v != "" || e != nil {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, "", v)
		}

		ZeroMode = false
		_, e = Convert(nil, TypeString)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to [%#v]!", name, nil, "")
		}
	}
	{
		ZeroMode = true
		input := ""
		zero := func() {}
		_, e := Convert(input, reflect.TypeOf(zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to func!", name, input)
		}
		ZeroMode = false
		_, e = Convert(input, reflect.TypeOf(zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to func!", name, input)
		}
	}
}
