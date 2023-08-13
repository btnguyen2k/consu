package g18

import "sort"

// Sortable is an interface that is implemented by all sortable types (string and numbers).
type Sortable interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Deduplicate removes duplicated elements from a slice.
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
// @available since v0.0.2
func PointerOf[K any](input K) *K {
	return &input
}
