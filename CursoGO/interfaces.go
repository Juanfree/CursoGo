package main

import (
	"github.com/example/CursoGO/structs"
)

func main()  {

	w1 := new(structs.Woman)
	w1.SetName("Maria")

	m1 := new(structs.Man)
	m1.SetName("Pepito")

	structs.TestInterface(w1)
	structs.TestInterface(m1)
}
