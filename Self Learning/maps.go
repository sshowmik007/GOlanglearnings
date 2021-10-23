package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"james":      32,
		"moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["james"])

	//asking comma ok idiom
	//checking v,ok :=
	//if
	v, ok := m["sakib"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["james"]; ok {
		fmt.Println("this is the if print", v)
	}
}
