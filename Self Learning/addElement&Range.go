package main

import (
	"fmt"
)

func main() {
	xi := []int{4, 5, 6, 9, 8, 2, 35}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
