package main

import "fmt"

type Course struct {
	Subject    string
	Coursecode int
}

type Student struct { //class
	Name    string
	Id      int
	Courses []Course
}

/*
func (s Student) PrintDetails() { //method
	fmt.Printf("Name : %v \n", s.Name)
	fmt.Printf("ID : %v", s.Id)
}
*/
func (s *Student) PrintDetails() { //  * pointer
	s.Name = "Invalid"
	s.Id = 0
}

func main() {

	student := Student{ //instance
		Name: " student ",
		Id:   454545,
	}

	fmt.Println(student)
	student.PrintDetails()
	fmt.Println(student)
}
