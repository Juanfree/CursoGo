package main

import (
	"fmt"
	"time"
)

func main() {
	go pingpong()
	time.Sleep(10 * time.Second)
}

func pingpong()  {
	ball := make(chan int)
	action := make(chan string)

	go referee(action)
	go ping(ball, action)

	for {
		value := <- ball
		switch value {
		case 1:
			go pong(ball, action)
		case 2:
			go ping(ball, action)
		}
	}
}

func ping(ball chan<- int, action chan<- string) {
	ball <- 1
	action <- "Player ping playing"
}

func pong (ball chan<- int, action chan<- string)  {
	ball   <- 2
	action <- "Player pong playing"
}

func referee(action <-chan string)  {
	for {
		fmt.Println("Action:" + <-action)
	}
}
