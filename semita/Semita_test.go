package semita

import (
	"github.com/btnguyen2k/consu/reddo"
	"reflect"
	"strconv"
	"testing"
	"time"
	"unsafe"
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
func TestCreateZero_Invalid(t *testing.T) {
	name := "TestCreateZero_Invalid"
	{
		z := CreateZero(nil)
		if z.IsValid() {
			t.Errorf("%s failed", name)
		}
	}
	{
		temp := "a string"
		v := &temp
		z := CreateZero(reflect.TypeOf(v))
		if z.IsValid() {
			t.Errorf("%s failed for data %#v %T", name, v, v)
		}
	}
}

func TestCreateZero_Primitives(t *testing.T) {
	name := "TestCreateZero_Primitives"
	{
		vList := []interface{}{false, int(0), int8(0), int16(0), int32(0), int64(0)}
		for _, v := range vList {
			z := CreateZero(reflect.TypeOf(v))
			if !z.IsValid() || z.Interface() != v {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
	{
		vList := []interface{}{uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0)}
		for _, v := range vList {
			z := CreateZero(reflect.TypeOf(v))
			if !z.IsValid() || z.Interface() != v {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
	{
		vList := []interface{}{uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0)}
		for _, v := range vList {
			z := CreateZero(reflect.TypeOf(v))
			if !z.IsValid() || z.Interface() != v {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
	{
		vList := []interface{}{float32(0.0), float64(0.0)}
		for _, v := range vList {
			z := CreateZero(reflect.TypeOf(v))
			if !z.IsValid() || z.Interface() != v {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
	{
		vList := []interface{}{complex64(0 + 0i), complex128(0 + 0i)}
		for _, v := range vList {
			z := CreateZero(reflect.TypeOf(v))
			if !z.IsValid() || z.Interface() != v {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
	{
		vList := []interface{}{"", unsafe.Pointer(nil)}
		for _, v := range vList {
			z := CreateZero(reflect.TypeOf(v))
			if !z.IsValid() || z.Interface() != v {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
}

func TestCreateZero_SliceAndArray(t *testing.T) {
	name := "TestCreateZero_SliceAndArray"
	{
		v := []int{1, 2, 3}
		z := CreateZero(reflect.TypeOf(v))
		if !z.IsValid() || z.Kind() != reflect.Slice || z.Len() != 0 {
			t.Errorf("%s failed for data %#v %T", name, v, v)
		}
		z = reflect.Append(z, reflect.ValueOf(0))
		z = reflect.Append(z, reflect.ValueOf(0))
		z = reflect.Append(z, reflect.ValueOf(0))
		if z.Len() != 3 {
			t.Errorf("%s failed for data %#v %T", name, v, v)
		}
		for i := 0; i < len(v); i++ {
			z.Index(i).Set(reflect.ValueOf(v[i]))
		}
		for i := 0; i < len(v); i++ {
			if z.Index(i).Int() != int64(v[i]) {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
	{
		v := [4]string{"0", "a", "false", "true"}
		z := CreateZero(reflect.TypeOf(v))
		if !z.IsValid() || z.Kind() != reflect.Slice || z.Len() != 0 {
			t.Errorf("%s failed for data %#v %T", name, v, v)
		}
		z = reflect.Append(z, reflect.ValueOf(""))
		z = reflect.Append(z, reflect.ValueOf(""))
		z = reflect.Append(z, reflect.ValueOf(""))
		z = reflect.Append(z, reflect.ValueOf(""))
		if z.Len() != 4 {
			t.Errorf("%s failed for data %#v %T", name, v, v)
		}
		for i := 0; i < len(v); i++ {
			z.Index(i).Set(reflect.ValueOf(v[i]))
		}
		for i := 0; i < len(v); i++ {
			if z.Index(i).String() != v[i] {
				t.Errorf("%s failed for data %#v %T", name, v, v)
			}
		}
	}
}

func TestCreateZero_Map(t *testing.T) {
	name := "TestCreateZero_Map"
	v := map[int]string{0: "", 1: "one", 2: "2"}
	z := CreateZero(reflect.TypeOf(v))
	if !z.IsValid() || z.Kind() != reflect.Map || z.Len() != 0 {
		t.Errorf("%s failed for data %#v %T", name, v, v)
	}
	z.SetMapIndex(reflect.ValueOf(0), reflect.ValueOf(""))
	z.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf("one"))
	z.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf("2"))
	if z.Len() != 3 {
		t.Errorf("%s failed for data %#v %T", name, v, v)
	}
	for i := 0; i < len(v); i++ {
		if z.MapIndex(reflect.ValueOf(i)).String() != v[i] {
			t.Errorf("%s failed for data %#v %T", name, v, v)
		}
	}
}

func TestCreateZero_Struct(t *testing.T) {
	name := "TestCreateZero_Struct"
	type MyStruct struct {
		I int
		s string
		B bool
		A []int
	}
	v := MyStruct{I: 103, s: "btnguyen2k", B: true, A: []int{1, 2, 3}}
	z := CreateZero(reflect.TypeOf(v))
	if !z.IsValid() || z.Kind() != reflect.Struct {
		t.Errorf("%s failed for data %#v %T", name, v, v)
	}
	z.FieldByName("I").Set(reflect.ValueOf(103))
	// z.FieldByName("s").Set(reflect.ValueOf("btnguyen2k"))
	z.FieldByName("B").Set(reflect.ValueOf(true))
	z.FieldByName("A").Set(reflect.ValueOf([]int{1, 2, 3}))
	if z.FieldByName("I").Int() != int64(v.I) {
		t.Errorf("%s failed for data %#v %T", name, v, v)
	}
	if z.FieldByName("B").Bool() != v.B {
		t.Errorf("%s failed for data %#v %T", name, v, v)
	}
	if !reflect.DeepEqual(z.FieldByName("A").Interface(), v.A) {
		t.Errorf("%s failed for data %#v %T", name, v, v)
	}
}

/*----------------------------------------------------------------------*/

func TestGetTypeOfMapKey(t *testing.T) {
	name := "TestGetTypeOfMapKey"
	{
		v := map[bool]string{true: "true", false: "false"}
		typ := GetTypeOfMapKey(v)
		if typ == nil || typ.Kind() != reflect.Bool {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfMapKey(&v)
		if typ == nil || typ.Kind() != reflect.Bool {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := map[int]string{1: "one", 2: "two"}
		typ := GetTypeOfMapKey(v)
		if typ == nil || typ.Kind() != reflect.Int {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfMapKey(&v)
		if typ == nil || typ.Kind() != reflect.Int {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := map[uint]string{1: "1", 2: "2"}
		typ := GetTypeOfMapKey(v)
		if typ == nil || typ.Kind() != reflect.Uint {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfMapKey(&v)
		if typ == nil || typ.Kind() != reflect.Uint {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := map[string]string{"1": "one", "2": "two"}
		typ := GetTypeOfMapKey(v)
		if typ == nil || typ.Kind() != reflect.String {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfMapKey(&v)
		if typ == nil || typ.Kind() != reflect.String {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}

	{
		v := "this is not a map"
		typ := GetTypeOfMapKey(v)
		if typ != nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfMapKey(&v)
		if typ != nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
}

func TestGetTypeOfElement(t *testing.T) {
	name := "TestGetTypeOfElement"
	{
		v := map[string]bool{"true": true, "false": false}
		typ := GetTypeOfElement(v)
		if typ == nil || typ.Kind() != reflect.Bool {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfElement(&v)
		if typ == nil || typ.Kind() != reflect.Bool {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := []int{1, 2, 3}
		typ := GetTypeOfElement(v)
		if typ == nil || typ.Kind() != reflect.Int {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfElement(&v)
		if typ == nil || typ.Kind() != reflect.Int {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := [4]uint{0, 1, 2, 3}
		typ := GetTypeOfElement(v)
		if typ == nil || typ.Kind() != reflect.Uint {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfElement(&v)
		if typ == nil || typ.Kind() != reflect.Uint {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := "a string"
		typ := GetTypeOfElement(v)
		if typ == nil || typ.Kind() != reflect.String {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfElement(&v)
		if typ == nil || typ.Kind() != reflect.String {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
}

func TestGetTypeOfStructAttibute(t *testing.T) {
	name := "TestGetTypeOfStructAttibute"
	type MyStruct struct {
		FieldInt     int
		FieldBool    bool
		FieldString  string
		fieldPrivate interface{}
	}
	v := MyStruct{FieldInt: 1, FieldBool: true, FieldString: "a string", fieldPrivate: 0.1}
	{
		typ := GetTypeOfStructAttibute(v, "FieldInt")
		if typ == nil || typ.Kind() != reflect.Int {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfStructAttibute(&v, "FieldInt")
		if typ == nil || typ.Kind() != reflect.Int {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		typ := GetTypeOfStructAttibute(v, "FieldBool")
		if typ == nil || typ.Kind() != reflect.Bool {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfStructAttibute(&v, "FieldBool")
		if typ == nil || typ.Kind() != reflect.Bool {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		typ := GetTypeOfStructAttibute(v, "FieldString")
		if typ == nil || typ.Kind() != reflect.String {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfStructAttibute(&v, "FieldString")
		if typ == nil || typ.Kind() != reflect.String {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		typ := GetTypeOfStructAttibute(v, "fieldPrivate")
		if typ == nil || typ.Kind() != reflect.Interface {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfStructAttibute(&v, "fieldPrivate")
		if typ == nil || typ.Kind() != reflect.Interface {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		typ := GetTypeOfStructAttibute(v, "invalid")
		if typ != nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfStructAttibute(&v, "invalid")
		if typ != nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := "this is not a struct"
		typ := GetTypeOfStructAttibute(v, "invalid")
		if typ != nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
		typ = GetTypeOfStructAttibute(&v, "invalid")
		if typ != nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
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

func TestSemita_GetTimeError(t *testing.T) {
	name := "TestSemita_GetTimeError"

	data := map[string]interface{}{
		"val_int": -1,
		"val_str": "-1",
	}
	s1 := NewSemita(data)
	s2 := NewSemita(&data)

	{
		p := "val_int"
		_, e := s1.GetTime(p)
		if e == nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		}
		_, e = s2.GetTime(p)
		if e == nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		}
	}
	{
		p := "val_str"
		_, e := s1.GetTime(p)
		if e == nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		}
		_, e = s2.GetTime(p)
		if e == nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		}
	}
}

func TestSemita_GetTime(t *testing.T) {
	name := "TestSemita_GetTime"

	now := time.Now()
	data := map[string]interface{}{
		"val_int": now.Unix(),
		"val_str": strconv.FormatInt(now.Unix(), 10),
	}
	s1 := NewSemita(data)
	s2 := NewSemita(&data)

	{
		p := "val_int"
		v, e := s1.GetTime(p)
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
		v, e = s2.GetTime(p)
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
	{
		p := "val_str"
		v, e := s1.GetTime(p)
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
		v, e = s2.GetTime(p)
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}
}

func TestSemita_GetTimeWithLayout(t *testing.T) {
	name := "TestSemita_GetTimeWithLayout"

	now := time.Now()
	input := "2019-04-29T20:59:10"
	layout := "2006-01-02T15:04:05"
	expected := time.Date(2019, 04, 29, 20, 59, 10, 0, time.UTC)
	data := map[string]interface{}{
		"val_int": now.Unix(),
		"val_str": strconv.FormatInt(now.Unix(), 10),
		"input":   input,
	}
	s1 := NewSemita(data)
	s2 := NewSemita(&data)

	{
		p := "val_int"
		v, e := s1.GetTimeWithLayout(p, "")
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
		v, e = s2.GetTimeWithLayout(p, "")
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		p := "val_str"
		v, e := s1.GetTimeWithLayout(p, "")
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
		v, e = s2.GetTimeWithLayout(p, "")
		if e != nil {
			t.Errorf("%s failed with data %v at path %s", name, data, p)
		} else if v.Unix() != now.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, now, v)
		}
	}

	{
		p := "input"
		v, e := s1.GetTimeWithLayout(p, layout)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
		} else if v.Unix() != expected.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
		v, e = s2.GetTimeWithLayout(p, layout)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
		} else if v.Unix() != expected.Unix() {
			t.Errorf("%s failed: expected [%#v] but received [%#v]", name, expected, v)
		}
	}
}

/*----------------------------------------------------------------------*/

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
		Overtime  bool
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

/*----------------------------------------------------------------------*/

func TestSemita_GetValueOfType_Invalid(t *testing.T) {
	name := "TestSemita_GetValueOfType_Invalid"
	p := "not_exists"

	data := generateDataMap()
	s1 := NewSemita(data)
	{
		v, _ := s1.GetValueOfType(p, reddo.TypeString)
		if v != nil {
			t.Errorf("%s getting value at {%#v} for data {%#v}", name, p, data)
		}
	}

	d := data.(map[string]interface{})
	s2 := NewSemita(&d)
	{
		v, _ := s2.GetValueOfType(p, reddo.TypeString)
		if v != nil {
			t.Errorf("%s getting value at {%#v} for data {%#v}", name, p, data)
		}
	}
}

func TestSemita_GetValueOfType_MultiLevelMap(t *testing.T) {
	name := "TestSemita_GetValueOfType_MultiLevelMap"
	data := generateDataMap()
	s1 := NewSemita(data)
	d := data.(map[string]interface{})
	s2 := NewSemita(&d)

	{
		p := "Name"
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != companyName {
			t.Errorf("%s getting value at {%#v} for data {%#v}", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != companyName {
			t.Errorf("%s getting value at {%#v} for data {%#v}", name, p, data)
		}
	}
	{
		p := "Year"
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].age"
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].email"
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].options.work_hours"
		v, e := s1.GetValueOfType(p, reflect.TypeOf([]int{}))
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
		v, e = s2.GetValueOfType(p, reflect.TypeOf([]int{}))
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
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
		v, e := s1.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].join_date"
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].join_date"
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
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
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Year"
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].Age"
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].Email"
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].Options.WorkHours"
		v, e := s1.GetValueOfType(p, reflect.TypeOf([]int{}))
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
		v, e = s2.GetValueOfType(p, reflect.TypeOf([]int{}))
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
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
		v, e := s1.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].JoinDate"
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].JoinDate"
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
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
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != companyName {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Year"
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(companyYear) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].age"
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(employee0Age) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].email"
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != employee1Email {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[0].options.WorkHours"
		v, e := s1.GetValueOfType(p, reflect.TypeOf([]int{}))
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if len(v.([]int)) != len(employee0WorkHours) {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		for i, n := 0, len(employee0WorkHours); i < n; i++ {
			if employee0WorkHours[i] != v.([]int)[i] {
				t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
			}
		}
		v, e = s2.GetValueOfType(p, reflect.TypeOf([]int{}))
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
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
		v, e := s1.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != employee1Overtime {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees.[0].join_date"
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		p := "Employees[1].join_date"
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s getting value at [%#v] for data %#v", name, p, data)
		}
	}
}

/*----------------------------------------------------------------------*/

func TestSemita_SetValue_MultiLevelMap(t *testing.T) {
	name := "TestSemita_SetValue_MultiLevelMap"

	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Name"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Year"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Employees.[0].age"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Employees[1].email"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Employees[0].options.work_hours.[0]"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Employees.[1].options.overtime"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Employees.[0].join_date"
		d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
		d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)

		vSet1 := d1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := d0
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
}

func TestSemita_SetValue_MultiLevelStruct(t *testing.T) {
	name := "TestSemita_SetValue_MultiLevelStruct"

	{
		data := generateDataStruct()
		// s1 := NewSemita(data)
		data2 := generateDataStruct().(Company)
		s2 := NewSemita(&data2)
		p := "Name"

		// vSet1 := "1"
		// e := s1.SetValue(p, vSet1)
		// ifFailed(t, name, e)
		// v, e := s1.GetValueOfType(p, reddo.TypeString)
		// ifFailed(t, name, e)
		// if v.(string) != vSet1 {
		// 	t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		// }

		vSet2 := "2"
		e := s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataStruct()
		// s1 := NewSemita(data)
		data2 := generateDataStruct().(Company)
		s2 := NewSemita(&data2)
		p := "Year"

		// vSet1 := "1"
		// e := s1.SetValue(p, vSet1)
		// ifFailed(t, name, e)
		// v, e := s1.GetValueOfType(p, reddo.TypeString)
		// ifFailed(t, name, e)
		// if v.(string) != vSet1 {
		// 	t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		// }

		vSet2 := 2
		e := s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataStruct()
		s1 := NewSemita(data)
		data2 := generateDataStruct().(Company)
		s2 := NewSemita(&data2)
		p := "Employees.[0].Age"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataStruct()
		s1 := NewSemita(data)
		data2 := generateDataStruct().(Company)
		s2 := NewSemita(&data2)
		p := "Employees[1].Email"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataStruct()
		s1 := NewSemita(data)
		data2 := generateDataStruct().(Company)
		s2 := NewSemita(&data2)
		p := "Employees[0].Options.WorkHours.[0]"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataStruct()
		s1 := NewSemita(data)
		data2 := generateDataStruct().(Company)
		s2 := NewSemita(&data2)
		p := "Employees.[1].Options.Overtime"

		vSet1 := !employee1Overtime
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := !vSet1
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeBool)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(bool) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataStruct()
		s1 := NewSemita(data)
		data2 := generateDataStruct().(Company)
		s2 := NewSemita(&data2)
		p := "Employees.[0].JoinDate"
		d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
		d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)

		vSet1 := d1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := d0
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
}

func TestSemita_SetValue_MultiLevelMixed(t *testing.T) {
	name := "TestSemita_SetValue_MultiLevelMixed"

	{
		data := generateDataMixed()
		// s1 := NewSemita(data)
		data2 := generateDataMixed().(CompanyMixed)
		s2 := NewSemita(&data2)
		p := "Name"

		// vSet1 := "1"
		// e := s1.SetValue(p, vSet1)
		// ifFailed(t, name, e)
		// v, e := s1.GetValueOfType(p, reddo.TypeString)
		// ifFailed(t, name, e)
		// if v.(string) != vSet1 {
		// 	t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		// }

		vSet2 := "2"
		e := s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMixed()
		// s1 := NewSemita(data)
		data2 := generateDataMixed().(CompanyMixed)
		s2 := NewSemita(&data2)
		p := "Year"

		// vSet1 := 1
		// e := s1.SetValue(p, vSet1)
		// if e != nil {
		// 	t.Errorf("%s failed: %e", name, e)
		// 	t.FailNow()
		// }
		// v, e := s1.GetValueOfType(p, reddo.TypeInt)
		// if e != nil {
		// 	t.Errorf("%s failed: %e", name, e)
		// 	t.FailNow()
		// }
		// if v.(int64) != int64(vSet1) {
		// 	t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		// }

		vSet2 := 2
		e := s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMixed()
		s1 := NewSemita(data)
		data2 := generateDataMixed().(CompanyMixed)
		s2 := NewSemita(&data2)
		p := "Employees.[0].age"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMixed()
		s1 := NewSemita(data)
		data2 := generateDataMixed().(CompanyMixed)
		s2 := NewSemita(&data2)
		p := "Employees[1].email"

		vSet1 := "1"
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := "2"
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != vSet2 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMixed()
		s1 := NewSemita(data)
		data2 := generateDataMixed().(CompanyMixed)
		s2 := NewSemita(&data2)
		p := "Employees[0].options.WorkHours.[0]"

		vSet1 := 1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 2
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	// {
	// 	p := "Employees.[1].options.Overtime"
	//
	// 	vSet1 := !employee1Overtime
	// 	e := s1.SetValue(p, vSet1)
	// 	if e != nil {
	// 		t.Errorf("%s failed: %e", name, e)
	// 		t.FailNow()
	// 	}
	// 	v, e := s1.GetValueOfType(p, reddo.TypeBool)
	// 	if e != nil {
	// 		t.Errorf("%s failed: %e", name, e)
	// 		t.FailNow()
	// 	}
	// 	if v.(bool) != vSet1 {
	// 		t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
	// 	}
	//
	// 	vSet2 := !vSet1
	// 	e = s2.SetValue(p, vSet2)
	// 	if e != nil {
	// 		t.Errorf("%s failed: %e", name, e)
	// 		t.FailNow()
	// 	}
	// 	v, e = s2.GetValueOfType(p, reddo.TypeBool)
	// 	if e != nil {
	// 		t.Errorf("%s failed: %e", name, e)
	// 		t.FailNow()
	// 	}
	// 	if v.(bool) != vSet2 {
	// 		t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
	// 	}
	// }
	{
		data := generateDataMixed()
		s1 := NewSemita(data)
		data2 := generateDataMixed().(CompanyMixed)
		s2 := NewSemita(&data2)
		p := "Employees.[0].join_date"
		d0, _ := time.Parse(employee0JoinDateFormat, employee0JoinDate)
		d1, _ := time.Parse(employee1JoinDateFormat, employee1JoinDate)

		vSet1 := d1
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee1JoinDateFormat) != employee1JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := d0
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeTime)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(time.Time).Format(employee0JoinDateFormat) != employee0JoinDate {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
}

/*----------------------------------------------------------------------*/

func TestSemita_SetValue_MultiLevelMap_CreateNodes(t *testing.T) {
	name := "TestSemita_SetValue_MultiLevelMap_CreateNodes"

	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "Employees.[].age" // append to end of slice

		_v, _ := s1.GetValueOfType("Employees", reflect.TypeOf([]interface{}{}))
		l1 := len(_v.([]interface{}))
		vSet1 := 19
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		_v, _ = s1.GetValueOfType("Employees", reflect.TypeOf([]interface{}{}))
		l2 := len(_v.([]interface{}))
		if l2 != l1+1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
		_p := "Employees[" + strconv.Itoa(l2-1) + "].age"
		v, e := s1.GetValueOfType(_p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(int64) != int64(vSet1) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		_v, _ = s2.GetValueOfType("Employees", reflect.TypeOf([]interface{}{}))
		l1 = len(_v.([]interface{}))
		vSet2 := 81
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		_v, _ = s2.GetValueOfType("Employees", reflect.TypeOf([]interface{}{}))
		l2 = len(_v.([]interface{}))
		if l2 != l1+1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
		_p = "Employees[" + strconv.Itoa(l2-1) + "].age"
		v, e = s2.GetValueOfType(_p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "a.b.c.d" // create all nodes for maps

		vSet1 := "19"
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if strconv.Itoa(int(v.(int64))) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 81
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != strconv.Itoa(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
	{
		data := generateDataMap()
		s1 := NewSemita(data)
		data2 := generateDataMap().(map[string]interface{})
		s2 := NewSemita(&data2)
		p := "a[].b.c[].d" // create all nodes for maps & slices
		_p := "a[0].b.c[0].d"

		vSet1 := "19"
		e := s1.SetValue(p, vSet1)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e := s1.GetValueOfType(_p, reddo.TypeInt)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if strconv.Itoa(int(v.(int64))) != vSet1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}

		vSet2 := 81
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		v, e = s2.GetValueOfType(_p, reddo.TypeString)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(string) != strconv.Itoa(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
}

func TestSemita_SetValue_MultiLevelMixed_CreateNodes(t *testing.T) {
	name := "TestSemita_SetValue_MultiLevelMixed_CreateNodes"

	{
		data := generateDataMixed()
		s1 := NewSemita(data)
		data2 := generateDataMixed().(CompanyMixed)
		s2 := NewSemita(&data2)
		p := "Employees.[].age" // append to end of slice

		vSet1 := 19
		e := s1.SetValue(p, vSet1)
		if e == nil {
			// s1 is not reference to strut --> can not append
			t.Errorf("%s failed: %e", name, e)
		}

		_v, _ := s2.GetValueOfType("Employees", reflect.TypeOf([]interface{}{}))
		l1 := len(_v.([]interface{}))
		vSet2 := 81
		e = s2.SetValue(p, vSet2)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		_v, _ = s2.GetValueOfType("Employees", reflect.TypeOf([]interface{}{}))
		l2 := len(_v.([]interface{}))
		if l2 != l1+1 {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
		_p := "Employees[" + strconv.Itoa(l2-1) + "].age"
		v, e := s2.GetValueOfType(_p, reddo.TypeUint)
		if e != nil {
			t.Errorf("%s failed: %e", name, e)
			t.FailNow()
		}
		if v.(uint64) != uint64(vSet2) {
			t.Errorf("%s setting value at [%#v] for data %#v", name, p, data)
		}
	}
}
