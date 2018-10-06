package structs

type Human struct {
	name string
	age int
	gender string
}

func (h* Human) SetName(name string) {
	h.name = name
}

type Woman struct {
	Human
}

func (w Woman) GetName() string {
	return "Woman: " +  w.name
}

type Man struct {
	Human
}

func (m Man) GetName() string {
	return "Man: " + m.name
}
