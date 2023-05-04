package reddo

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func testToBool(t *testing.T, input interface{}, expected bool) {
	testName := "TestToBool"
	{
		v, e := ToBool(input)
		if e != nil {
			t.Fatalf("%s failed for input <%#v>: %e", testName, input, e)
		} else if v != expected {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, input, expected, v)
		}
	}
	if input != nil {
		v, e := Convert(input, TypeBool)
		if e != nil {
			t.Fatalf("%s failed for input <%#v>: %e", testName, input, e)
		} else if v.(bool) != expected {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, input, expected, v)
		}
	}
}

func testToBoolError(t *testing.T, input interface{}) {
	testName := "TestToBoolError"
	{
		_, e := ToBool(input)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to bool!", testName, input)
		}
	}
	if input != nil {
		_, e := Convert(input, TypeBool)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to bool!", testName, input)
		}
	}
}

// TestToBool tests if values are converted correctly to bool
func TestToBool(t *testing.T) {
	var i = 0
	var p1 *int
	var p2 = &i
	testCases := []struct {
		name     string
		input    []interface{}
		expected []bool
	}{
		{
			"bool",
			[]interface{}{false, true},
			[]bool{false, true},
		},
		{
			"int",
			[]interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)},
			[]bool{false, true, true, false, true, true, false, true, true, false, true, true, false, true, true},
		},
		{
			"uint",
			[]interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)},
			[]bool{false, true, false, true, false, true, false, true, false, true, false, true},
		},
		{
			"float",
			[]interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)},
			[]bool{false, true, true, true, true, false, true, true, true, true},
		},
		{
			"complex",
			[]interface{}{0 + 0i, 0 - 0i, 0 + 2i, 0 - 3i, -1 + 0i, 2 + 0i, 3 - 2i, 3 + 3i, -4 + 5i, -5 + 6i},
			[]bool{false, false, true, true, true, true, true, true, true, true},
		},
		{
			"string",
			[]interface{}{"false", "true", "False", "True", "FALSE", "TRUE"},
			[]bool{false, true, false, true, false, true},
		},
		{
			"pointer",
			[]interface{}{p1, p2},
			[]bool{false, true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, n := 0, len(tc.input); i < n; i++ {
				testToBool(t, tc.input[i], tc.expected[i])
			}
		})
	}

	testToBoolError(t, "blabla")
	testToBoolError(t, struct {
	}{})

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

func testToFloatError(t *testing.T, input interface{}) {
	testName := "TestToFloatError"
	{
		_, e := ToFloat(input)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to float!", testName, input)
		}
	}
	if input != nil {
		_, e := Convert(input, TypeFloat)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to float!", testName, input)
		}
	}
}

// TestToFloat tests if values are converted correctly to float
func TestToFloat(t *testing.T) {
	testCases := []struct {
		name     string
		input    []interface{}
		expected []float64
	}{
		{
			"bool",
			[]interface{}{false, true},
			[]float64{0.0, 1.0},
		},
		{
			"int",
			[]interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)},
			[]float64{0.0, -1.0, 2.0, 0.0, -2.0, 3.0, 0.0, -3.0, 4.0, 0.0, -4.0, 5.0, 0.0, -5.0, 6.0},
		},
		{
			"uint",
			[]interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)},
			[]float64{0.0, 1.0, 0.0, 2.0, 0.0, 3.0, 0.0, 4.0, 0.0, 5.0, 0.0, 6.0},
		},
		{
			"float",
			[]interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)},
			[]float64{0.0, 0.001, -0.001, -1.2, 3.4, 0.0, 0.001, -0.001, -5.6, 7.8},
		},
		{
			"string",
			[]interface{}{"0", "0.0", "0.001", "-0.001", "-1.2", "3.4", "-1E9", "1e9", "-1e-9", "1E-9"},
			[]float64{0.0, 0.0, 0.001, -0.001, -1.2, 3.4, -1e9, 1e9, -1e-9, 1e-9},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, n := 0, len(tc.input); i < n; i++ {
				testToFloat(t, tc.input[i], tc.expected[i])
			}
		})
	}

	testToFloatError(t, "blabla")
	testToFloatError(t, struct {
	}{})

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

func testToIntError(t *testing.T, input interface{}) {
	testName := "TestToIntError"
	{
		_, e := ToInt(input)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to int!", testName, input)
		}
	}
	if input != nil {
		_, e := Convert(input, TypeInt)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to int!", testName, input)
		}
	}
}

// TestToInt tests if values are converted correctly to int
func TestToInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    []interface{}
		expected []int64
	}{
		{
			"bool",
			[]interface{}{false, true},
			[]int64{0, 1},
		},
		{
			"int",
			[]interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)},
			[]int64{0, -1, 2, 0, -2, 3, 0, -3, 4, 0, -4, 5, 0, -5, 6},
		},
		{
			"uint",
			[]interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)},
			[]int64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6},
		},
		{
			"float",
			[]interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)},
			[]int64{0, 0, -0, -1, 3, 0, 0, -0, -5, 7},
		},
		{
			"string",
			[]interface{}{"0", "-1", "2", "-3", "4"},
			[]int64{0, -1, 2, -3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, n := 0, len(tc.input); i < n; i++ {
				testToInt(t, tc.input[i], tc.expected[i])
			}
		})
	}

	testToIntError(t, "blabla")
	testToIntError(t, struct {
	}{})
	testToIntError(t, "-1.2")

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

func testToUintError(t *testing.T, input interface{}) {
	testName := "TestToUintError"
	{
		_, e := ToUint(input)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to uint!", testName, input)
		}
	}
	if input != nil {
		_, e := Convert(input, TypeUint)
		if e == nil {
			t.Fatalf("[%s] failed: [%#v] should not be convertible to uint!", testName, input)
		}
	}
}

// TestToUint tests if values are converted correctly to uint
func TestToUint(t *testing.T) {
	testCases := []struct {
		name     string
		input    []interface{}
		expected []uint64
	}{
		{
			"bool",
			[]interface{}{false, true},
			[]uint64{0, 1},
		},
		{
			"int",
			[]interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)},
			[]uint64{0, MaxUint, 2, 0, MaxUint - 1, 3, 0, MaxUint - 2, 4, 0, MaxUint - 3, 5, 0, MaxUint - 4, 6},
		},
		{
			"uint",
			[]interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)},
			[]uint64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6},
		},
		{
			"float",
			[]interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)},
			[]uint64{0, 0, 0, MaxUint, 3, 0, 0, 0, MaxUint - 4, 7},
		},
		{
			"string",
			[]interface{}{"0", "1", "2", "3", "4"},
			[]uint64{0, 1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, n := 0, len(tc.input); i < n; i++ {
				testToUint(t, tc.input[i], tc.expected[i])
			}
		})
	}

	testToUintError(t, "blabla")
	testToUintError(t, struct {
	}{})
	testToUintError(t, "-2.3")
	testToUintError(t, "-1")

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
	testCases := []struct {
		name     string
		input    []interface{}
		expected []string
	}{
		{
			"bool",
			[]interface{}{false, true},
			[]string{"false", "true"},
		},
		{
			"int",
			[]interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)},
			[]string{"0", "-1", "2", "0", "-2", "3", "0", "-3", "4", "0", "-4", "5", "0", "-5", "6"},
		},
		{
			"uint",
			[]interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)},
			[]string{"0", "1", "0", "2", "0", "3", "0", "4", "0", "5", "0", "6"},
		},
		{
			"float",
			[]interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)},
			[]string{"0", "0.001", "-0.001", "-1.2", "3.4", "0", "0.001", "-0.001", "-5.6", "7.8"},
		},
		{
			"string",
			[]interface{}{[]byte("a string"), "0", "-1", "2", "-3", "4", "a", "b", "c", ""},
			[]string{"a string", "0", "-1", "2", "-3", "4", "a", "b", "c", ""},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, n := 0, len(tc.input); i < n; i++ {
				expected := tc.expected[i]
				v := reflect.ValueOf(tc.input[i])
				if v.Kind() == reflect.Float32 {
					expected = strconv.FormatFloat(v.Float(), 'g', -1, 64)
				}
				testToString(t, tc.input[i], expected)
			}
		})
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
// TestToTimeError tests if error should be raised when invalid values are being convertedto time.Time
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

func TestToTime(t *testing.T) {
	testName := "TestToTime"
	now := time.Now()
	testCases := []struct {
		name       string
		input      interface{}
		expected   time.Time
		resolution string
	}{
		{"struct", now, now, "nano"},
		{"int(seconds)", now.Unix(), now, "second"},
		{"int(milliseconds)", now.UnixMilli(), now, "milli"},
		{"int(microseconds)", now.UnixMicro(), now, "micro"},
		{"int(nanoseconds)", now.UnixNano(), now, "nano"},
		{"string(seconds)", strconv.FormatInt(now.Unix(), 10), now, "second"},
		{"string(milliseconds)", strconv.FormatInt(now.UnixMilli(), 10), now, "milli"},
		{"string(microseconds)", strconv.FormatInt(now.UnixMicro(), 10), now, "micro"},
		{"string(nanoseconds)", strconv.FormatInt(now.UnixNano(), 10), now, "nano"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, e := ToStruct(tc.input, TypeTime)
			if e != nil {
				t.Fatalf("%s failed for input <%#v>: %e", testName, tc.input, e)
			} else if tc.resolution == "second" && v.(time.Time).Unix() != tc.expected.Unix() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			} else if tc.resolution == "milli" && v.(time.Time).UnixMilli() != tc.expected.UnixMilli() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			} else if tc.resolution == "micro" && v.(time.Time).UnixMicro() != tc.expected.UnixMicro() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			} else if tc.resolution == "nano" && v.(time.Time).UnixNano() != tc.expected.UnixNano() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			}
		})
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
	testName := "TestToTimeWithLayout"
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	testCases := []struct {
		name, layout string
		input        interface{}
		expected     time.Time
		resolution   string
	}{
		{"struct", "", now, now, "nano"},
		{"int(seconds)", "", now.Unix(), now, "second"},
		{"int(milliseconds)", "", now.UnixMilli(), now, "milli"},
		{"int(microseconds)", "", now.UnixMicro(), now, "micro"},
		{"int(nanoseconds)", "", now.UnixNano(), now, "nano"},
		{"string(seconds)", "", strconv.FormatInt(now.Unix(), 10), now, "second"},
		{"string(milliseconds)", "", strconv.FormatInt(now.UnixMilli(), 10), now, "milli"},
		{"string(microseconds)", "", strconv.FormatInt(now.UnixMicro(), 10), now, "micro"},
		{"string(nanoseconds)", "", strconv.FormatInt(now.UnixNano(), 10), now, "nano"},
		{"UTC", "2006-01-02T15:04:05", "2019-04-29T20:59:10", time.Date(2019, 04, 29, 20, 59, 10, 0, time.UTC), "nano"},
		{"Asia/Ho_Chi_Minh", "2006-01-02 15:04:05.999999 -0700 -07", "2019-04-29 20:59:10.165067 +0700 +07", time.Date(2019, 04, 29, 20, 59, 10, 0, loc), "second"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, e := ToTimeWithLayout(tc.input, tc.layout)
			if e != nil {
				t.Fatalf("%s failed for input <%#v>: %e", testName, tc.input, e)
			} else if tc.resolution == "second" && v.Unix() != tc.expected.Unix() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			} else if tc.resolution == "milli" && v.UnixMilli() != tc.expected.UnixMilli() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			} else if tc.resolution == "micro" && v.UnixMicro() != tc.expected.UnixMicro() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			} else if tc.resolution == "nano" && v.UnixNano() != tc.expected.UnixNano() {
				t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", testName, tc.input, tc.expected, v)
			}
		})
	}

	{
		ZeroMode = true
		v, e := ToTimeWithLayout(nil, "")
		if e != nil {
			t.Fatalf("%s failed: %e", testName, e)
		} else if v.Unix() != zeroTime.Unix() {
			t.Fatalf("%s failed: expected [%#v] but received [%#v]", testName, zeroTime, v)
		}

		ZeroMode = false
		_, e = ToTimeWithLayout(nil, "")
		if e == nil {
			t.Fatalf("%s failed: [nil] should not be convertible to time.Time when ZeroMode=false!", testName)
		}
	}

	{
		// invalid input
		input := "abc"
		layout := "Jan"
		_, e := ToTimeWithLayout(input, layout)
		if e == nil {
			t.Fatalf("%s failed: %e", testName, e)
		}
	}
	{
		// invalid layout
		input := "2019-01-01"
		layout := "month"
		_, e := ToTimeWithLayout(input, layout)
		if e == nil {
			t.Fatalf("%s failed: %e", testName, e)
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
			t.Fatalf("%s failed for input <nil>: %v - %e", name, v, e)
		}
		return
	}
	if typ == nil {
		if v != nil || e == nil {
			t.Fatalf("%s failed for type <nil>: %v - %e", name, v, e)
		}
		return
	}

	if e != nil {
		t.Fatalf("%s failed for input <%#v>: %e", name, input, e)
	} else {
		from := reflect.ValueOf(v)
		to := reflect.ValueOf(expected)
		if from.Len() != to.Len() {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", name, input, expected, v)
		}
		if !reflect.DeepEqual(from.Interface(), to.Interface()) {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", name, input, expected, v)
		}
	}

	if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		v, e = Convert(input, typ)
	} else {
		v, e = Convert(input, reflect.SliceOf(typ))
	}
	if e != nil {
		t.Fatalf("%s failed for input <%#v>: %e", name, input, e)
	} else {
		from := reflect.ValueOf(v)
		to := reflect.ValueOf(expected)
		if from.Len() != to.Len() {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", name, input, expected, v)
		}
		if !reflect.DeepEqual(from.Interface(), to.Interface()) {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", name, input, expected, v)
		}
	}
}

// TestToSlice tests if values are converted correctly to slice
func TestToSlice(t *testing.T) {
	testName := "TestToSlice"
	testCases := []struct {
		name            string
		input, expected interface{}
		typ             reflect.Type
		nonZeroMode     bool
	}{
		{"nil", nil, nil, reflect.TypeOf([0]int{}), true},
		{"[]bool to []int/type=nil", []bool{true, false}, []int{1, 0}, nil, true},
		{"[]bool to []float/type=[0]float64", []bool{false, true}, []float64{0.0, 1.0}, reflect.TypeOf([0]float64{}), true},
		{"string to []byte/type=[0]byte", "a string", []byte("a string"), reflect.TypeOf([0]byte{}), true},
		{"[5]int to []string/type=[]string", [5]int{-2, 1, 0, 1, 2}, []string{"-2", "1", "0", "1", "2"}, reflect.TypeOf([]string{}), true},
		{"[]bool to []string/type=string", []bool{true, false}, []string{"true", "false"}, TypeString, true},
		{"[]interface to []string/type=string", []interface{}{"str", 1, false, nil}, []string{"str", "1", "false", ""}, TypeString, false},
		{"[]interface to []int/type=[]int", []interface{}{"1", -2, true, nil}, []int{1, -2, 1, 0}, reflect.TypeOf([]int{}), false},
		{"[]interface to []float/type=[]float", []interface{}{"1.2", -3.4, false, nil}, []float32{1.2, -3.4, 0.0, 0.0}, reflect.TypeOf([]float32{}), false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ZeroMode = true
			testToSlice(t, tc.input, tc.expected, tc.typ)
			if tc.nonZeroMode {
				ZeroMode = false
				testToSlice(t, tc.input, tc.expected, tc.typ)
			}
		})
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
			t.Fatalf("%s failed: input %#v should result error", testName, input)
		}
		// nil value can be converted to interface{}(nil)!
		input = []interface{}{1, "2", 3.4, false, now, nil}
		testToSlice(t, input, []interface{}{1, "2", 3.4, false, now, nil}, reflect.TypeOf([]interface{}{}))
	}

	{
		input := ""
		_, e := ToSlice(input, reflect.TypeOf([0]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", testName, input)
		}
	}
	{
		input := ""
		_, e := Convert(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", testName, input)
		}
	}

	{
		input := []string{"a", "b", "c"}
		_, e := ToSlice(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", testName, input)
		}
	}
	{
		input := []string{"a", "b", "c"}
		_, e := Convert(input, reflect.TypeOf([]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to []int!", testName, input)
		}
	}
}

// TestToSliceNested tests ToSlice with edge/complex cases
func TestToSliceNested(t *testing.T) {
	testCases := []struct {
		name            string
		input, expected interface{}
		typ             reflect.Type
	}{
		{"slice_of_slice", []interface{}{[]interface{}{1, true, "str"}, []int{1, 2, 3}, nil}, [][]string{[]string{"1", "true", "str"}, []string{"1", "2", "3"}, nil}, reflect.TypeOf([][]string{})},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ZeroMode = true
			testToSlice(t, tc.input, tc.expected, tc.typ)
			ZeroMode = false
			testToSlice(t, tc.input, tc.expected, tc.typ)
		})
	}
}

/*----------------------------------------------------------------------*/
func testToMap(t *testing.T, input interface{}, expected interface{}, typ reflect.Type) {
	name := "TestToMap"

	v, err := ToMap(input, typ)
	if input == nil {
		// if input is nil, returned value must be nil and no error
		if v != nil || err != nil {
			t.Fatalf("%s failed for input <nil>: %v - %err", name, v, err)
		}
		return
	}
	if typ == nil {
		// if type is nil, returned value must be nil and there is error
		if v != nil || err == nil {
			t.Fatalf("%s failed for type <nil>: %v - %err", name, v, err)
		}
		return
	}

	if err != nil {
		t.Fatalf("%s failed: %e / ZeroMode: %v / Input: %#v", name, err, ZeroMode, input)
	} else {
		if !reflect.DeepEqual(expected, v) {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", name, input, expected, v)
		}
	}

	v, err = Convert(input, typ)
	if err != nil {
		t.Fatalf("%s failed: %e / ZeroMode: %v / Input: %#v", name, err, ZeroMode, input)
	} else {
		if !reflect.DeepEqual(expected, v) {
			t.Fatalf("%s failed for input <%#v>: expected [%#v] but received [%#v]", name, input, expected, v)
		}
	}
}

// TestToMap tests if values are converted correctly to map
func TestToMap(t *testing.T) {
	testName := "TestToMap"

	testCases := []struct {
		name        string
		input       interface{}
		expected    interface{}
		typ         reflect.Type
		nonZeroMode bool
	}{
		{"all_nil", nil, nil, nil, true},
		{"nil_input", nil, nil, reflect.TypeOf(map[string]interface{}{}), true},
		{"nil_type", map[string]bool{"1": true, "0": false}, nil, nil, true},
		{"map[int]bool -> map[string]interface", map[int]bool{1: true, 0: false}, map[string]interface{}{"1": true, "0": false}, reflect.TypeOf(map[string]interface{}{}), true},
		{"map[int]interface -> map[string]interface", map[int]interface{}{0: false, 1: "one", 2: 2}, map[string]interface{}{"0": false, "1": "one", "2": 2}, reflect.TypeOf(map[string]interface{}{}), true},
		{"map[string]bool -> map[int]string", map[string]bool{"1": true, "0": false}, map[int]string{0: "false", 1: "true"}, reflect.TypeOf(map[int]string{}), true},
		{"nil_value", map[string]interface{}{"k1": "value 1", "k2": 1, "k3": true, "k4": nil}, map[string]string{"k1": "value 1", "k2": "1", "k3": "true", "k4": ""}, reflect.TypeOf(map[string]string{}), false},
		{"nil_key", map[interface{}]string{"value 1": "k1", 1.2: "k2", false: "k3", nil: "k4"}, map[string]string{"value 1": "k1", "1.2": "k2", "false": "k3", "": "k4"}, reflect.TypeOf(map[string]string{}), false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ZeroMode = true
			testToMap(t, tc.input, tc.expected, tc.typ)
			if tc.nonZeroMode {
				ZeroMode = false
				testToMap(t, tc.input, tc.expected, tc.typ)
			}
		})
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
			t.Fatalf("%s failed: input %#v should result error", testName, input)
		}
	}

	{
		input := map[string]bool{"one": true, "0": false}
		_, e := ToMap(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map[int]string!", testName, input)
		}
	}
	{
		input := map[bool]string{true: "1", false: "zero"}
		_, e := ToMap(input, reflect.TypeOf(map[bool]int{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map[bool]int!", testName, input)
		}
	}
	{
		input := ""
		_, e := ToMap(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map!", testName, input)
		}
	}
	{
		input := ""
		_, e := Convert(input, reflect.TypeOf(map[int]string{}))
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to map!", testName, input)
		}
	}
	{
		input := map[string]bool{"1": true, "0": false}
		_, e := ToMap(input, TypeString)
		if e == nil {
			t.Fatalf("%s failed: [%#v] should not be convertible to string!", testName, input)
		}
	}
}

// TestToMapComplex tests ToMap with edge cases.
func TestToMapComplex(t *testing.T) {
	testName := "TestToMapComplex"
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
			t.Fatalf("%s failed: input %#v should result error", testName, input)
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
