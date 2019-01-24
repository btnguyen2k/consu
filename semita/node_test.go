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
		_, err := n.next("[0]")
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
		_, err := n.next("[0]")
		if err == nil {
			t.Errorf("TestNode_nextInvalid failed for value {%#v}", v)
		}
	}
	{
		v := map[int]string{1: "one"}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("1")
		if err == nil {
			t.Errorf("TestNode_nextInvalid failed for value {%#v}", v)
		}
	}
	{
		v := []int{0, 1, 2, 3}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("[a]")
		if err == nil {
			t.Errorf("TestNode_nextInvalid failed for value {%#v}", v)
		}
	}
}

type Inner struct {
	b interface{}
	f interface{}
	i interface{}
	s interface{}
}

type Outter struct {
	A       interface{}
	B       interface{}
	M       interface{}
	S       interface{}
	private interface{}
}

func genDataOuter() Outter {
	return Outter{
		A: []int{0, 1, 2, 3, 4, 5},
		B: [3]string{"a", "b", "c"},
		M: map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		S:       Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
		private: 1.2,
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
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}, error: %e", v, p, err)
	}

	for _, path := range []string{"a.[0]", "b.[1]", "m.z", "s.A.[0]", "s.B.[1]", "s.M.z", "s.S.s"} {
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

	for _, path := range []string{"A.[0]", "B.[1]", "M.z", "S.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextStruct failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func genDataSlice() []interface{} {
	return []interface{}{
		"a string",
		103,
		19.81,
		true,
		[]int{0, 1, 2, 3, 4, 5},
		[3]string{"a", "b", "c"},
		map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		genDataOuter(),
	}
}

func TestNode_nextSlice(t *testing.T) {
	v := genDataSlice()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "abc"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}

	for _, path := range []string{"[4].[0]", "[5].[1]", "[6].z", "[7].A.[0]", "[7].B.[1]", "[7].M.z", "[7].S.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextMap failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func genDataArray() [8]interface{} {
	return [8]interface{}{
		"a string",
		103,
		19.81,
		true,
		[]int{0, 1, 2, 3, 4, 5},
		[3]string{"a", "b", "c"},
		map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		genDataOuter(),
	}
}

func TestNode_nextArray(t *testing.T) {
	v := genDataArray()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "abc"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}

	for _, path := range []string{"[4].[0]", "[5].[1]", "[6].z", "[7].A.[0]", "[7].B.[1]", "[7].M.z", "[7].S.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextMap failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func TestNode_setValueInvalid(t *testing.T) {
	{
		v := false
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("path", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
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
		_, err := n.setValue("[0]", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
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
		_, err := n.setValue("path", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
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
		_, err := n.setValue("[0]", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
		}
	}
	{
		v := map[int]string{1: "one"}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("1", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
		}
	}
}

func TestNode_setValueMapInvalidType(t *testing.T) {
	{
		v := map[int]string{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		node, err := root.setValue("1", reflect.ValueOf("string"))
		if node != nil || err == nil {
			// invalid key type
			t.Errorf("TestNode_setValueMapInvalidType failed with data %#v", v)
		}
	}
	{
		v := map[string]int{"one": 1}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("two", reflect.ValueOf("2"))
		if err == nil {
			// invalid element type
			t.Errorf("TestNode_setValueMapInvalidType failed for value {%#v}", v)
		}
	}
}

func TestNode_setValueMap(t *testing.T) {
	v := genDataMap()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}

	p = "notExist"
	node, err = root.setValue(p, data)
	if node == nil || err != nil || node.unwrap() != data.Interface() {
		// non-exists entry
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"a", "b", "m", "s"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
		}
	}
}

func TestNode_setValueStructInvalidType(t *testing.T) {
	type MyStruct struct {
		S string
		I int
	}

	v := MyStruct{S: "string", I: 123}
	n := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	_, err := n.setValue("S", reflect.ValueOf(1981))
	if err == nil {
		// invalid element type
		t.Errorf("TestNode_setValueStructInvalidType failed for value {%#v}", v)
	}
}

func TestNode_setValueStruct(t *testing.T) {
	type MyStruct struct {
		A       interface{}
		B       interface{}
		M       interface{}
		S       interface{}
		private interface{}
	}
	v := MyStruct{
		A: []int{0, 1, 2, 3, 4, 5},
		B: [3]string{"a", "b", "c"},
		M: map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		S:       Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
		private: 1.2,
	}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(&v), // for struct: only addressable struct is settable
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "notExist"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// non-exists entry
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "private"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// non-exists entry
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"A", "B", "M", "S"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
		}
	}
}

func TestNode_setValueSliceInvalidType(t *testing.T) {
	v := []string{"a", "b", "c"}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	{
		node, err := root.setValue("[1]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueSliceInvalidType failed with data %#v", v)
		}
	}
	{
		node, err := root.setValue("[a]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueSliceInvalidType failed with data %#v", v)
		}
	}
}

func TestNode_setValueSlice(t *testing.T) {
	v := genDataSlice()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "xyz"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}

	l := root.value.Len()
	p = "[]"
	node, err = root.setValue(p, data)
	if node == nil || err != nil || node.unwrap() != data.Interface() || root.value.Len() != l+1 {
		// non-exists entry
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"[0]", "[1]", "[2]", "[]"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
		}
	}
}

func TestNode_setValueArrayInvalidType(t *testing.T) {
	v := [3]string{"a", "b", "c"}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	{
		node, err := root.setValue("[1]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueArrayInvalidType failed with data %#v", v)
		}
	}
	{
		node, err := root.setValue("[a]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueArrayInvalidType failed with data %#v", v)
		}
	}
}

func TestNode_setValueArray(t *testing.T) {
	v := genDataArray()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(&v), // for array: only addressable array is settable
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "xyz"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"[0]", "[1]", "[2]", "[3]"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
		}
	}
}

/*----------------------------------------------------------------------*/
func TestNode_createChildMap(t *testing.T) {
	{
		v := genDataArray()
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: only addressable array is settable
		}
		path := "[0]"
		node, err := root.createChildMap(path)
		if node == nil || err != nil || node.value.Elem().Kind() != reflect.Map || node.value.Elem().Len() != 0 {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
	{
		v := genDataSlice()
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "[0]"
		node, err := root.createChildMap(path)
		if node == nil || err != nil || node.value.Elem().Kind() != reflect.Map || node.value.Elem().Len() != 0 {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
	{
		v := genDataMap()
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "xyz"
		node, err := root.createChildMap(path)
		if node == nil || err != nil {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
	{
		type MyStruct struct {
			A       interface{}
			B       interface{}
			M       interface{}
			S       interface{}
			private interface{}
		}
		v := MyStruct{
			A: []int{0, 1, 2, 3, 4, 5},
			B: [3]string{"a", "b", "c"},
			M: map[string]interface{}{
				"x": "x",
				"y": 19.81,
				"z": true,
			},
			S:       Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
			private: 1.2,
		}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: only addressable struct is settable
		}
		path := "A"
		node, err := root.createChildMap(path)
		if node == nil || err != nil || node.value.Elem().Kind() != reflect.Map || node.value.Elem().Len() != 0 {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
}

func createChildSlice(t *testing.T) {
	{
		v := genDataArray()
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: only addressable array is settable
		}
		path := "[0]"
		node, err := root.createChildMap(path)
		if node == nil || err != nil {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
	{
		v := genDataSlice()
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "[0]"
		node, err := root.createChildMap(path)
		if node == nil || err != nil {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
	{
		v := genDataMap()
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "xyz"
		node, err := root.createChildMap(path)
		if node == nil || err != nil {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
	{
		type MyStruct struct {
			A       interface{}
			B       interface{}
			M       interface{}
			S       interface{}
			private interface{}
		}
		v := MyStruct{
			A: []int{0, 1, 2, 3, 4, 5},
			B: [3]string{"a", "b", "c"},
			M: map[string]interface{}{
				"x": "x",
				"y": 19.81,
				"z": true,
			},
			S:       Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
			private: 1.2,
		}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: only addressable struct is settable
		}
		path := "A"
		node, err := root.createChildMap(path)
		if node == nil || err != nil {
			t.Errorf("TestNode_createChildMap failed with data %#v at index {%#v}", v, path)
		}
	}
}
