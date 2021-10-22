package main

import "fmt"

func main() {
	var a []string = []string{" sfa ", "sad"}

	ab := append(a, " name  ", " newname")
	fmt.Println(ab)

}
