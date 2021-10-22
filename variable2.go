package main

import (
	"fmt"
)

func main() {
	//	const name, age = "Kim", 22
	//	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.
	var (
		a int = 45
		b     = "hello"
	)

	const (
		ip     = "192.168.0.1"
		dbName = "she"
	)

	fmt.Println(a)
	fmt.Println(b)

}
