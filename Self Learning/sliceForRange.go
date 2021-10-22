package main

import "fmt"

func main() {

	x := []int{5, 6, 9, 2, 8, 12, 456, 5}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[5])
	fmt.Println(x[3])
	fmt.Println(x[1])
	//If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
	/*
	   for key, value := range collection {
	       //block of statements
	   } */
	for i, v := range x {
		fmt.Println(i, v)
	}

}
