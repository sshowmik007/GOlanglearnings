// Take the code from the previous exercise,
// then store the values of type person in a map
//  with the key of last name. Access each value
// in the map. Print out the values, ranging over the slice.

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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range v.favFlavour {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
