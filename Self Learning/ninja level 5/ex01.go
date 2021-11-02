// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavour []string
}

func main() {
	p1 := person{
		first:      "sadman",
		last:       "showmik",
		favFlavour: []string{"5", "6", "7"},
	}
	p2 := person{
		first:      "mardia",
		last:       "mehzabin",
		favFlavour: []string{"2", "3", "9"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavour {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavour {
		fmt.Println(i, v)
	}
}
