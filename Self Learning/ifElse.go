package main

import "fmt"

func main() {
	x := 40
	if x == 40 {
		fmt.Println("value is 40")
	} else if x == 41 {
		fmt.Println("value is 41")
	} else if x == 42 {
		fmt.Println("value is 42")
	} else {
		fmt.Println("none of he above")
	}
}
