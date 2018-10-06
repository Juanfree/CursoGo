package main

import "fmt"

func main()  {
	arr1 := getArray()

	fmt.Println(arr1)
	fmt.Println(getCoercionArray() [2])
	fmt.Println(getSlice() [0])
}

func getArray() [2]string {
	var arr1 [2] string
	arr1[0] = "value 1"
	arr1[1] = "value 2"
	return arr1
}

func getCoercionArray() [3]int {
	arr := [3]int{1, 2, 3}
	return arr
}

func getSlice() []string {
	var slice [] string

	slice = append(slice, "hola", "que", "tal", "estas")

	return slice
}

//func getWrongArray() [2]string {
//	var arr1 [2] string
//	arr1[0] = "value 1"
//	arr1[1] = "value 2"
//	arr1[2] = "value 2"
//	return arr1
//}
