package main

import "fmt"

func main()  {
	for i := 0; i < 5 ; i++ {
		funcName(i)
	}

	i := 0
	for i < 5 {
		i++
		funcName(i)
	}

	i = 0
	for i < 10 {
		i++

		if i == 5 {
			break
		}

		funcName(i)
	}
}

func funcName(i int) {
	if i < 2 {
		fmt.Println("Less than 2")
	} else if i == 3 {
		fmt.Println("Equal to 3")
	} else {
		fmt.Println("Greater than 4")
	}
}
