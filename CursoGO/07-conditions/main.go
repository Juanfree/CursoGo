package main

import "fmt"

func main()  {
	println(".............................")
	ifExample()
	println(".............................")
	switchExample()
	println(".............................")
}

func ifExample()  {
	a := numbers()

	b, c, d := a[0], a[1], a[2]

	println(ifTest(b))
	println(ifTest(c))
	println(ifTest(d))
}

func ifTest(number int) bool {
	return number % 2 == 0
}

func numbers() [3]int {
	n := [3]int{1, 2, 3}
	return n
}

func switchExample()  {
	switchTest(1)
	switchTest(2)
	switchTest(3)
	switchTest(10)
}

func switchTest(a int)  {
	switch a {
	case 1:
		fmt.Println("Is 1")
	case 2:
		fmt.Println("Is 2")
	case 3, 4:
		fmt.Println("Is 3 or 4")
		fallthrough
	case 5:
		fmt.Println("Is 5")
	default:
		fmt.Println("Is greater than 4")
	}
}