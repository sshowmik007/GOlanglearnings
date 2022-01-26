

//func (receiver) identifier(parameters) (returns) { code }
//know the difference between parameters and arguments
//		we define our func with parameters (if any)
//		we call our func and pass in arguments (in any)
//Everything in Go is PASS BY VALUE
//		purpose of functions
//		abstract code
//		code reusability

package main

import "fmt"

func main() {
	foo() //foo(arguments)
	bar("james")
	s1 := woo("moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "fleming")
	fmt.Println(x)
	fmt.Println(y)
}

//func (r receiver) identifier (parameter) (return(s)) {code..}

func foo() {
	fmt.Println("hy i am foo")
}

//everything in GO pass by VALUE
func bar(s string) {
	fmt.Println("hello , ", s)
}

//single return
func woo(st string) string {
	return fmt.Sprint("hello from woo ", st)
}

//mutiple return
// func mouse(fn,ln string)
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, ` says "Hello" `)
	b := false
	return a, b
}
