package main

import "fmt"

//it is annonymous bcz it doesnt have a name for struct
// We can create anonymous structs also. An anonymous struct is a struct which is not associated with a specific identifier.

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   21,
	}
	fmt.Println(p1)

}
