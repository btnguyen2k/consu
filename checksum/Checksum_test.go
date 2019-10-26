package checksum

import (
	"crypto/md5"
	"fmt"
	"testing"
)

var hfList = []HashFunc{Crc32HashFunc, Md5HashFunc, Sha1HashFunc, Sha256HashFunc, Sha512HashFunc}
var csfList = []func(interface{}) []byte{Crc32Checksum, Md5Checksum, Sha1Checksum, Sha256Checksum, Sha512Checksum}

func TestChecksum_Bool(t *testing.T) {
	name := "TestChecksum_Bool"
	v1 := true
	v2 := false
	v3 := true
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, &v3))
		if checksum1 != checksum3 || !(checksum1 != checksum2) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}

func TestChecksum_Int(t *testing.T) {
	name := "TestChecksum_Int"
	v1 := int(103)
	v2 := int32(103)
	v3 := int64(103)
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, &v3))
		if checksum1 != checksum2 || checksum1 != checksum3 {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
		}

		v4 := int(301)
		checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
		if !(checksum1 != checksum4) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v4}, []interface{}{checksum1, checksum4})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}

func TestChecksum_Uint(t *testing.T) {
	name := "TestChecksum_Uint"
	v1 := uint(103)
	v2 := uint32(103)
	v3 := uint64(103)
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, &v3))
		if checksum1 != checksum2 || checksum1 != checksum3 {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
		}

		v4 := uint(301)
		checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
		if !(checksum1 != checksum4) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v4}, []interface{}{checksum1, checksum4})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}

func TestChecksum_Float(t *testing.T) {
	name := "TestChecksum_Float"
	v1 := float32(103)
	v2 := float32(103.0)
	v3 := float64(103)
	v4 := float64(103.0)
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
		checksum4 := fmt.Sprintf("%x", Checksum(hf, &v4))
		if checksum1 != checksum2 || checksum1 != checksum3 || checksum1 != checksum4 {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3, v4}, []interface{}{checksum1, checksum2, checksum3, checksum4})
		}

		vi := int(103)
		vui := uint(103)
		checksumi := fmt.Sprintf("%x", Checksum(hf, vi))
		checksumui := fmt.Sprintf("%x", Checksum(hf, vui))
		if !(checksum1 != checksumi) || !(checksum1 != checksumui) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, vi, vui}, []interface{}{checksum1, checksumi, checksumui})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}

func TestChecksum_String(t *testing.T) {
	name := "TestChecksum_String"
	v1 := "a"
	v2 := "a "
	v3 := " a"
	v4 := "A"
	v5 := "a"
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
		checksum4 := fmt.Sprintf("%x", Checksum(hf, &v4))
		if checksum1 == checksum2 || checksum1 == checksum3 || checksum1 == checksum4 || checksum2 == checksum3 || checksum2 == checksum4 || checksum3 == checksum4 {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3, v4}, []interface{}{checksum1, checksum2, checksum3, checksum4})
		}

		checksum5 := fmt.Sprintf("%x", Checksum(hf, &v5))
		if checksum1 != checksum5 {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v5}, []interface{}{checksum1, checksum5})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}

	s := "1"
	i := 1
	for _, hf := range hfList {
		checksumS := fmt.Sprintf("%x", Checksum(hf, s))
		checksumI := fmt.Sprintf("%x", Checksum(hf, i))
		if !(checksumS != checksumI) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{s, i}, []interface{}{checksumS, checksumI})
		}
	}
}

func TestChecksum_SliceArray(t *testing.T) {
	name := "TestChecksum_SliceArray"
	v1 := []int{1, 2}
	v2 := [2]uint{1, 2}
	v3 := []interface{}{int(1), uint(2)}
	v4 := [2]int{2, 1}
	v5 := []float64{1, 2}
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
		if checksum1 != checksum2 || checksum1 != checksum3 {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
		}

		checksum4 := fmt.Sprintf("%x", Checksum(hf, &v4))
		if !(checksum1 != checksum4) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v4}, []interface{}{checksum1, checksum4})
		}

		checksum5 := fmt.Sprintf("%x", Checksum(hf, &v5))
		if !(checksum1 != checksum5) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v5}, []interface{}{checksum1, checksum5})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}

	v0 := make([]interface{}, 0)
	for _, hf := range hfList {
		checksum := Checksum(hf, v0)
		if checksum == nil || len(checksum) != 0 {
			t.Fatalf("%s failed for input %#v - received %#v", name, v0, checksum)
		}
	}
}

func TestChecksum_Map(t *testing.T) {
	name := "TestChecksum_Map"
	v1 := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true}
	v2 := map[string]interface{}{"b": 2.3, "d": true, "c": "a string", "a": 1}
	v3 := map[string]interface{}{"x": 1, "y": 2.3, "z": "a string", "t": true}
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
		if !(checksum1 == checksum2) || !(checksum1 != checksum3) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}

type MyStructAllPublic struct {
	S string
	N int
	F float64
	A []interface{}
	M map[string]interface{}
}

func TestChecksum_StructAllPublic(t *testing.T) {
	name := "TestChecksum_StructAllPublic"
	a := []interface{}{1, 2.3, true, "a string"}
	m := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true}

	v1 := MyStructAllPublic{S: "string", N: 1, F: 2.3, A: a, M: m}
	v2 := MyStructAllPublic{N: 1, A: a, M: m, F: 2.3, S: "string"}
	v3 := MyStructAllPublic{S: "string", N: 1, F: 2.3}
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
		if !(checksum1 == checksum2) || !(checksum1 != checksum3) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}

type MyStructPubPriv struct {
	S string
	N int
	F float64
	a []interface{}
	m map[string]interface{}
}

func TestChecksum_StructPubPriv(t *testing.T) {
	name := "TestChecksum_StructPubPriv"
	a := []interface{}{1, 2.3, true, "a string"}
	m := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true}

	v1 := MyStructPubPriv{S: "string", N: 1, F: 2.3, a: a, m: m}
	v2 := MyStructPubPriv{N: 3, a: a, m: m, F: 1.2, S: "string"}
	v3 := MyStructPubPriv{S: "string", N: 1, F: 2.3, a: a, m: m}
	v4 := MyStructPubPriv{N: 1, F: 2.3, S: "string"}
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
		checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
		checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
		if !(checksum1 != checksum2) || !(checksum1 == checksum3) || !(checksum1 != checksum4) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2, v3, v4}, []interface{}{checksum1, checksum2, checksum3, checksum4})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}

type MyStruct struct {
	S string
	N uint
	F float64
	s string
	n int
	f float32
}

func (s MyStruct) Checksum() interface{} {
	h := md5.New()
	h.Write([]byte(s.S))
	h.Write(uintToBytes(uint64(s.N)))
	h.Write(floatToBytes(s.F))
	result := h.Sum(nil)
	return result
}

func TestChecksum_StructChecksum(t *testing.T) {
	name := "TestChecksum_StructChecksum"

	v1 := &MyStruct{S: "string", N: 1, F: 2.3, s: "STRING", n: -1, f: -2.3}
	v2 := MyStruct{S: "string", N: 1, F: 2.3, s: "String", n: 2, f: 4.6}
	for i, hf := range hfList {
		checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
		checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
		if !(checksum1 == checksum2) {
			t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v2}, []interface{}{checksum1, checksum2})
		}

		csf := csfList[i]
		if checksum1 != fmt.Sprintf("%x", csf(v1)) {
			t.Fatalf("%s failed at index %d", name, i)
		}
	}
}
