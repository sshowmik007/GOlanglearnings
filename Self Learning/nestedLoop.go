package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("outer loop : %d \n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("outer loop : %d\t  Inner Loop: %d \n", i, j)
		}
	}

}
