//using composite literal crate an arrary which holds 5 values of TYPE in
//assaign values to aeach index

package main

import "fmt"

func main() {

	x := [5]int{45, 50, 68, 94, 20}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x) //type

}
