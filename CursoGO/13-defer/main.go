package main

import "fmt"

func main()  {
	defer testDefer3()
	defer testDefer1()
	defer testDefer2()
	hello()
	defer testDefer4()
}

func hello()  {
	fmt.Println("Hello!")
}

func testDefer1() {
	fmt.Println("Bye 1!")
}

func testDefer2() {
	fmt.Println("Bye 2!")
}

func testDefer3() {
	fmt.Println("Bye 3!")
}

func testDefer4() {
	fmt.Println("Bye 4!")
}