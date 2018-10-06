package main

import "fmt"

func main()  {
	println(getName())
	println(getLastName())
	a, b, c := numbers()
	fmt.Printf("The numbers are: %d %d %d ", a, b, c)

	float32, float64 := floats()

	fmt.Println(float32, float64)
}

func scan() string{
	var value string
	fmt.Scanf("%s", &value)
	return value
}

func getName() string {
	fmt.Println("Enter your name.go:")
	return scan()
}

func getLastName() string {
	fmt.Println("Enter your lastname:")
	return scan()
}

func numbers() (int, int, int) {
	return 1,2,3
}

func suma(a int, b int)  int{
	return a + b
}

func floats() (float32, float64) {
	return float32(0.1), float64(0.1)
}
