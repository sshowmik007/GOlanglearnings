package main

import (
	"errors"
	"fmt"
)

func main() {

	// var condition = true

	// if condition {
	// 	fmt.Println("if block")
	// } else {
	// 	fmt.Println("else block")
	// }
	// //nested
	// if !condition {
	// 	if condition {
	// 		fmt.Println("if block")
	// 	} else {
	// 		fmt.Println("else block")
	// 	}
	// }
	//x:=5 function assign

	var input string
	fmt.Println("insert a name")
	fmt.Scanln(&input)

	if userExists, length, err := validateUser("sadman"); err == nil {
		fmt.Printf("user exists : %v \n", userExists)
		fmt.Printf("length of the user name %v ", length)
	} else {
		fmt.Println(err)
	}
}

func validateUser(name string) (bool, int, error) {
	const user = "sadman"
	// var userExists bool
	if name == user {
		return true, len(name), nil
	} else {
		return false, 0, errors.New("user doesn't exist")
	}
}
