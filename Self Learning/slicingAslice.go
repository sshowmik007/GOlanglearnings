package main

import "fmt"

func main() {
	x := []int{5, 6, 9, 2, 8, 12, 456, 5}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}
	for j := 0; j <= 4; j++ {
		fmt.Println(j, x[j])
	}

}
