package main

import "fmt"

const TRUE =  true

func main()  {
	var name string
	lastname := "Testing"
	name = "Testing"


	var (
		a = 1
		b = 2
		c = 3
	)

	d := 6

	clauseLeft := lastname == name
	clauseRight := (a + b + c) == d

	fmt.Println(clauseLeft == clauseRight == TRUE)
}
