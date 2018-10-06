package main

import (
	"fmt"
	"strings"
)

const STRING = "Hola que tal estas? Hola"

func main()  {
	fmt.Println(replaceOne(upperCase(STRING)))
	fmt.Println(replaceAll(upperCase(STRING)))
	fmt.Println(split(STRING))
	fmt.Println(len(split(STRING)))
}

func upperCase(value string) string{
	return strings.ToUpper(value)
}

func replaceOne(value string) string {
	return strings.Replace(value, "HOLA", "Hello", 1)
}

func replaceAll(value string) string {
	return strings.Replace(value, "HOLA", "Hello", -1)
}

func split(value string) []string  {
	return strings.Split(value, " ")
}
