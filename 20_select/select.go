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

// timeout
// sometime we use external resources and want to implement timeout if the external resources not responding for some periode of time

func slowFunc(res chan<- string) {
	time.Sleep(time.Second * 3)
	res <- "result slow"
}

func fastFunc(res chan<- string) {
	time.Sleep(time.Second * 1)
	res <- "result fast"
}

// nonblocking channel
// basic usage of channel are blocking. If we want to make channel to be non blocking we can use default in select case
func procNonBlock(messages chan<- string) {
	messages <- "Hello World"
}

func main() {
	//select
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

	// timeout
	// this code will timeout because func > 2 sec
	resultSlow := make(chan string, 1)
	resultFast := make(chan string, 1)

	go slowFunc(resultSlow)

	select {
	case msg := <-resultSlow:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Slow Timeout")
	}

	go fastFunc(resultFast)

	select {
	case msg := <-resultFast:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Fast Timeout")
	}

	// non blocking channel

	nonBlock := make(chan string, 1)

	// this will make the channel non blocking if it not received any value
	select {
	case msg := <-nonBlock:
		fmt.Println("Received msg: ", msg)
	default:
		fmt.Println("Message not received")
	}

	// because this will blocking and return an error
	// msg2 := <-nonBlock
	// fmt.Println(msg2)

	// but if we process the channel
	procNonBlock(nonBlock)

	// this function will return a value
	select {
	case msg := <-nonBlock:
		fmt.Println("Received msg: ", msg)
	default:
		fmt.Println("Message not received")
	}

}
