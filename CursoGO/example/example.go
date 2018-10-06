package example

import "fmt"

func Hola()  {
	privateFunction()
}

func privateFunction() {
	fmt.Println("private function")
}

func NotPrivateFunction() {
	fmt.Println("not private function")
}
