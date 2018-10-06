package main

import (
	"errors"
	"fmt"
)

func main() {
	b, _ := add(10,10)

	fmt.Println(b)

	c, err := add("10",10)

	fmt.Println(c, err)

	if (err != nil) {
		panic(err)
	}
}

func add(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("Param should be an integer")
	}


	return a.(int) + b.(int), nil
}
