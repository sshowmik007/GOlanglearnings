package main

import "encoding/json"

type Person struct {
	Name  string
	Id    int
	Phone int
}

func main() {
	//annonymous struct
	p := struct {
		Name  string
		Id    int
		Phone int
	}{
		// if  we write the name we can ommit the values else we have to give all the values

		//we can ommit

		Name:  "Sadman",
		Id:    18101124,
		Phone: 454545,
	}
	//or
	json.Unmarshal(&p)
	//defer

	// fmt.Println(("first line"))
	// defer fmt.Println("defer")
	// fmt.Println("lastline")
}
