package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"} //[]string - slice of string
	fmt.Println(x)
	y := []string{"m", "n", "l", "o"}
	fmt.Println(y)

	multidimentionalSlice := [][]string{x, y}
	fmt.Println(multidimentionalSlice) //multi dimentional
}
