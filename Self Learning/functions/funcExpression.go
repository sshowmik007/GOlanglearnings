package main

import "fmt"

func main (){
	fmt.Println("hello their")
	f:= func() {
		fmt.Println(" I am Showmik")
	}
	f()

	g:= func(x int) {
		fmt.Println(" I am Showmik", x)
	}
	g(1999)
}
