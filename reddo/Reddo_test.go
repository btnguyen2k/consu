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
	var inputList = []interface{}{nil, false, true}
	var expectedList = []bool{false, false, true}
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
			t.Fatalf("TestToBool failed: [%#v] should not be convertable to bool!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeBool)
		if e == nil {
			t.Fatalf("TestToBool failed: [%#v] should not be convertable to bool!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToBool(input)
		if e == nil {
			t.Fatalf("TestToBool failed: [%#v] should not be convertable to bool!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeBool)
		if e == nil {
			t.Fatalf("TestToBool failed: [%#v] should not be convertable to bool!", input)
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

	inputList = []interface{}{nil, "0", "0.0", "0.001", "-0.001", "-1.2", "3.4", "-1E9", "1e9", "-1e-9", "1E-9"}
	expectedList = []float64{0.0, 0.0, 0.0, 0.001, -0.001, -1.2, 3.4, -1e9, 1e9, -1e-9, 1e-9}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i])
	}

	{
		input := "blabla"
		_, e := ToFloat(input)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertable to float!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeFloat)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertable to float!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToFloat(input)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertable to float!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeFloat)
		if e == nil {
			t.Fatalf("TestToFloat failed: [%#v] should not be convertable to float!", input)
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

	inputList = []interface{}{nil, "0", "-1", "2", "-3", "4"}
	expectedList = []int64{0, 0, -1, 2, -3, 4}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i])
	}

	{
		input := "-1.2"
		_, e := ToInt(input)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertable to int!", input)
		}
	}
	{
		input := "3.4"
		_, e := Convert(input, TypeInt)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertable to int!", input)
		}
	}

	{
		input := "blabla"
		_, e := ToInt(input)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertable to int!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeInt)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertable to int!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToInt(input)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertable to int!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeInt)
		if e == nil {
			t.Fatalf("TestToInt failed: [%#v] should not be convertable to int!", input)
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

	inputList = []interface{}{nil, "0", "1", "2", "3", "4"}
	expectedList = []uint64{0, 0, 1, 2, 3, 4}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i])
	}

	{
		input := "-1"
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertable to uint!", input)
		}
	}
	{
		input := "-1.2"
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertable to uint!", input)
		}
	}
	{
		input := "3.4"
		_, e := Convert(input, TypeUint)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertable to uint!", input)
		}
	}

	{
		input := "blabla"
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertable to uint!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, TypeUint)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertable to uint!", input)
		}
	}

	{
		input := struct {
		}{}
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertable to uint!", input)
		}
	}
	{
		input := struct {
		}{}
		_, e := Convert(input, TypeUint)
		if e == nil {
			t.Fatalf("TestToUint failed: [%#v] should not be convertable to uint!", input)
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

	inputList = []interface{}{nil, []byte("a string"), "0", "-1", "2", "-3", "4", "a", "b", "c", ""}
	expectedList = []string{"", "a string", "0", "-1", "2", "-3", "4", "a", "b", "c", ""}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i])
	}

	{
		input := struct {
		}{}
		testToString(t, input, fmt.Sprint(input))
	}
}

/*----------------------------------------------------------------------*/
// TestToTimeError tests if values are converted correctly to time.Time
func TestToTimeError(t *testing.T) {
	{
		input := -1
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertable to time.Time!", input)
		}
	}

	{
		input := "-1"
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertable to time.Time!", input)
		}
	}

	{
		input := "-1.abc"
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertable to time.Time!", input)
		}
	}

	{
		input := struct {
		}{}
		_, err := ToStruct(input, TypeTime)
		if err == nil {
			t.Fatalf("TestToTime failed: [%#v] should not be convertable to time.Time!", input)
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
	v, e := ToTime(nil)
	if e != nil {
		t.Fatalf("%s failed: %e", name, e)
	} else if v.Unix() != zeroTime.Unix() {
		t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, zeroTime, v)
	}
}

// TestToTimeWithLayout tests if strings are converted correctly to time.Time using layout
func TestToTimeWithLayout(t *testing.T) {
	name := "TestToTimeWithLayout"

	{
		// convert nil to 'time.Time'
		v, e := ToTimeWithLayout(nil, "")
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v.Unix() != zeroTime.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, zeroTime, v)
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

	type Def struct {
		Abc
		Key2 string
	}
	typeDef := reflect.TypeOf(Def{})

	{
		// Abc is convertable to Abc
		input := Abc{}
		v, e := ToStruct(input, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}
	{
		// Abc is convertable to Abc
		input := Abc{}
		v, e := Convert(input, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}

	{
		// Abc is NOT convertable to Def
		input := Abc{}
		_, e := ToStruct(input, typeDef)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to struct Def!", name, input)
		}
	}
	{
		// Abc is NOT convertable to Def
		input := Abc{}
		_, e := Convert(input, typeDef)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to struct Def!", name, input)
		}
	}

	{
		// Def is convertable to Def
		input := Def{}
		v, e := ToStruct(input, typeDef)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}
	{
		// Def is convertable to Def
		input := Def{}
		v, e := Convert(input, typeDef)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}

	{
		// Def is convertable to Abc
		input := Def{}
		v, e := ToStruct(input, typeAbc)
		if e != nil {
			t.Fatalf("%s failed: %e", name, e)
		} else if v != input.Abc {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, input, v)
		}
	}
	{
		// Def is convertable to Abc
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
			t.Fatalf("%s failed: [%#v] should not be convertable to string!", name, input)
		}
	}
	{
		input := ""
		_, e := ToStruct(input, typeAbc)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to struct Abc!", name, input)
		}
	}
	{
		input := ""
		_, e := Convert(input, typeAbc)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to struct Abc!", name, input)
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
			t.Fatalf("%s failed: [%#v] should not be convertable to []int!", name, input)
		}
	}
	{
		input := ""
		_, e := Convert(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to []int!", name, input)
		}
	}
	// {
	// 	input := []bool{true, false}
	// 	_, e := ToSlice(input, TypeString)
	// 	if e == nil {
	// 		t.Fatalf("%s failed: [%#v] should not be convertable to string!", name, input)
	// 	}
	// }

	{
		input := []string{"a", "b", "c"}
		_, e := ToSlice(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to []int!", name, input)
		}
	}
	{
		input := []string{"a", "b", "c"}
		_, e := Convert(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to []int!", name, input)
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
		t.Fatalf("%s failed: %e", name, e)
	} else {
		from := reflect.ValueOf(v)
		to := reflect.ValueOf(expected)
		if from.Len() != to.Len() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
		for _, k := range from.MapKeys() {
			if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
				t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
			}
		}
		for _, k := range to.MapKeys() {
			if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
				t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
			}
		}
	}

	v, e = Convert(input, typ)
	if e != nil {
		t.Fatalf("%s failed: %e", name, e)
	} else {
		from := reflect.ValueOf(v)
		to := reflect.ValueOf(expected)
		if from.Len() != to.Len() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
		for _, k := range from.MapKeys() {
			if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
				t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
			}
		}
		for _, k := range to.MapKeys() {
			if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
				t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
			}
		}
	}
}

// TestToMap tests if values are converted correctly to map
func TestToMap(t *testing.T) {
	name := "TestToMap"

	{
		testToMap(t, nil, nil, reflect.TypeOf(map[string]interface{}{}))
		testToMap(t, map[string]bool{"1": true, "0": false}, map[int]string{0: "false", 1: "true"}, nil)
	}

	{
		input := map[string]bool{"1": true, "0": false}
		testToMap(t, input, map[int]string{0: "false", 1: "true"}, reflect.TypeOf(map[int]string{}))
	}

	{
		input := map[string]bool{"one": true, "0": false}
		_, e := ToMap(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to map[int]string!", name, input)
		}
	}

	{
		input := map[bool]string{true: "1", false: "zero"}
		_, e := ToMap(input, reflect.TypeOf(map[bool]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to map[bool]int!", name, input)
		}
	}

	{
		input := ""
		_, e := ToMap(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to map!", name, input)
		}
	}
	{
		input := ""
		_, e := Convert(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to map!", name, input)
		}
	}
	{
		input := map[string]bool{"1": true, "0": false}
		_, e := ToMap(input, TypeString)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to string!", name, input)
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
			t.Fatalf("%s failed: [%#v] should not be convertable to [%#v]!", name, &a, &zero)
		}
	}

	{
		a := ""
		zero := int64(0)
		_, e := ToPointer(a, reflect.TypeOf(&zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to [%#v]!", name, a, &zero)
		}
	}

	{
		a := ""
		zero := int64(0)
		_, e := ToPointer(&a, reflect.TypeOf(&zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to [%#v]!", name, &a, zero)
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
		v, e := Convert("", nil)
		if e != nil || v != "" {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", name, "", v)
		}
	}
	{
		_, e := Convert(nil, TypeString)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to [%#v]!", name, nil, "")
		}
	}
	{
		input := ""
		zero := func() {}
		_, e := Convert(input, reflect.TypeOf(zero))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertable to func!", name, input)
		}
	}
}
