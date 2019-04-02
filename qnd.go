package main

import (
	"fmt"
	"github.com/btnguyen2k/consu/reddo"
	"reflect"
	"time"
)

func toTime(layout, source string) {
	t, e := time.Parse(layout, source)
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println(t.UnixNano(), t)
	}
}

func main() {
	s := "Nguyễn Bá Thành"
	barr := []byte(s)

	{
		v1, err := reddo.ToString(barr)
		fmt.Println(v1, err)
		v2, err := reddo.Convert(barr, reddo.TypeString)
		fmt.Println(v2, err)
	}

	{
		v1, err := reddo.ToSlice(s, reflect.TypeOf(barr))
		fmt.Println(v1, string(v1.([]byte)), err)
		v2, err := reddo.Convert(s, reflect.TypeOf(barr))
		fmt.Println(v2, string(v2.([]byte)), err)
	}
}
