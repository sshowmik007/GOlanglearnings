// that's going to print out a slice for you have a slice with 10 values in it a slice of ants with
// 10 values in it.

package main

import "fmt"

func main() {

	x := []int{45, 50, 68, 94, 20, 45, 46, 4, 50, 65} //slice
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)

}
