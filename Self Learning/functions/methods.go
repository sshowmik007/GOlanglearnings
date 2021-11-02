/*A method is nothing more than a FUNC attached to a
TYPE. When you attach a func to a type it is a method
of that type. You attach a func to a type with a RECEIVER.*/

package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(param) (return(s)) {CODES}

func (s secretAgent) speak() { //taking the argument and calling the method
	fmt.Println("I am ", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"sadman",
			"showmik",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
