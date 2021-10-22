package main

func main() {
	//declaring array

	// var arr1 [5] int //declaring an array var name [length] datatype
	// arr1[4] = 45

	//another way

	/*
		var arr1 [5]int = [5]int{1, 2, 3, 4, 5} //declaring+ initialization
		var arr2 [5]int = [5]int{1, 2, 3}       //declaring+ initialization
		var names [3]string = [3]string{"a", "b", "c"}
		fmt.Println(len(arr1))
		fmt.Println(arr2)
		//slicing of an array
		fmt.Println(arr1[1:3])
		fmt.Println(names)
	*/

	/*
		//declaring array new system
		var arr2 = new([5]string) //reference type
		arr2[4] = "sadman"
		fmt.Println(arr2)

		var arr_new = arr2
		arr_new[4] = "new data"
		fmt.Println(arr2)
	*/

	var a [5]*int

	a[1] = nil

}
