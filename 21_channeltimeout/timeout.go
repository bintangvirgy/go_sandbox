package main

import (
	"fmt"
	"time"
)

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

func main() {
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
}
