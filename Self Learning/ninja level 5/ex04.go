// Create and use an anonymous struct.

package main

import "fmt"

func main() {
	s := struct {
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first: "sadman",
		friends: map[string]int{
			"sadman": 555,
			"Q":      666,
			"R":      999,
		},
		favDrink: []string{
			"COCACOLA",
			"Pepso",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrink)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	for i, val := range s.favDrink {
		fmt.Println(i, val)
	}
}
