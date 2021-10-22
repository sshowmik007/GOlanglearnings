package main

import "fmt"

func main() {
	x := []int{5, 6, 9, 2, 8, 12, 456, 5}
	fmt.Println(x)
	// What append does is append the elements to
	//the end of the slice and return the result. The result needs to be returned because, as with our hand-written Append, the underlying array may change.
	// func append(slice []T, elements ...T) []T
	x = append(x, 77, 78, 50, 10506, 10)
	fmt.Println(x)

	y := []int{10, 62, 930, 52, 48, 212, 56, 50}
	x = append(x, y...) //here ... 3 dot means rest numbers from y
	fmt.Println(x)
}
