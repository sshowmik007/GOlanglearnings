//using SLICE crate an arrary which holds 10 values of TYPE in
//assaign values to aeach index

package main

import "fmt"

func main() {

	x := []int{45, 50, 68, 94, 20, 45, 46, 4, 50, 65} //slice
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x) //type

}
