package semita

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

var (
	genericSliceType = reflect.TypeOf([]interface{}{})
	genericMapType   = reflect.TypeOf(map[string]interface{}{})
)

func createEmptyGenericSlice() reflect.Value {
	return reflect.MakeSlice(genericSliceType, 0, 0)
}

func createEmptyGenericMap() reflect.Value {
	return reflect.MakeMap(genericMapType)
}

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
	if n.value.Kind() == reflect.Invalid {
		return nil
	}
	return n.value.Interface()
}

// next returns a child node according to 'index'
func (n *node) next(index string) (*node, error) {
	if n.value.Kind() == reflect.Invalid {
		return nil, errors.New("current node is nil")
	}
	vNode := n.value
	if vNode.Kind() == reflect.Interface {
		vNode = n.value.Elem()
	}
	if match := patternIndex.FindStringSubmatch(index); len(match) > 0 {
		// current node should be an array or slice
		if vNode.Kind() == reflect.Array || vNode.Kind() == reflect.Slice {
			i, e := strconv.Atoi(match[1])
			if e != nil {
				return nil, e
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

// setValue inserts the value as a child into the correct position specified by 'index'
func (n *node) setValue(index string, value reflect.Value) (*node, error) {
	if n.value.Kind() == reflect.Invalid {
		return n, errors.New("current node is nil")
	}
	vNode := n.value
	if vNode.Kind() == reflect.Interface {
		vNode = n.value.Elem()
	}

	if match := patternIndex.FindStringSubmatch(index); len(match) > 0 || "[]" == index {
		if vNode.Kind() == reflect.Slice || vNode.Kind() == reflect.Array {
			var i int
			var e error
			if "[]" == index {
				i = vNode.Len()
			} else {
				i, e = strconv.Atoi(match[1])
				if e != nil {
					return n, e
				}
			}
			if i < 0 || i > vNode.Len() {
				return n, errors.New("index out of bound " + index)
			}
			if i == vNode.Len() && vNode.Kind() == reflect.Slice {
				// special case
				vNode = reflect.Append(vNode, value)
				n.prev.setValue(n.key, vNode)
			} else {
				vNode.Index(i).Set(value)
			}
			return n, nil
		}
		return n, errors.New("expecting current node is array, but it is {" + vNode.Type().String() + "}")
	} else {
		if vNode.Kind() == reflect.Map {
			vNode.SetMapIndex(reflect.ValueOf(index), value)
			return n, nil
		}
		return n, errors.New("expecting current node is map, but it is {" + vNode.Type().String() + "}")
	}
	return n, nil
}

// createChildMap creates an empty map and insert it as a child node
func (n *node) createChildMap(index string) (*node, error) {
	return n.setValue(index, createEmptyGenericMap())
}

// createChildSlice creates an empty slice and insert it as a child node
func (n *node) createChildSlice(index string) (*node, error) {
	return n.setValue(index, createEmptyGenericSlice())
}
