// We can take one struct and embed it in another struct. When you do this the inner type gets promoted to the outer type

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
			age:   21,
		},
		ltk: true,
	}

	p2 := person{
		first: "sadman",
		last:  "showmik",
		age:   25,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}
