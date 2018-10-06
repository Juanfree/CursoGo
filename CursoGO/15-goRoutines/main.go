package main

import (
	"fmt"
	"time"
)

func main() {
	goRoutineWithSleep()
}

func goRoutineWithSleep() {
	for i := 1; i < 1000; {
		go hello(i)
		i++
	}
	// 5 seconds ..
	time.Sleep(5 * time.Second)
}

func hello(index int)  {
	fmt.Println("Hello index! " , index)
}
