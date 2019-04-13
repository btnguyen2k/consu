package main

import (
	"fmt"
	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/consu/semita"
)

func main() {
	i := 123
	s, e := reddo.ToString(i)
	fmt.Println(s, e)

	m := make(map[string]interface{})
	m["i"] = 456
	s, e = reddo.ToString(m["i"])
	fmt.Println(s, e)
	s, e = reddo.ToString(m["s"])
	fmt.Println(s, e)

	sem := semita.NewSemita(m)
	s1, e := sem.GetValueOfType("i", reddo.TypeString)
	fmt.Println(s1, e)
	s1, e = sem.GetValueOfType("s", reddo.TypeString)
	fmt.Println(s1, e)
}
