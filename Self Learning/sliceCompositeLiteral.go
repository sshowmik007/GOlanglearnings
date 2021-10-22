package main

import "fmt"

func main() {
	//composite literals x:= type{values}
	// slice allows u to group together values of the same type
	x := []int{5, 6, 9, 2, 8, 12, 456, 5}
	fmt.Println(x)
}
