package checksum

import (
	"crypto/md5"
	"fmt"
	"testing"
	"time"
)

var nameList = []string{"CRC32", "MD5", "SHA1", "SHA256", "SHA512"}
var hfList = []HashFunc{Crc32HashFunc, Md5HashFunc, Sha1HashFunc, Sha256HashFunc, Sha512HashFunc}
var csfList = []func(interface{}) []byte{Crc32Checksum, Md5Checksum, Sha1Checksum, Sha256Checksum, Sha512Checksum}

func TestChecksum_Bool(t *testing.T) {
	testName := "TestChecksum_Bool"
	v1 := true
	v2 := false
	v3 := true
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, &v3))
			if checksum1 != checksum3 || !(checksum1 != checksum2) {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

func TestChecksum_Int(t *testing.T) {
	testName := "TestChecksum_Int"
	v1 := int(103)
	v2 := int32(103)
	v3 := int64(103)
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, &v3))
			if checksum1 != checksum2 || checksum1 != checksum3 || checksum2 != checksum3 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
			}

			v4 := int(301)
			checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
			if !(checksum1 != checksum4) {
				t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v4}, []interface{}{checksum1, checksum4})
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

func TestChecksum_Uint(t *testing.T) {
	testName := "TestChecksum_Uint"
	v1 := uint(103)
	v2 := uint32(103)
	v3 := uint64(103)
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, &v3))
			if checksum1 != checksum2 || checksum1 != checksum3 || checksum2 != checksum3 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v2, v3}, []interface{}{checksum1, checksum2, checksum3})
			}

			v4 := uint(301)
			checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
			if !(checksum1 != checksum4) {
				t.Fatalf("%s failed for input %#v - received %#v", name, []interface{}{v1, v4}, []interface{}{checksum1, checksum4})
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

func TestChecksum_Float(t *testing.T) {
	testName := "TestChecksum_Float"
	v1 := float32(103)
	v2 := float32(103.0)
	v3 := float64(103)
	v4 := float64(103.0)
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			checksum4 := fmt.Sprintf("%x", Checksum(hf, &v4))
			if checksum1 != checksum2 || checksum1 != checksum3 || checksum1 != checksum4 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v2, v3, v4}, []interface{}{checksum1, checksum2, checksum3, checksum4})
			}

			vi := int(103)
			vui := uint(103)
			checksumi := fmt.Sprintf("%x", Checksum(hf, vi))
			checksumui := fmt.Sprintf("%x", Checksum(hf, vui))
			if !(checksum1 != checksumi) || !(checksum1 != checksumui) {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, vi, vui}, []interface{}{checksum1, checksumi, checksumui})
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

func TestChecksum_String(t *testing.T) {
	testName := "TestChecksum_String"
	v1 := "a"
	v2 := "a "
	v3 := " a"
	v4 := "A"
	v5 := "a"
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			checksum4 := fmt.Sprintf("%x", Checksum(hf, &v4))
			if checksum1 == checksum2 || checksum1 == checksum3 || checksum1 == checksum4 || checksum2 == checksum3 || checksum2 == checksum4 || checksum3 == checksum4 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v2, v3, v4}, []interface{}{checksum1, checksum2, checksum3, checksum4})
			}

			checksum5 := fmt.Sprintf("%x", Checksum(hf, &v5))
			if checksum1 != checksum5 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v5}, []interface{}{checksum1, checksum5})
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

func TestChecksum_StringVsNumber(t *testing.T) {
	testName := "TestChecksum_StringVsNumber"
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			s := "1"
			n := 1
			ui := 1
			f32 := float32(1)
			f64 := float64(1)

			checksumS := fmt.Sprintf("%x", Checksum(hf, s))

			checksumN := fmt.Sprintf("%x", Checksum(hf, n))
			if !(checksumS != checksumN) {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{s, n}, []interface{}{checksumS, checksumN})
			}

			checksumUI := fmt.Sprintf("%x", Checksum(hf, ui))
			if !(checksumS != checksumUI) {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{s, ui}, []interface{}{checksumS, checksumUI})
			}

			checksumF32 := fmt.Sprintf("%x", Checksum(hf, f32))
			if !(checksumS != checksumF32) {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{s, f32}, []interface{}{checksumS, checksumF32})
			}

			checksumF64 := fmt.Sprintf("%x", Checksum(hf, f64))
			if !(checksumS != checksumF64) {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{s, f64}, []interface{}{checksumS, checksumF64})
			}
		})
	}
}

func TestChecksum_Time(t *testing.T) {
	testName := "TestChecksum_Time"
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	timeLayout := "2006-01-02T15:04:05.999999999-07:00"
	now := time.Now()
	v1 := now
	v2 := now.Add(1 * time.Hour)
	v3 := now.Add(-1 * time.Minute)
	v4 := now.Add(2 * time.Second)
	v5 := now.Add(-2 * time.Millisecond)
	v6 := now.Add(3 * time.Microsecond)
	v7 := now.Add(-3 * time.Nanosecond)
	v0, _ := time.Parse(timeLayout, now.Format(timeLayout))
	v0 = v0.In(loc)
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			checksum4 := fmt.Sprintf("%x", Checksum(hf, &v4))
			checksum5 := fmt.Sprintf("%x", Checksum(hf, v5))
			checksum6 := fmt.Sprintf("%x", Checksum(hf, &v6))
			checksum7 := fmt.Sprintf("%x", Checksum(hf, v7))
			if checksum1 == checksum2 || checksum1 == checksum3 || checksum1 == checksum4 || checksum1 == checksum5 || checksum1 == checksum6 || checksum1 == checksum7 ||
				checksum2 == checksum3 || checksum2 == checksum4 || checksum2 == checksum5 || checksum2 == checksum6 || checksum2 == checksum7 ||
				checksum3 == checksum4 || checksum3 == checksum5 || checksum3 == checksum6 || checksum3 == checksum7 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v2, v3, v4, v5, v6, v7}, []interface{}{checksum1, checksum2, checksum3, checksum4, checksum5, checksum6, checksum7})
			}

			checksum0 := fmt.Sprintf("%x", Checksum(hf, &v0))
			if checksum1 != checksum0 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v0}, []interface{}{checksum1, checksum0})
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
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

func TestChecksum_SliceArrayWithTime(t *testing.T) {
	name := "TestChecksum_SliceArrayWithTime"
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	v1 := []interface{}{int(1), int32(2), int64(3), now}
	v2 := []interface{}{uint(1), uint32(2), uint64(3), now.UTC()}
	v3 := []interface{}{int(1), uint(2), int16(3), now.In(loc)}
	v4 := [4]interface{}{now, int64(3), int32(2), int(1)}
	v5 := []interface{}{float32(1), float64(2), float32(3), now}
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
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	nowInLoc := now.In(loc)
	v1 := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true, "time": now, "timep": &nowInLoc}
	v2 := map[string]interface{}{"b": 2.3, "d": true, "c": "a string", "a": 1, "time": nowInLoc, "timep": &now}
	v3 := map[string]interface{}{"x": 1, "y": 2.3, "z": "a string", "t": true, "time": now, "timep": &now}
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
	S  string
	N  int
	F  float64
	A  []interface{}
	M  map[string]interface{}
	T  time.Time
	TP *time.Time
}

func TestChecksum_StructAllPublic(t *testing.T) {
	name := "TestChecksum_StructAllPublic"
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	nowInLoc := now.In(loc)
	a := []interface{}{1, 2.3, true, "a string", now, &nowInLoc}
	m := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true, "t": now, "tp": &nowInLoc}

	v1 := MyStructAllPublic{S: "string", N: 1, F: 2.3, A: a, M: m, T: now, TP: &nowInLoc}
	v2 := MyStructAllPublic{N: 1, A: a, M: m, F: 2.3, S: "string", T: nowInLoc, TP: &now}
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
	S  string
	N  int
	F  float64
	a  []interface{}
	m  map[string]interface{}
	t  time.Time
	TP *time.Time
}

func TestChecksum_StructPubPriv(t *testing.T) {
	name := "TestChecksum_StructPubPriv"
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	nowInLoc := now.In(loc)
	a := []interface{}{1, 2.3, true, "a string", now, &nowInLoc}
	m := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true, "t": now, "tp": &nowInLoc}

	v1 := MyStructPubPriv{S: "string", N: 1, F: 2.3, a: a, m: m, t: now, TP: &nowInLoc}
	v2 := MyStructPubPriv{N: 3, a: a, m: m, F: 1.2, S: "string", t: now, TP: &nowInLoc}
	v3 := MyStructPubPriv{S: "string", N: 1, F: 2.3, a: a, m: m, t: nowInLoc, TP: &now}
	v4 := MyStructPubPriv{N: 1, F: 2.3, S: "string", t: nowInLoc, TP: &now}
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
