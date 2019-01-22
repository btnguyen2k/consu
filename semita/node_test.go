package semita

import (
	"reflect"
	"strings"
	"testing"
)

func Test_createEmptyGenericSlice(t *testing.T) {
	v := createEmptyGenericSlice()
	if v.Kind() != reflect.Slice || v.Len() != 0 || v.Type().Elem().Kind() != reflect.Interface {
		t.Errorf("Test_createEmptyGenericSlice failed, expected empty generic slice, but received {%#v}", v)
	}
}

func Test_createEmptyGenericMap(t *testing.T) {
	v := createEmptyGenericMap()
	if v.Kind() != reflect.Map || v.Len() != 0 || v.Type().Key().Kind() != reflect.String || v.Type().Elem().Kind() != reflect.Interface {
		t.Errorf("Test_createEmptyGenericMap failed, expected empty generic map, but received {%#v}", v)
	}
}

func TestNode_unwrap(t *testing.T) {
	{
		v := true
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(bool) != v {
			t.Errorf("TestNode_unwrap failed, expected {%#v}, but received {%#v}", v, uv)
		}
	}
	{
		v := 1
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(int) != v {
			t.Errorf("TestNode_unwrap failed, expected {%#v}, but received {%#v}", v, uv)
		}
	}
	{
		v := 1.2
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(float64) != v {
			t.Errorf("TestNode_unwrap failed, expected {%#v}, but received {%#v}", v, uv)
		}
	}
	{
		v := "123"
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(string) != v {
			t.Errorf("TestNode_unwrap failed, expected {%#v}, but received {%#v}", v, uv)
		}
	}
}

func TestNode_nextInvalid(t *testing.T) {
	{
		v := false
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("path")
		if err == nil {
			t.Errorf("TestNode_nextInvalid failed for value {%#v}", v)
		}
	}
	{
		v := 1
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("path")
		if err == nil {
			t.Errorf("TestNode_nextInvalid failed for value {%#v}", v)
		}
	}
	{
		v := 1.2
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("path")
		if err == nil {
			t.Errorf("TestNode_nextInvalid failed for value {%#v}", v)
		}
	}
	{
		v := "123"
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("path")
		if err == nil {
			t.Errorf("TestNode_nextInvalid failed for value {%#v}", v)
		}
	}
}

type Inner struct {
	b bool
	f float64
	i int
	s string
}

type Outter struct {
	a []int
	b [3]string
	m map[string]interface{}
	s Inner
}

func genDataOuter() Outter {
	return Outter{
		a: []int{0, 1, 2, 3, 4, 5},
		b: [3]string{"a", "b", "c"},
		m: map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		s: Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
	}
}

func genDataMap() map[string]interface{} {
	return map[string]interface{}{
		"a": []int{0, 1, 2, 3, 4, 5},
		"b": [3]string{"a", "b", "c"},
		"m": map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		"s": genDataOuter(),
	}
}

func TestNode_nextMap(t *testing.T) {
	v := genDataMap()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}

	p = "not exist"
	node, err = root.next(p)
	if node != nil || err != nil {
		// non-exists entry
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}

	for _, path := range []string{"a.[0]", "b.[1]", "m.z", "s.a.[0]", "s.b.[1]", "s.m.z", "s.s.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextMap failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func TestNode_nextStruct(t *testing.T) {
	v := genDataOuter()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "not exist"
	node, err = root.next(p)
	if node != nil || err != nil {
		// non-exists entry
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}

	for _, path := range []string{"a.[0]", "b.[1]", "m.z", "s.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextStruct failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}
