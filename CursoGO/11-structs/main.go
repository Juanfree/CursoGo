package main

import "fmt"

// Class...
type Example struct {
	a int
	b string
	c []int
}

func (e Example) sub(value int) {
	e.a -= value
}

func (e* Example) add(value int) {
	e.a += value
}

func main()  {
	//firstExample()

	a := new(Example)

	fmt.Println(a)

	a.add(10)

	fmt.Println(a)

	a.sub(10)

	fmt.Println(a)
}

func firstExample() {
	a := Example{a: 1, b: "hola", c: []int{1, 2, 3, 4, 5}}
	//Explicit way
	var b = new(Example)
	b.a = 10
	b.b = "dewww"
	b.c = []int{1, 2}
	c := new(Example)
	c.a = 10
	c.b = "dewww"
	c.c = []int{1}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
