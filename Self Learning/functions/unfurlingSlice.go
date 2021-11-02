// When you have a slice of some type,
// you can pass in the individual values in
// a slice by using the “...” operator.

package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9} //unfurling a slice of int
	x := sum(xi...)
	fmt.Println("The total is", x)
}

// func main() {
// 	// passing in zero or more values
// 	x := sum(xi...)
// 	fmt.Println("The total is", x)
// }

//Unlimited number of values of type int
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
