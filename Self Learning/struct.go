// A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types). Structs allow us to compose together values of different types.

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   21,
	}

	p2 := person{
		first: "sadman",
		last:  "showmik",
		age:   25,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
