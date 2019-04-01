package main

import (
	"fmt"
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
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	t := time.Now().In(loc)
	s := t.String()
	fmt.Println(s)
	toTime("2006-01-02 15:04:05.000000 -0700 -07", s)
	expected := time.Date(2019, 04, 29, 20, 59, 10, 0, time.UTC)
	fmt.Println(expected)

	// data := map[string]interface{}{
	// 	"ValueInt":     1547549353,
	// 	"ValueString":  "1547549353123",
	// 	"ValueInvalid": -1,
	// 	"ValueDateStr": "January 15, 2019 20:49:13.123",
	// }
	// s := semita.NewSemita(data)
	//
	// fmt.Println(s.GetTime("ValueInt"))
	// fmt.Println(s.GetTime("ValueString"))
	// fmt.Println(s.GetTime("ValueInvalid"))
	//
	// fmt.Println("==================================================")
	//
	// fmt.Println(s.GetTimeWithLayout("ValueInt", ""))
	// fmt.Println(s.GetTimeWithLayout("ValueString", ""))
	// fmt.Println(s.GetTimeWithLayout("ValueInvalid", ""))
	// fmt.Println(s.GetTimeWithLayout("ValueDateStr", "January 02, 2006 15:04:05.000"))
}
