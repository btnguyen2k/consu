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

			if checksum1 == checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum3)
			}

			if checksum1 != checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

func TestChecksum_BoolVsNumber(t *testing.T) {
	testName := "TestChecksum_BoolVsNumber"
	vArr1 := []interface{}{true, int(1), int8(1), int16(1), int32(1), int64(1), uint(1), uint8(1), uint16(1), uint32(1), uint64(1), float32(1), float64(1), []byte{1}, byte(1)}
	vArr2 := []interface{}{false, int(0), int8(0), int16(0), int32(0), int64(0), uint(0), uint8(0), uint16(0), uint32(0), uint64(0), float32(0), float64(0), []byte{0}, byte(0)}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, vArr1[0]))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, vArr2[0]))
			for j := 1; j < len(vArr1); j++ {
				checksum3 := fmt.Sprintf("%x", Checksum(hf, vArr1[j]))
				if checksum1 == checksum3 {
					t.Fatalf("%s failed: Checksum(%#v)=%s must be NOT the same as Checksum(%#v)=%s", testName+"/"+name, vArr1[0], checksum1, vArr1[j], checksum3)
				}

				checksum4 := fmt.Sprintf("%x", Checksum(hf, vArr2[j]))
				if checksum2 == checksum4 {
					t.Fatalf("%s failed: Checksum(%#v)=%s must be NOT the same as Checksum(%#v)=%s", testName+"/"+name, vArr2[0], checksum2, vArr2[j], checksum4)
				}
			}
		})
	}
}

func TestChecksum_Int(t *testing.T) {
	testName := "TestChecksum_Int"
	vArr := []interface{}{int(103), int8(103), int16(103), int32(103), int64(103)}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				if j%2 == 0 {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, &v))
				} else {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, v))
				}
			}
			for j := 0; j < len(checksumArr)-1; j++ {
				for k := j + 1; k < len(checksumArr); k++ {
					if checksumArr[j] != checksumArr[k] {
						t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[j], checksumArr[j], vArr[k], checksumArr[k])
					}
				}
			}

			v0 := int(301)
			checksum0 := fmt.Sprintf("%x", Checksum(hf, v0))
			if checksumArr[0] == checksum0 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be NOT the same as Checksum(%#v)=%s", testName+"/"+name, vArr[0], checksumArr[0], v0, checksum0)
			}

			v := fmt.Sprintf("%x", csfList[i](vArr[0]))
			if v != checksumArr[0] {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksumArr[0], v)
			}
		})
	}
}

func TestChecksum_Uint(t *testing.T) {
	testName := "TestChecksum_Uint"
	vArr := []interface{}{uint(103), uint8(103), uint16(103), uint32(103), uint64(103)}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				if j%2 == 0 {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, &v))
				} else {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, v))
				}
			}
			for j := 0; j < len(checksumArr)-1; j++ {
				for k := j + 1; k < len(checksumArr); k++ {
					if checksumArr[j] != checksumArr[k] {
						t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[j], checksumArr[j], vArr[k], checksumArr[k])
					}
				}
			}

			v0 := uint(301)
			checksum0 := fmt.Sprintf("%x", Checksum(hf, v0))
			if checksumArr[0] != checksum0 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be NOT the same as Checksum(%#v)=%s", testName+"/"+name, vArr[0], checksumArr[0], v0, checksum0)
			}

			v := fmt.Sprintf("%x", csfList[i](vArr[0]))
			if v != checksumArr[0] {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksumArr[0], v)
			}
		})
	}
}

func TestChecksum_IntVsUInt(t *testing.T) {
	testName := "TestChecksum_IntVsUInt"
	viArr := []interface{}{int(103), int8(103), int16(103), int32(103), int64(103)}
	vuiArr := []interface{}{uint(103), uint8(103), uint16(103), uint32(103), uint64(103)}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumiArr := make([]string, len(viArr))
			checksumuiArr := make([]string, len(vuiArr))
			for j := range viArr {
				if j%2 == 0 {
					checksumiArr[j] = fmt.Sprintf("%x", Checksum(hf, viArr[j]))
					checksumuiArr[j] = fmt.Sprintf("%x", Checksum(hf, &vuiArr[j]))
				} else {
					checksumiArr[j] = fmt.Sprintf("%x", Checksum(hf, &viArr[j]))
					checksumuiArr[j] = fmt.Sprintf("%x", Checksum(hf, vuiArr[j]))
				}
			}
			for j, ci := range checksumiArr {
				for k, cui := range checksumuiArr {
					if ci != cui {
						t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, viArr[j], checksumiArr[j], vuiArr[k], checksumuiArr[k])
					}
				}
			}
		})
	}
}

func TestChecksum_Float(t *testing.T) {
	testName := "TestChecksum_Float"
	vArr := []interface{}{float32(103), float32(103.0), float64(103), float64(103.0)}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				if j%2 == 0 {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, &v))
				} else {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, v))
				}
			}
			for j := 0; j < len(checksumArr)-1; j++ {
				for k := j + 1; k < len(checksumArr); k++ {
					if checksumArr[j] != checksumArr[k] {
						t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[j], checksumArr[j], vArr[k], checksumArr[k])
					}
				}
			}

			v0 := uint(301)
			checksum0 := fmt.Sprintf("%x", Checksum(hf, v0))
			if checksumArr[0] != checksum0 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[0], checksumArr[0], v0, checksum0)
			}

			v := fmt.Sprintf("%x", csfList[i](vArr[0]))
			if v != checksumArr[0] {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksumArr[0], v)
			}
		})
	}
}

func TestChecksum_FloatVsIntUint(t *testing.T) {
	testName := "TestChecksum_FloatVsIntUint"
	v32 := float32(103)
	v64 := float64(103)
	vi := int(103)
	vui := uint(103)
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumf32 := fmt.Sprintf("%x", Checksum(hf, v32))
			checksumf64 := fmt.Sprintf("%x", Checksum(hf, v64))
			checksumi := fmt.Sprintf("%x", Checksum(hf, vi))
			checksumui := fmt.Sprintf("%x", Checksum(hf, vui))

			if checksumf32 == checksumi {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v32, checksumf32, vi, checksumi)
			}
			if checksumf32 == checksumui {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v32, checksumf32, vui, checksumui)
			}

			if checksumf64 == checksumi {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v64, checksumf64, vi, checksumi)
			}
			if checksumf64 == checksumui {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v64, checksumf64, vui, checksumui)
			}
		})
	}
}

func TestChecksum_String(t *testing.T) {
	testName := "TestChecksum_String"
	vArr := []string{"a", "a ", " a", "A", "A ", " A", "aA"}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				if j%2 == 0 {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, &v))
				} else {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, v))
				}
			}
			for j := 0; j < len(checksumArr)-1; j++ {
				for k := j + 1; k < len(checksumArr); k++ {
					if checksumArr[j] == checksumArr[k] {
						t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[j], checksumArr[j], vArr[k], checksumArr[k])
					}
				}
			}

			for _, v := range vArr {
				c := fmt.Sprintf("%x", Checksum(hf, v))
				cp := fmt.Sprintf("%x", Checksum(hf, &v))
				if c != cp {
					t.Fatalf("%s failed: Checksum(%#v)=%s must sbe the same as Checksum(%#v)=%s", testName+"/"+name, v, c, &v, cp)
				}
			}

			v := fmt.Sprintf("%x", csfList[i](vArr[0]))
			if v != checksumArr[0] {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksumArr[0], v)
			}
		})
	}
}

func TestChecksum_StringVsNumber(t *testing.T) {
	testName := "TestChecksum_StringVsNumber"
	vArr := []interface{}{"1", int(1), uint(1), float32(1), float64(1)}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, v))
			}
			for j := 1; j < len(checksumArr); j++ {
				if checksumArr[0] == checksumArr[j] {
					t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[0], checksumArr[0], vArr[j], checksumArr[j])
				}
			}
		})
	}
}

func TestChecksum_TimeZone(t *testing.T) {
	testName := "TestChecksum_TimeZone"
	now := time.Now()
	zones := []string{
		"Australia/Adelaide",
		"Australia/Brisbane",
		"Australia/Canberra",
		"Australia/Darwin",
		"Australia/Melbourne",
		"Australia/Perth",
		"Australia/Sydney",
		"Australia/Tasmania",
		"Asia/Ho_Chi_Minh",
	}
	timeLayout := "2006-01-02T15:04:05.999999999-07:00"
	vArr := []time.Time{now, now.UTC()}
	for _, zone := range zones {
		loc, _ := time.LoadLocation(zone)
		vArr = append(vArr, now.In(loc))
		v0, _ := time.Parse(timeLayout, now.Format(timeLayout))
		vArr = append(vArr, v0)
		vArr = append(vArr, v0.In(loc))
		vArr = append(vArr, v0.UTC())
	}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				if j%2 == 0 {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, &v))
				} else {
					checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, v))
				}
			}
			for j := 0; j < len(checksumArr)-1; j++ {
				for k := j + 1; k < len(checksumArr); k++ {
					if checksumArr[j] != checksumArr[k] {
						t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[j], checksumArr[j], vArr[k], checksumArr[k])
					}
				}
			}
		})
	}
}

func TestChecksum_Time(t *testing.T) {
	testName := "TestChecksum_Time"
	now := time.Now()
	vArr := []time.Time{
		now,
		now.Add(1 * time.Hour),
		now.Add(-1 * time.Minute),
		now.Add(2 * time.Second),
		now.Add(-2 * time.Millisecond),
		now.Add(3 * time.Microsecond),
		now.Add(-3 * time.Nanosecond),
	}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, &v))
			}
			for j := 0; j < len(checksumArr)-1; j++ {
				for k := j + 1; k < len(checksumArr); k++ {
					if checksumArr[j] == checksumArr[k] {
						t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[j], checksumArr[j], vArr[k], checksumArr[k])
					}
				}
			}

			v := fmt.Sprintf("%x", csfList[i](vArr[0]))
			if v != checksumArr[0] {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksumArr[0], v)
			}
		})
	}
}

func TestChecksum_SliceArray(t *testing.T) {
	testName := "TestChecksum_SliceArray"
	v1 := []int{1, 2}
	v2 := [2]uint{1, 2}
	v3 := []uint{2, 1}
	v4 := [3]int{1, 2, 0}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
			if checksum1 != checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum2)
			}
			if checksum1 == checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}
			if checksum1 == checksum4 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v4, checksum4)
			}
		})
	}
}

func TestChecksum_SliceArrayEmpty(t *testing.T) {
	testName := "TestChecksum_SliceArrayEmpty"
	vArr := []interface{}{[]int{}, [0]uint{}, []interface{}{}, [0]interface{}{}, []string{}, [0]string{}, []bool{}, [0]bool{}}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			for _, v := range vArr {
				checksum1 := Checksum(hf, v)
				if checksum1 == nil || len(checksum1) != 0 {
					t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, v, checksum1)
				}
				checksum2 := Checksum(hf, &v)
				if checksum2 == nil || len(checksum2) != 0 {
					t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, v, checksum2)
				}
			}
		})
	}
}

func TestChecksum_SliceVsArray(t *testing.T) {
	testName := "TestChecksum_SliceVsArray"
	vArr := []interface{}{
		[]int{1, 2},
		[2]uint{1, 2},
		[]interface{}{int(1), uint(2)},
		[2]interface{}{uint(1), int(2)},
	}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksumArr := make([]string, len(vArr))
			for j, v := range vArr {
				checksumArr[j] = fmt.Sprintf("%x", Checksum(hf, v))
			}
			for j := 0; j < len(checksumArr)-1; j++ {
				for k := j + 1; k < len(checksumArr); k++ {
					if checksumArr[j] != checksumArr[k] {
						t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, vArr[j], checksumArr[j], vArr[k], checksumArr[k])
					}
				}
			}
		})
	}
}

func TestChecksum_SliceArrayWithTime(t *testing.T) {
	testName := "TestChecksum_SliceArrayWithTime"
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	v1 := []interface{}{int(1), int32(2), int64(3), now}
	v2 := []interface{}{uint(1), uint32(2), uint64(3), now.UTC()}
	v3 := []interface{}{int(1), uint(2), int16(3), now.In(loc)}
	v4 := [4]interface{}{now, int64(3), int32(2), int(1)}
	v5 := []interface{}{float32(1), float64(2), float32(3), now}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			if checksum1 != checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum2)
			}
			if checksum1 != checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}

			checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
			if checksum1 == checksum4 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v4, checksum4)
			}

			checksum5 := fmt.Sprintf("%x", Checksum(hf, v5))
			if checksum1 == checksum5 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v5, checksum5)
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

func TestChecksum_Map(t *testing.T) {
	testName := "TestChecksum_Map"
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	nowInLoc := now.In(loc)
	nowUTC := now.UTC()
	v1 := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true, "time": now, "timep": &nowInLoc}
	v2 := map[string]interface{}{"b": 2.3, "d": true, "c": "a string", "a": 1, "time": &nowInLoc, "timep": now}
	v3 := map[string]interface{}{"x": 1, "y": 2.3, "z": "a string", "t": true, "time": &now, "timep": &nowUTC}
	v4 := map[string]interface{}{"A": 1, "B": 2.3, "C": "a string", "D": true, "time": now, "timep": &nowInLoc}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
			if checksum1 != checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum2)
			}

			if checksum1 == checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}
			if checksum1 == checksum4 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v4, checksum4)
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
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
	testName := "TestChecksum_StructAllPublic"
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	nowInLoc := now.In(loc)
	a := []interface{}{1, 2.3, true, "a string", now, &nowInLoc}
	m := map[string]interface{}{"a": 1, "b": 2.3, "c": "a string", "d": true, "t": now, "tp": &nowInLoc}

	v1 := MyStructAllPublic{S: "string", N: 1, F: 2.3, A: a, M: m, T: now, TP: &nowInLoc}
	v2 := MyStructAllPublic{N: 1, A: a, M: m, F: 2.3, S: "string", T: nowInLoc, TP: &now}
	v3 := MyStructAllPublic{S: "string", N: 1, F: 2.3}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			if checksum1 != checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum2)
			}
			if checksum1 == checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

type MyStructPubPriv struct {
	S string
	N int
	F float64
	s string
	n int
	f float64
}

func TestChecksum_StructPubPriv(t *testing.T) {
	testName := "TestChecksum_StructPubPriv"
	v1 := MyStructPubPriv{S: "string", N: 1, F: 2.3, s: "a string", n: 2, f: 3.4}
	v2 := MyStructPubPriv{s: "a string", n: 2, f: 3.4, S: "string", N: 1, F: 2.3}
	v3 := MyStructPubPriv{S: "string", N: 1, F: 2.3}
	v4 := MyStructPubPriv{s: "a string", n: 2, f: 3.4}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
			if checksum1 != checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum2)
			}

			if checksum1 == checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}
			if checksum1 == checksum4 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v4, checksum4)
			}

			if checksum3 == checksum4 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v3, checksum3, v4, checksum4)
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

type MyStructPubPrivPointer struct {
	S *string
	N *int
	F *float64
	s *string
	n *int
	f *float64
}

func TestChecksum_StructPubPrivPointer(t *testing.T) {
	testName := "TestChecksum_StructPubPrivPointer"
	var s string = "string"
	var n int = 1
	var f float64 = 2.3
	var s1 string = "a string"
	var n1 int = -1
	var f1 float64 = -3.4
	v1 := MyStructPubPrivPointer{S: &s, N: &n, F: &f, s: &s1, n: &n1, f: &f1}
	v2 := MyStructPubPrivPointer{s: &s1, n: &n1, f: &f1, S: &s, N: &n, F: &f}
	v3 := MyStructPubPrivPointer{S: &s, N: &n, F: &f}
	v4 := MyStructPubPrivPointer{s: &s, n: &n, f: &f}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, &v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			checksum4 := fmt.Sprintf("%x", Checksum(hf, v4))
			if checksum1 != checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum2)
			}

			if checksum1 == checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}
			if checksum1 == checksum4 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v4, checksum4)
			}

			if checksum3 == checksum4 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v3, checksum3, v4, checksum4)
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
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
	testName := "TestChecksum_StructChecksum"

	v1 := MyStruct{S: "string", N: 1, F: 2.3, s: "STRING", n: -1, f: -2.3}
	v2 := &MyStruct{S: "string", N: 1, F: 2.3, s: "String", n: 1, f: 2.3}
	v3 := MyStruct{S: "String", N: 1, F: 2.3, s: "STRING", n: -1, f: -2.3}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
			checksum3 := fmt.Sprintf("%x", Checksum(hf, v3))
			if checksum1 != checksum2 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v2, checksum2)
			}
			if checksum1 == checksum3 {
				t.Fatalf("%s failed: Checksum(%#v)=%s must NOT be the same as Checksum(%#v)=%s", testName+"/"+name, v1, checksum1, v3, checksum3)
			}

			v := fmt.Sprintf("%x", csfList[i](v1))
			if v != checksum1 {
				t.Fatalf("%s failed, expected %#v but received %#v", testName+"/"+name, checksum1, v)
			}
		})
	}
}

type MyStruct1 struct {
	S string
}

type MyStruct2 struct {
	S string
}

func TestChecksum_TwoStructs(t *testing.T) {
	testName := "TestChecksum_TwoStructs"
	v1 := MyStruct1{S: "a string"}
	v2 := MyStruct2{S: "a string"}
	for i, name := range nameList {
		t.Run(name, func(t *testing.T) {
			hf := hfList[i]
			checksum1 := fmt.Sprintf("%x", Checksum(hf, v1))
			checksum2 := fmt.Sprintf("%x", Checksum(hf, v2))
			if checksum1 == checksum2 {
				t.Fatalf("%s failed for input %#v - received %#v", testName+"/"+name, []interface{}{v1, v2}, []interface{}{checksum1, checksum2})
			}
		})
	}
}
