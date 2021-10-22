package main

import "fmt"

func main() {
	/*

		switch {
		case false:
			fmt.Println("this should not print")
		case (2 == 4):
			fmt.Println("wont print")
		case (4 == 4):
			fmt.Println("print")
			fallthrough //check next line
		case (8 > 4):
			fmt.Println("print this one ")
			// if we use fallthrough always it will repetadly print the nxt lines
			//if nothing is true we use defult:
		default:
			fmt.Println("nothing was true")
		}
	*/

	// we can assign the value to check
	/*
		n := "bond"
		switch n {
		case "moneyPenny":
			fmt.Println("miss money")
		case "bond":
			fmt.Println("james bond")
		case "Q":
			fmt.Println("The QQQ")
		default:
			fmt.Println("this is defult")
		}
	*/

	// we can use multiple check cases
	n := "bond"
	switch n {
	case "moneyPenny", "bond", "hello":
		fmt.Println("miss money or bond or hellow")
	case "x":
		fmt.Println("X")
	case "Q":
		fmt.Println("The QQQ")
	default:
		fmt.Println("this is defult")
	}

}
