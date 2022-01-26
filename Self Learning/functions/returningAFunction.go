

//we can also return a func() {}

package main

import (
"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()

	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

}

func foo() string {
	return "Hello world"
}

func bar() func() int {
	return func() int {   //TYPE
		return 451
	}
}
