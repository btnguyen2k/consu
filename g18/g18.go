package g18

import (
	"sort"
)

// Sortable is an interface that is implemented by all sortable types (string and numbers).
//
// Note: equivalent to built-in type cmp.Ordered introduced in Go 1.21.
type Sortable interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~uintptr
}

// Deduplicate removes duplicated elements from a slice.
//
// Note: the elements in the final slice may not be in the same order as those in the input. If you want to keep the order, use DeduplicateStable.
func Deduplicate[K Sortable](input []K) []K {
	if len(input) == 0 {
		return make([]K, 0)
	}
	result := make([]K, len(input))
	copy(result, input)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	prev := 0
	for i, n := 1, len(result); i < n; i++ {
		if result[i] != result[prev] {
			prev++
			result[prev] = result[i]
		}
	}
	return result[:prev+1]
}

// DeduplicateStable removes duplicated elements from a slice, preserving the order of the elements.
//
// @Available since v0.1.0
func DeduplicateStable[K comparable](input []K) []K {
	if len(input) == 0 {
		return make([]K, 0)
	}
	result := make([]K, 0)
	for _, v := range input {
		if FindInSlice(v, result) == -1 {
			result = append(result, v)
		}
	}
	return result
}

// FindInSlice returns the position of needle in haystack. -1 is return if not found.
func FindInSlice[K comparable](needle K, haystack []K) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

// PointerOf returns a "pointer" version of the input.
//
// @Available since v0.0.2
func PointerOf[K any](input K) *K {
	return &input
}

// Max returns the maximum value of the input value list.
//
// Note: if Go 1.21+, use built-in operator max() instead.
//
// @Available since v0.1.0
func Max[K Sortable](values ...K) K {
	if len(values) == 0 {
		panic("empty input")
	}
	result := values[0]
	for _, v := range values[1:] {
		if v > result {
			result = v
		}
	}
	return result
}

// Min returns the minimum value of the input value list.
//
// Note: if Go 1.21+, use built-in operator min() instead.
//
// @Available since v0.1.0
func Min[K Sortable](values ...K) K {
	if len(values) == 0 {
		panic("empty input")
	}
	result := values[0]
	for _, v := range values[1:] {
		if v < result {
			result = v
		}
	}
	return result
}
