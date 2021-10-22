package main

import (
	"fmt"
)

// we cant use  ":=" outside of the scope

//importing multiple library
// we cant keep unused import in our code
func main() {

	// we cant keep unused variable in our code
	// by using var we can declare variable
	var a int = 45
	var b string = "hello sadman"
	// type inference
	var c = "how are you??"
	x, y := 10, 20

	var (
		g string = "bello"
	)

	// another type
	// variable_name := value
	d := 65
	fmt.Println(a, b, c)
	fmt.Println(d)
	fmt.Println(x, y)
	fmt.Println(g)

}
