package semita

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

var (
	genericSliceType = reflect.TypeOf([]interface{}{})
	genericMapType   = reflect.TypeOf(map[string]interface{}{})
)

// createEmptyGenericSlice creates an empty slice of type interface{}
func createEmptyGenericSlice() reflect.Value {
	return reflect.MakeSlice(genericSliceType, 0, 0)
}

// createEmptyGenericMap creates an empty map of type [string]interface{}
func createEmptyGenericMap() reflect.Value {
	return reflect.MakeMap(genericMapType)
}

// isExportedField returns true if 'fieldName' indicates a struct' exported field.
func isExportedField(fieldName string) bool {
	return len(fieldName) >= 0 && string(fieldName[0]) == strings.ToUpper(string(fieldName[0]))
}

/*----------------------------------------------------------------------*/

type node struct {
	// this_node = prev->[key]
	prev     *node
	prevType reflect.Type
	key      string
	value    reflect.Value
}

// unwrap returns the underlying 'value' as an interface
func (n *node) unwrap() interface{} {
	return n.value.Interface()
}

func (n *node) elem() reflect.Value {
	v := n.value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Interface {
		v = n.value.Elem()
	}
	return v
}

// next returns a child node located at 'index'
func (n *node) next(index string) (*node, error) {
	vNode := n.elem()
	if match := patternIndex.FindStringSubmatch(index); len(match) > 0 {
		// current node should be an array or slice
		if vNode.Kind() == reflect.Array || vNode.Kind() == reflect.Slice {
			var i = vNode.Len()
			if "[]" != index {
				var e error
				i, e = strconv.Atoi(match[1])
				if e != nil {
					return nil, e
				}
			}
			if i < 0 || i >= vNode.Len() {
				return nil, nil
			}
			return &node{
				prev:     n,
				prevType: vNode.Type(),
				key:      index,
				value:    vNode.Index(i),
			}, nil
		}
		return nil, errors.New("invalid type {" + vNode.Type().String() + "}")
	} else if vNode.Kind() == reflect.Map {
		// map's key must be string
		if vNode.Type().Key().Kind() != reflect.String {
			return nil, errors.New("node of type {" + vNode.Type().String() + "} is not supported, map type must be string")
		}
		v := vNode.MapIndex(reflect.ValueOf(index))
		if v.Kind() == reflect.Invalid {
			return nil, nil
		}
		return &node{
			prev:     n,
			prevType: vNode.Type(),
			key:      index,
			value:    v,
		}, nil
	} else if vNode.Kind() == reflect.Struct {
		f := vNode.FieldByName(index)
		if f.Kind() == reflect.Invalid {
			// non-exist field
			return nil, nil
		}
		if !isExportedField(index) {
			// handle unexported field
			rv := reflect.New(vNode.Type()).Elem()
			rv.Set(vNode)
			f = rv.FieldByName(index)
			f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
		}
		return &node{
			prev:     n,
			prevType: vNode.Type(),
			key:      index,
			value:    f,
		}, nil
	}
	return nil, errors.New("invalid type {" + vNode.Type().String() + "}")
}

// setValue inserts the value as a child into the correct position specified by 'index'.
// when successful, this function returns the newly created child node.
func (n *node) setValue(index string, value reflect.Value) (*node, error) {
	vNode := n.elem()
	if match := patternIndex.FindStringSubmatch(index); len(match) > 0 {
		if vNode.Kind() == reflect.Slice || vNode.Kind() == reflect.Array {
			var i = vNode.Len()
			var e error
			if "[]" != index {
				i, e = strconv.Atoi(match[1])
				if e != nil {
					return nil, e
				}
			}
			if i < 0 || i > vNode.Len() || (i == vNode.Len() && vNode.Kind() != reflect.Slice) {
				return nil, errors.New("index out of bound {" + index + "}")
			}
			if !value.Type().AssignableTo(vNode.Type().Elem()) {
				return nil, errors.New("value of type {" + value.Type().String() + "} is not assignable to element of {" + vNode.Type().String() + "}")
			}
			if i == vNode.Len() {
				// special case: append to tail of slice
				vNode = reflect.Append(vNode, value)
				if n.prev == nil {
					n.value = vNode
				} else {
					_n, _e := n.prev.setValue(n.key, vNode)
					if _n == nil || _e != nil {
						return nil, errors.New("error setting appended slice to previous node {" + n.key + "}")
					}
					return _n.next(fmt.Sprintf("[%d]", i))
				}
			} else {
				childNode := vNode.Index(i)
				if !childNode.CanSet() {
					// final check
					return nil, errors.New("entry at index {" + index + "} is not settable")
				}
				childNode.Set(value)
			}
			return n.next(fmt.Sprintf("[%d]", i))
		}
		return nil, errors.New("expecting array or slice, but it is {" + vNode.Type().String() + "}")
	}
	if vNode.Kind() == reflect.Map {
		if vNode.Type().Key().Kind() != reflect.String {
			// map's key must be string
			return nil, errors.New("node of type {" + vNode.Type().String() + "} is not supported, map type must be string")
		}
		if !value.Type().AssignableTo(vNode.Type().Elem()) {
			// map's element type must match
			return nil, errors.New("value of type {" + value.Type().String() + "} is not assignable to element of map {" + vNode.Type().String() + "}")
		}
		vNode.SetMapIndex(reflect.ValueOf(index), value)
		return n.next(index)
	} else if vNode.Kind() == reflect.Struct {
		f := vNode.FieldByName(index)
		if f.Kind() == reflect.Invalid || !isExportedField(index) {
			// field must exist and is exported
			return nil, errors.New("{" + vNode.Type().String() + "} does not has exported field {" + index + "}")
		}
		if !value.Type().AssignableTo(f.Type()) {
			// field type must match
			return nil, errors.New("value of type {" + value.Type().String() + "} is not assignable to field {" + f.Type().String() + "}")
		}
		if !f.CanSet() {
			// final check
			return nil, errors.New("field {" + index + "} is not settable")
		}
		f.Set(value)
		return n.next(index)
	}
	return nil, errors.New("expecting map or struct, but it is {" + vNode.Type().String() + "}")
}

// createChildMap creates an empty map and insert it as a child node
func (n *node) createChildMap(index string) (*node, error) {
	return n.setValue(index, createEmptyGenericMap())
}

// createChildSlice creates an empty slice and insert it as a child node
func (n *node) createChildSlice(index string) (*node, error) {
	return n.setValue(index, createEmptyGenericSlice())
}
