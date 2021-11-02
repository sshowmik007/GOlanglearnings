package main

import "fmt"

//we can made type of func
type fn func(int, int)int{



// type Resolver interface{
// 	GResolver()  //function's signature

// 
}

func main() {
	Arithmatic(5, 5, "sum", Sum)
}
func Sum(a int, b int) int {
	return a + b
}
func Arithmatic(a int, b int, opType string, f fn) { //callback
	var res int
	if opType == "sum" {
		res := f(a, b)
	}

	fmt.Println(res)
}
