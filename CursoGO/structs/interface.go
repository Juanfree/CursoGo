package structs

import "fmt"

type HumanInterface interface {
	GetName() string
}

func GetName (h HumanInterface) string {
	return h.GetName()
}

func TestInterface(h HumanInterface) {
	fmt.Println(h.GetName())
}