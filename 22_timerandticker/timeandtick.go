package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer
	// create a timer to execute some process / function in the future
	timer1 := time.NewTimer(time.Second * 2)

	fmt.Println(<-timer1.C)
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println(<-timer2.C)
		fmt.Println("Timer 2 fired")
	}()

	// if stop is called before timer end, it will stop the timer
	// timer2Stop := timer2.Stop()
	// if timer2Stop {
	// 	fmt.Println("Timer 2 stopped and not fired")
	// }

	time.Sleep(time.Second * 2)

	// Ticker
	// Timer only run one time, ticker can run multiple time

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Ticker stopped")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop()
	done <- true
}
