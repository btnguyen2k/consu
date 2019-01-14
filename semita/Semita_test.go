package semita

import (
	"testing"
)

func ifFailed(t *testing.T, f string, e error) {
	if e != nil {
		t.Errorf("%s failed: %e", f, e)
	}
}

/*----------------------------------------------------------------------*/

// TestNewSemita test if Semita instance can be created correctly.
func TestNewSemita(t *testing.T) {
	// only Array, Slice, Map and Struct can be wrapped
	{
		input := struct {
			a int
			b string
			c bool
		}{a: 1, b: "2", c: true}
		s := NewSemita(input)
		if s == nil {
			t.Errorf("TestNewSemita failed for input %#v", input)
		}
	}
	{
		input := map[string]interface{}{}
		s := NewSemita(input)
		if s == nil {
			t.Errorf("TestNewSemita failed for input %#v", input)
		}
	}
	{
		input := [3]int{1, 2, 3}
		s := NewSemita(input)
		if s == nil {
			t.Errorf("TestNewSemita failed for input %#v", input)
		}
	}
	{
		input := []string{"a", "b", "c"}
		s := NewSemita(input)
		if s == nil {
			t.Errorf("TestNewSemita failed for input %#v", input)
		}
	}

	{
		input := 1
		s := NewSemita(input)
		if s != nil {
			t.Errorf("TestNewSemita failed for input %#v", input)
		}
	}
	{
		input := "string"
		s := NewSemita(input)
		if s != nil {
			t.Errorf("TestNewSemita failed for input %#v", input)
		}
	}
	{
		input := false
		s := NewSemita(input)
		if s != nil {
			t.Errorf("TestNewSemita failed for input %#v", input)
		}
	}
}

/*----------------------------------------------------------------------*/

func testSplitPath(t *testing.T, path string, expected []string) {
	tokens := SplitPath(path)
	if len(tokens) != len(expected) {
		t.Errorf("TestSplitPath failed for input [%s], expected %#v but received %#v.", path, expected, tokens)
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

func TestSemita_GetValueArray(t *testing.T) {
	input := [3]int{1, 2, 3}
	s := NewSemita(input)

	{
		// index out-of-bound
		p := "-1"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueArray getting value at [%#v] for input %#v", p, input)
		}
	}
	{
		// index out-of-bound
		p := "3"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueArray getting value at [%#v] for input %#v", p, input)
		}
	}

	{
		p := "[0]"
		v, e := s.GetValue(p)
		if e != nil || v != input[0] {
			t.Errorf("TestSemita_GetValueArray getting value at [%#v] for input %#v", p, input)
		}
	}
}

func TestSemita_GetValueSlice(t *testing.T) {
	input := []string{"1", "2", "3"}
	s := NewSemita(input)

	{
		// index out-of-bound
		p := "-1"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueSlice getting value at [%#v] for input %#v", p, input)
		}
	}
	{
		// index out-of-bound
		p := "3"
		_, e := s.GetValue(p)
		if e == nil {
			t.Errorf("TestSemita_GetValueSlice getting value at [%#v] for input %#v", p, input)
		}
	}

	{
		p := "[0]"
		v, e := s.GetValue(p)
		if e != nil || v != input[0] {
			t.Errorf("TestSemita_GetValueSlice getting value at [%#v] for input %#v", p, input)
		}
	}
}

func TestSemita_GetValueMap(t *testing.T) {
	input := map[string]interface{}{
		"a": "string",
		"b": 1,
		"c": true,
	}
	s := NewSemita(input)

	{
		p := "a"
		v, e := s.GetValue(p)
		if e != nil || v != input[p] {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for input %#v", p, input)
		}
	}
	{
		p := "b"
		v, e := s.GetValue(p)
		if e != nil || v != input[p] {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for input %#v", p, input)
		}
	}
	{
		p := "b"
		v, e := s.GetValue(p)
		if e != nil || v != input[p] {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for input %#v", p, input)
		}
	}

	{
		p := "z"
		v, e := s.GetValue(p)
		if v != nil && e != nil {
			t.Errorf("TestSemita_GetValueMap getting value at [%#v] for input %#v", p, input)
		}
	}
}

func TestSemita_GetValueStruct(t *testing.T) {
	type MyStruct struct {
		A string
		B int
		C bool
		x string // un-exported field
	}
	input := MyStruct{
		A: "string",
		B: 1,
		C: true,
		x: "another string",
	}
	s := NewSemita(input)

	{
		p := "A"
		v, e := s.GetValue(p)
		if e != nil || v != input.A {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for input %#v", p, input)
		}
	}
	{
		p := "B"
		v, e := s.GetValue(p)
		if e != nil || v != input.B {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for input %#v", p, input)
		}
	}
	{
		p := "C"
		v, e := s.GetValue(p)
		if e != nil || v != input.C {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for input %#v", p, input)
		}
	}
	{
		p := "x"
		v, e := s.GetValue(p)
		if e != nil || v != input.x {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for input %#v", p, input)
		}
	}

	{
		p := "z"
		v, e := s.GetValue(p)
		if v != nil && e != nil {
			t.Errorf("TestSemita_GetValueStruct getting value at [%#v] for input %#v", p, input)
		}
	}
}
