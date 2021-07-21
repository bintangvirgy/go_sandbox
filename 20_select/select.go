package main

import (
	"fmt"
	"time"
)

// select
// waiting all channel to be available to receive
// use case statement, it will block until one of channel is available to run

func wait1(val1 chan string) {
	time.Sleep(time.Second)
	val1 <- "val1"
}

func wait2(val2 chan string) {
	time.Sleep(time.Second * 2)
	val2 <- "val2"
}

func main() {
	val1 := make(chan string, 1)
	val2 := make(chan string, 1)

	go wait1(val1)
	go wait2(val2)

	for i := 0; i < 2; i++ {
		// select will block execution and wait until one of its channel available to receive
		// we need loop because select only run once for channel availability
		// if there're multiple channel available, it will choose randomly
		select {
		case msg2 := <-val2:
			fmt.Println("message: ", msg2)
		case msg1 := <-val1:
			fmt.Println("message: ", msg1)
		}
	}

}
