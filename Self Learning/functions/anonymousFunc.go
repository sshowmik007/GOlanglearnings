package main

import (

	"fmt"
)
func main(){
	foo()

	func (){
		fmt.Println("um ana anonymous func")
	}()

	func (x int){
		fmt.Println("hy i am : " , x)
	}(23)

}

func foo() {
	fmt.Println("foo ran")
}