package semita

import (
	"reflect"
	"testing"
)

func TestNode_valueInterface(t *testing.T) {
	v := "a string"
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	if root.unwrap() == nil {
		t.Errorf("TestNode_valueInterface failed, value is nil")
	} else if root.unwrap().(string) != v {
		t.Errorf("TestNode_valueInterface failed, expected value is %#v, but received %#v", v, root.unwrap())
	}
}

func TestNode_nextMap(t *testing.T) {
	v := map[string]interface{}{
		"s": "string",
		"i": 103,
		"a": []int{1, 2, 3},
		"m": map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
	}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	node, err := root.next("xyz")
	if node != nil || err != nil {
		t.Errorf("TestNode_nextMap failed with data %#v at index %#v", v, "xyz")
	}
}
