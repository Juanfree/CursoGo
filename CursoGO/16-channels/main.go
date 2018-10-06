package main

import "fmt"

func main() {
	//firstExample()
	//secondExample()

	withBufferExample()
}

func firstExample() {
	ch := make(chan string)
	go func() {
		defer close(ch)
		ch <- "First channel"
	}()
	fmt.Println(<-ch)
}

func secondExample() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
	}()
	for i := range ch {
		fmt.Println(ch, i)
	}
}

func withBufferExample() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func withBufferAndErrorExample() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}