package main

import (
	"github.com/example/CursoGO/maps"
	"fmt"
)

func main()  {
	m := maps.GetMap()
	fmt.Println(len(m))
	fmt.Println(m)
	fmt.Println(m["value2"])


	delete(m, "value2")
	fmt.Println(len(m))
	fmt.Println(m)
}
