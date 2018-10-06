package maps

func GetMap() map[string]int {
	m := make(map[string]int)

	m["value1"] = 1
	m["value2"] = 2
	m["value3"] = 3

	return m
}