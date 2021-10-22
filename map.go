package main

import ("fmt" 
	encoding/json
)


func main() {

	// var map1 map[string]string

	// map1 = make(map[string]string)
	//interface can take anuthing
	// another way

	// var map1 map[string]string
	// map1 = map[string]string{"Key1": "value1", "Key2": "value2"}

	//interface can take anything

	var map1 map[string]interface{}
	map1 = make(map[string]interface{})

	data:= {
		"name" : "john" ,
		"age" : 29
		"hobbies" :[
			"martial arts" ,
			"breakfast foods" ,
			"piano"
		]
	}

	//err := json.Unmarshal([]byte(data) , &map1)   //convertion the data into array
	//ignoring error
	//err := json.Unmarshal([]byte(data) , &map1)   //convertion the data into array
		//to ignore something we would use " _ "
	
	/*
		if err!=nil{
		fmt.Println(err)
	}
	map1 = make(map[string]interface{})
	map1["income"]=2.2
	map1["id"]=18101124
	map1["firstname"] = "sadman"
	map1["lastname"] = "showmik"

	fmt.Println(map1)
	*/
	j.Unmarshal([]byte(data),&map1)
	fmt.Println(map1)

}
