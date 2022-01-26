package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
Name string  //capital letter means export
Phone int
address string
}

func main() {
	p:= Person {"showmik" , 01303 , "mirpur"}
	data,_:= json.Marshal(p)
	fmt.Println(string(data))
}


//{
//	//p:=45
//	//const (
//	//	pp int = 4
//	//	k string = "hello"
//	//	l int32 = 4
//	//)
//	//	fmt.Println(p)
//
//}
