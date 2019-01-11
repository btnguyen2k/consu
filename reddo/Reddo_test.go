package reddo

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func ifFailed(t *testing.T, f string, e error) {
	if e != nil {
		t.Errorf("%s failed: %e", f, e)
	}
}

/*----------------------------------------------------------------------*/

func testToBool(t *testing.T, input interface{}, expected bool, zero bool) {
	v, e := ToBool(input)
	ifFailed(t, "TestToBool", e)
	if v != expected {
		t.Errorf("TestToBool failed: expected [%v] but received [%v]", expected, v)
	}

	v, e = Convert(input, zero)
	ifFailed(t, "TestToBool", e)
	if v != expected {
		t.Errorf("TestToBool failed: expected [%v] but received [%v]", expected, v)
	}
}

// TestToBool tests if values are converted correctly to bool
func TestToBool(t *testing.T) {
	var zero = false
	var inputList = []interface{}{false, true}
	var expectedList = []bool{false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []bool{false, true, true, false, true, true, false, true, true, false, true, true, false, true, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []bool{false, true, false, true, false, true, false, true, false, true, false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []bool{false, true, true, true, true, false, true, true, true, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{0 + 0i, 0 - 0i, 0 + 2i, 0 - 3i, -1 + 0i, 2 + 0i, 3 - 2i, 3 + 3i, -4 + 5i, -5 + 6i}
	expectedList = []bool{false, false, true, true, true, true, true, true, true, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i], zero)
	}

	var i = 0
	var p1 *int
	var p2 = &i
	inputList = []interface{}{p1, p2}
	expectedList = []bool{false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{"false", "true", "False", "True", "FALSE", "TRUE"}
	expectedList = []bool{false, true, false, true, false, true}
	for i, n := 0, len(inputList); i < n; i++ {
		testToBool(t, inputList[i], expectedList[i], zero)
	}

	{
		input := "blabla"
		_, e := ToBool(input)
		if e == nil {
			t.Errorf("TestToBool failed: [%v] should not be convertable to bool!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, zero)
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
	{
		input := struct {
		}{}
		_, e := Convert(input, zero)
		if e == nil {
			t.Errorf("TestToBool failed: [%v] should not be convertable to bool!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

func testToFloat(t *testing.T, input interface{}, expected float64, zero float64) {
	v, e := ToFloat(input)
	ifFailed(t, "TestToFloat", e)
	if v != expected {
		t.Errorf("TestToFloat failed: expected [%f] but received [%f]", expected, v)
	}

	v, e = Convert(input, zero)
	ifFailed(t, "TestToFloat", e)
	if v != expected {
		t.Errorf("TestToFloat failed: expected [%f] but received [%f]", expected, v)
	}
}

// TestToFloat tests if values are converted correctly to float
func TestToFloat(t *testing.T) {
	var zero = float64(0.0)
	var inputList = []interface{}{false, true}
	var expectedList = []float64{0.0, 1.0}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []float64{0.0, -1.0, 2.0, 0.0, -2.0, 3.0, 0.0, -3.0, 4.0, 0.0, -4.0, 5.0, 0.0, -5.0, 6.0}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []float64{0.0, 1.0, 0.0, 2.0, 0.0, 3.0, 0.0, 4.0, 0.0, 5.0, 0.0, 6.0}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []float64{0.0, 0.001, -0.001, -1.2, 3.4, 0.0, 0.001, -0.001, -5.6, 7.8}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{"0", "0.0", "0.001", "-0.001", "-1.2", "3.4", "-1E9", "1e9", "-1e-9", "1E-9"}
	expectedList = []float64{0.0, 0.0, 0.001, -0.001, -1.2, 3.4, -1e9, 1E9, -1E-9, 1e-9}
	for i, n := 0, len(inputList); i < n; i++ {
		testToFloat(t, inputList[i], expectedList[i], zero)
	}

	{
		input := "blabla"
		_, e := ToFloat(input)
		if e == nil {
			t.Errorf("TestToFloat failed: [%v] should not be convertable to float!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, zero)
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
	{
		input := struct {
		}{}
		_, e := Convert(input, zero)
		if e == nil {
			t.Errorf("TestToFloat failed: [%v] should not be convertable to float!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

func testToInt(t *testing.T, input interface{}, expected int64, zero int64) {
	v, e := ToInt(input)
	ifFailed(t, "TestToInt", e)
	if v != expected {
		t.Errorf("TestToInt failed: expected [%d] but received [%d]", expected, v)
	}

	v, e = Convert(input, zero)
	ifFailed(t, "TestToInt", e)
	if v != expected {
		t.Errorf("TestToInt failed: expected [%d] but received [%d]", expected, v)
	}
}

// TestToInt tests if values are converted correctly to int
func TestToInt(t *testing.T) {
	var zero = int64(0)
	var inputList = []interface{}{false, true}
	var expectedList = []int64{0, 1}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []int64{0, -1, 2, 0, -2, 3, 0, -3, 4, 0, -4, 5, 0, -5, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []int64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []int64{0, 0, -0, -1, 3, 0, 0, -0, -5, 7}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{"0", "-1", "2", "-3", "4"}
	expectedList = []int64{0, -1, 2, -3, 4}
	for i, n := 0, len(inputList); i < n; i++ {
		testToInt(t, inputList[i], expectedList[i], zero)
	}

	{
		input := "-1.2"
		_, e := ToInt(input)
		if e == nil {
			t.Errorf("TestToInt failed: [%v] should not be convertable to int!", input)
		}
	}
	{
		input := "3.4"
		_, e := Convert(input, zero)
		if e == nil {
			t.Errorf("TestToInt failed: [%v] should not be convertable to int!", input)
		}
	}

	{
		input := "blabla"
		_, e := ToInt(input)
		if e == nil {
			t.Errorf("TestToInt failed: [%v] should not be convertable to int!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, zero)
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
	{
		input := struct {
		}{}
		_, e := Convert(input, zero)
		if e == nil {
			t.Errorf("TestToInt failed: [%v] should not be convertable to int!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

const (
	MaxUint = ^uint64(0)
)

func testToUint(t *testing.T, input interface{}, expected uint64, zero uint64) {
	v, e := ToUint(input)
	ifFailed(t, "TestToUint", e)
	if v != expected {
		t.Errorf("TestToUint failed: expected [%d] but received [%d]", expected, v)
	}

	v, e = Convert(input, zero)
	ifFailed(t, "TestToUint", e)
	if v != expected {
		t.Errorf("TestToUint failed: expected [%d] but received [%d]", expected, v)
	}
}

// TestToUint tests if values are converted correctly to uint
func TestToUint(t *testing.T) {
	var zero = uint64(0)
	var inputList = []interface{}{false, true}
	var expectedList = []uint64{0, 1}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []uint64{0, MaxUint, 2, 0, MaxUint - 1, 3, 0, MaxUint - 2, 4, 0, MaxUint - 3, 5, 0, MaxUint - 4, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []uint64{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{float32(0.0), float32(0.001), float32(-0.001), float32(-1.2), float32(3.4), float64(0.0), float64(0.001), float64(-0.001), float64(-5.6), float64(7.8)}
	expectedList = []uint64{0, 0, 0, MaxUint, 3, 0, 0, 0, MaxUint - 4, 7}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{"0", "1", "2", "3", "4"}
	expectedList = []uint64{0, 1, 2, 3, 4}
	for i, n := 0, len(inputList); i < n; i++ {
		testToUint(t, inputList[i], expectedList[i], zero)
	}

	{
		input := "-1"
		_, e := ToUint(input)
		if e == nil {
			t.Errorf("TestToUint failed: [%v] should not be convertable to uint!", input)
		}
	}
	{
		input := "-1.2"
		_, e := ToUint(input)
		if e == nil {
			t.Errorf("TestToUint failed: [%v] should not be convertable to uint!", input)
		}
	}
	{
		input := "3.4"
		_, e := Convert(input, zero)
		if e == nil {
			t.Errorf("TestToUint failed: [%v] should not be convertable to uint!", input)
		}
	}

	{
		input := "blabla"
		_, e := ToUint(input)
		if e == nil {
			t.Errorf("TestToUint failed: [%v] should not be convertable to uint!", input)
		}
	}
	{
		input := "blabla"
		_, e := Convert(input, zero)
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
	{
		input := struct {
		}{}
		_, e := Convert(input, zero)
		if e == nil {
			t.Errorf("TestToUint failed: [%v] should not be convertable to uint!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

func testToString(t *testing.T, input interface{}, expected string, zero string) {
	v, e := ToString(input)
	ifFailed(t, "TestToString", e)
	if v != expected {
		t.Errorf("TestToString failed: expected [%s] but received [%s]", expected, v)
	}

	v, e = Convert(input, zero)
	ifFailed(t, "TestToString", e)
	if v != expected {
		t.Errorf("TestToString failed: expected [%s] but received [%s]", expected, v)
	}
}

// TestToString tests if values are converted correctly to string
func TestToString(t *testing.T) {
	var zero = ""
	var inputList = []interface{}{false, true}
	var expectedList = []string{"false", "true"}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{int(0), int(-1), int(2), int8(0), int8(-2), int8(3), int16(0), int16(-3), int16(4), int32(0), int32(-4), int32(5), int64(0), int64(-5), int64(6)}
	expectedList = []string{"0", "-1", "2", "0", "-2", "3", "0", "-3", "4", "0", "-4", "5", "0", "-5", "6"}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i], zero)
	}

	inputList = []interface{}{uint(0), uint(1), uint8(0), uint8(2), uint16(0), uint16(3), uint32(0), uint32(4), uint64(0), uint64(5), uintptr(0), uintptr(6)}
	expectedList = []string{"0", "1", "0", "2", "0", "3", "0", "4", "0", "5", "0", "6"}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i], zero)
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
		testToString(t, inputList[i], expected, zero)
	}

	inputList = []interface{}{"0", "-1", "2", "-3", "4", "a", "b", "c", ""}
	expectedList = []string{"0", "-1", "2", "-3", "4", "a", "b", "c", ""}
	for i, n := 0, len(inputList); i < n; i++ {
		testToString(t, inputList[i], expectedList[i], zero)
	}

	{
		input := struct {
		}{}
		testToString(t, input, fmt.Sprint(input), zero)
	}
}

/*----------------------------------------------------------------------*/

// TestToStruct tests if values are converted correctly to struct
func TestToStruct(t *testing.T) {
	type Abc struct{ Key1 int }
	zeroAbc := Abc{}

	type Def struct {
		Abc
		Key2 string
	}
	zeroDef := Def{}

	{
		// Abc is convertable to Abc
		input := Abc{}
		v, err := ToStruct(input, zeroAbc)
		ifFailed(t, "TestToStruct", err)
		if v != input {
			t.Errorf("TestToStruct failed: expected [%v] but received [%v]", input, v)
		}
	}
	{
		// Abc is convertable to Abc
		input := Abc{}
		v, err := Convert(input, zeroAbc)
		ifFailed(t, "TestToStruct", err)
		if v != input {
			t.Errorf("TestToStruct failed: expected [%v] but received [%v]", input, v)
		}
	}

	{
		// Abc is NOT convertable to Def
		input := Abc{}
		_, err := ToStruct(input, zeroDef)
		if err == nil {
			t.Errorf("TestToStruct failed: [%v] should not be convertable to struct Def!", input)
		}
	}
	{
		// Abc is NOT convertable to Def
		input := Abc{}
		_, err := Convert(input, zeroDef)
		if err == nil {
			t.Errorf("TestToStruct failed: [%v] should not be convertable to struct Def!", input)
		}
	}

	{
		// Def is convertable to Def
		input := Def{}
		v, err := ToStruct(input, zeroDef)
		ifFailed(t, "TestToStruct", err)
		if v != input {
			t.Errorf("TestToStruct failed: expected [%v] but received [%v]", input, v)
		}
	}
	{
		// Def is convertable to Def
		input := Def{}
		v, err := Convert(input, zeroDef)
		ifFailed(t, "TestToStruct", err)
		if v != input {
			t.Errorf("TestToStruct failed: expected [%v] but received [%v]", input, v)
		}
	}

	{
		// Def is convertable to Abc
		input := Def{}
		v, err := ToStruct(input, zeroAbc)
		ifFailed(t, "TestToStruct", err)
		if v != input.Abc {
			t.Errorf("TestToStruct failed: expected [%v] but received [%v]", input.Abc, v)
		}
	}
	{
		// Def is convertable to Abc
		input := Def{}
		v, err := Convert(input, zeroAbc)
		ifFailed(t, "TestToStruct", err)
		if v != input.Abc {
			t.Errorf("TestToStruct failed: expected [%v] but received [%v]", input.Abc, v)
		}
	}

	{
		input := Abc{}
		_, err := ToStruct(input, "")
		if err == nil {
			t.Errorf("TestToStruct failed: [%v] should not be convertable to string!", input)
		}
	}
	{
		input := ""
		_, err := ToStruct(input, zeroAbc)
		if err == nil {
			t.Errorf("TestToStruct failed: [%v] should not be convertable to struct Abc!", input)
		}
	}
	{
		input := ""
		_, err := Convert(input, zeroAbc)
		if err == nil {
			t.Errorf("TestToStruct failed: [%v] should not be convertable to struct Abc!", input)
		}
	}
}

/*----------------------------------------------------------------------*/
func testToSlice(t *testing.T, input interface{}, expected interface{}, zero interface{}) {
	v, e := ToSlice(input, zero)
	ifFailed(t, "TestToSlice", e)
	from := reflect.ValueOf(v)
	to := reflect.ValueOf(expected)
	if from.Len() != to.Len() {
		t.Errorf("TestToSlice failed: expected [%v] but received [%v]", expected, v)
	}
	for i, n := 0, from.Len(); i < n; i++ {
		if from.Index(i).Interface() != to.Index(i).Interface() {
			t.Errorf("TestToSlice failed: expected [%v] but received [%v]", expected, v)
			break
		}
	}

	v, e = Convert(input, zero)
	ifFailed(t, "TestToSlice", e)
	from = reflect.ValueOf(v)
	to = reflect.ValueOf(expected)
	if from.Len() != to.Len() {
		t.Errorf("TestToSlice failed: expected [%v] but received [%v]", expected, v)
	}
	for i, n := 0, from.Len(); i < n; i++ {
		if from.Index(i).Interface() != to.Index(i).Interface() {
			t.Errorf("TestToSlice failed: expected [%v] but received [%v]", expected, v)
			break
		}
	}
}

// TestToSlice tests if values are converted correctly to slice
func TestToSlice(t *testing.T) {
	{
		input := []bool{true, false}
		zero := [0]int{}
		testToSlice(t, input, []int{1, 0}, zero)
	}
	{
		input := [5]int{-2, 1, 0, 1, 2}
		zero := []string{""}
		testToSlice(t, input, []string{"-2", "1", "0", "1", "2"}, zero)
	}

	{
		input := ""
		_, err := ToSlice(input, [0]int{})
		if err == nil {
			t.Errorf("TestToSlice failed: [%v] should not be convertable to []int!", input)
		}
	}
	{
		input := ""
		_, err := Convert(input, []int{})
		if err == nil {
			t.Errorf("TestToSlice failed: [%v] should not be convertable to []int!", input)
		}
	}
	{
		input := []bool{true, false}
		_, err := ToSlice(input, "")
		if err == nil {
			t.Errorf("TestToSlice failed: [%v] should not be convertable to string!", input)
		}
	}

	{
		input := []string{"a", "b", "c"}
		_, err := ToSlice(input, []int{})
		if err == nil {
			t.Errorf("TestToSlice failed: [%v] should not be convertable to []int!", input)
		}
	}
	{
		input := []string{"a", "b", "c"}
		_, err := Convert(input, []int{})
		if err == nil {
			t.Errorf("TestToSlice failed: [%v] should not be convertable to []int!", input)
		}
	}
}

/*----------------------------------------------------------------------*/
func testToMap(t *testing.T, input interface{}, expected interface{}, zero interface{}) {
	v, e := ToMap(input, zero)
	ifFailed(t, "TestToMap", e)
	from := reflect.ValueOf(v)
	to := reflect.ValueOf(expected)
	if from.Len() != to.Len() {
		t.Errorf("TestToMap failed: expected [%v] but received [%v]", expected, v)
	}
	for _, k := range from.MapKeys() {
		if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
			t.Errorf("TestToMap failed: expected [%v] but received [%v]", expected, v)
		}
	}
	for _, k := range to.MapKeys() {
		if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
			t.Errorf("TestToMap failed: expected [%v] but received [%v]", expected, v)
		}
	}

	v, e = Convert(input, zero)
	ifFailed(t, "TestToMap", e)
	from = reflect.ValueOf(v)
	to = reflect.ValueOf(expected)
	if from.Len() != to.Len() {
		t.Errorf("TestToMap failed: expected [%v] but received [%v]", expected, v)
	}
	for _, k := range from.MapKeys() {
		if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
			t.Errorf("TestToMap failed: expected [%v] but received [%v]", expected, v)
		}
	}
	for _, k := range to.MapKeys() {
		if from.MapIndex(k).Interface() != to.MapIndex(k).Interface() {
			t.Errorf("TestToMap failed: expected [%v] but received [%v]", expected, v)
		}
	}
}

// TestToMap tests if values are converted correctly to map
func TestToMap(t *testing.T) {
	{
		input := map[string]bool{"1": true, "0": false}
		zero := map[int]string{}
		testToMap(t, input, map[int]string{0: "false", 1: "true"}, zero)
	}

	{
		input := map[string]bool{"one": true, "0": false}
		_, err := ToMap(input, map[int]string{})
		if err == nil {
			t.Errorf("TestToMap failed: [%v] should not be convertable to map[int]string!", input)
		}
	}

	{
		input := map[bool]string{true: "1", false: "zero"}
		_, err := ToMap(input, map[bool]int{})
		if err == nil {
			t.Errorf("TestToMap failed: [%v] should not be convertable to map[bool]int!", input)
		}
	}

	{
		input := ""
		_, err := ToMap(input, map[int]string{})
		if err == nil {
			t.Errorf("TestToMap failed: [%v] should not be convertable to map!", input)
		}
	}
	{
		input := ""
		_, err := Convert(input, map[int]string{})
		if err == nil {
			t.Errorf("TestToMap failed: [%v] should not be convertable to map!", input)
		}
	}
	{
		input := map[string]bool{"1": true, "0": false}
		_, err := ToMap(input, "")
		if err == nil {
			t.Errorf("TestToMap failed: [%v] should not be convertable to string!", input)
		}
	}
}

/*----------------------------------------------------------------------*/

// TestToPointer tests if values are converted correctly to pointer
func TestToPointer(t *testing.T) {
	{
		a := float64(1.23)
		zero := int32(0)
		output, err := ToPointer(&a, &zero)
		ifFailed(t, "TestToPointer", err)
		i32 := *output.(*interface{})
		if i32.(int32) != 1 {
			t.Errorf("TestToPointer failed: received [%v]", output)
		}
	}
	{
		a := float64(1.23)
		zero := int32(0)
		output, err := Convert(&a, &zero)
		ifFailed(t, "TestToPointer", err)
		i32 := *output.(*interface{})
		if i32.(int32) != 1 {
			t.Errorf("TestToPointer failed: received [%v]", output)
		}
	}

	{
		a := string("1.23")
		zero := float64(0)
		output, err := ToPointer(&a, &zero)
		ifFailed(t, "TestToPointer", err)
		f64 := *output.(*interface{})
		if f64.(float64) != 1.23 {
			t.Errorf("TestToPointer failed: received [%v]", output)
		}
	}
	{
		a := string("1.23")
		zero := float64(0)
		output, err := Convert(&a, &zero)
		ifFailed(t, "TestToPointer", err)
		f64 := *output.(*interface{})
		if f64.(float64) != 1.23 {
			t.Errorf("TestToPointer failed: received [%v]", output)
		}
	}

	{
		a := string("blabla")
		zero := float64(0)
		_, err := ToPointer(&a, &zero)
		if err == nil {
			t.Errorf("TestToPointer failed: [%v] should not be convertable to [%v]!", &a, &zero)
		}
	}

	{
		a := ""
		zero := int64(0)
		_, err := ToPointer(a, &zero)
		if err == nil {
			t.Errorf("TestToPointer failed: [%v] should not be convertable to [%v]!", a, &zero)
		}
	}

	{
		a := ""
		zero := int64(0)
		_, err := ToPointer(&a, zero)
		if err == nil {
			t.Errorf("TestToPointer failed: [%v] should not be convertable to [%v]!", &a, zero)
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
		output, err := ToPointer(&a, &Abc{})
		ifFailed(t, "TestToPointer", err)
		abc := *output.(*interface{})
		if abc.(Abc).A != 1 {
			t.Errorf("TestToPointer failed: received [%v]", output)
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
		output, err := Convert(&a, &Abc{})
		ifFailed(t, "TestToPointer", err)
		abc := *output.(*interface{})
		if abc.(Abc).A != 1 {
			t.Errorf("TestToPointer failed: received [%v]", output)
		}
	}
}

/*----------------------------------------------------------------------*/

func TestConvert(t *testing.T) {
	{
		_, err := Convert("", nil)
		if err == nil {
			t.Errorf("TestToPointer failed: [%v] should not be convertable to [%v]!", "", nil)
		}
	}
	{
		_, err := Convert(nil, "")
		if err == nil {
			t.Errorf("TestToPointer failed: [%v] should not be convertable to [%v]!", nil, "")
		}
	}
	{
		input := ""
		zero := func() {}
		_, err := Convert(input, zero)
		if err == nil {
			t.Errorf("TestToPointer failed: [%v] should not be convertable to func!", input)
		}
	}

}
