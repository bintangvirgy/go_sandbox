package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
		time.Sleep(time.Second)
	}
}

// go routine is one of golang most powerful feature. It create lightweight thread of execution
func main() {
	// this is a blockin function, it will not process next line if this function isn't done processed
	f("non-goroutine")

	// this is go routine, it will start thread and continue to next line
	go f("goroutine1")
	time.Sleep(time.Second)
	go f("goroutine2")

	// if we dont make a wait group, after go routine started it will end the main program
	time.Sleep(time.Second * 10)
}
