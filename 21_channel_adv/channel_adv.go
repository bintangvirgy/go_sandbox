package main

import (
	"fmt"
	"time"
)

// close channel
// used to close the channel, if receive channel to 2 var, second param will return channel state, either it's still open or closed
func closeChan(jobs <-chan int, done chan<- bool) {
	for {
		j, more := <-jobs
		if more {
			fmt.Println("Receive job ", j)
		} else {
			fmt.Println("All job received")
			done <- true
			return
		}
	}
}

// range over channel
// close channel function can be used to range the channel
func chanRange(rangeV chan string) {
	rangeV <- "one"
	rangeV <- "two"
	rangeV <- "three"
	// if the channel not closed, the range operation will error after channel emptied
	close(rangeV)

	for v := range rangeV {
		fmt.Println(v)
	}
}

func main() {

	fmt.Println("Closing Channel")
	jobs := make(chan int, 5)
	done := make(chan bool)

	go closeChan(jobs, done)

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("Sending job", i)
		time.Sleep(time.Microsecond * 500)
	}
	close(jobs)
	fmt.Println("All job sent")
	<-done

	fmt.Println()
	fmt.Println("Range over channel")
	rangeV := make(chan string, 5)
	chanRange(rangeV)
}
