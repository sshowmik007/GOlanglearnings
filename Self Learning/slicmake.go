package main

import "fmt"

func main() {
	//make([]T, length, capacity)
	//make([]int, 50, 100)
	//new([100]int)[0:50]
	x := make([]int, 10, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3566)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 35)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3556)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3999)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	//capacity increases 2x
}
