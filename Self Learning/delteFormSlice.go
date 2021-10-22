package main

import "fmt"

func main() {
	x := []int{5, 6, 9, 2, 8, 12, 456, 5}
	fmt.Println(x)

	x = append(x, 77, 78, 50, 10506, 10)
	fmt.Println(x)

	y := []int{10, 62, 930, 52, 48, 212, 56, 50}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	//func append(slice []Type, elems ...Type) []Type removing indx 2-4
	fmt.Println(x)
}
