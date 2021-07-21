package main

// channel
// channel buffering
// channel sync
// channel direction
import (
	"fmt"
	"time"
)

//channel can make pipeline from one go routine to other goroutine
var messages = make(chan string)
var messages1 = make(chan string)
var messages2 = make(chan string)

//make buffered channel with max 2
var buffmessages = make(chan string, 2)

func ping() {
	time.Sleep(time.Second)
	messages <- "ping"
}

func doubleping() {
	messages1 <- "ping1"
	time.Sleep(time.Second)
	messages2 <- "ping2"
}

func bufferedping() {
	buffmessages <- "pingbuff1"
	buffmessages <- "pingbuff2"

	// but adding this still works
	// i use 2 buffer
	buffmessages <- "pingbuff4"
}

func goBlock(done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func sendChan(pings chan<- string, msg string) {
	// this func only send pings channel
	pings <- msg

	// this line will error because we can't receive channel value
	// <-pings

}

func recChan(pings <-chan string, pongs chan<- string) {
	// this func only receive channel
	msg := <-pings

	pongs <- msg

	// this line will error
	// pings <- "test"
}

func main() {
	go ping()
	go doubleping()
	go bufferedping()

	// this channel is blocking, it will wait value from go function
	msg := <-messages
	fmt.Print("blocking: ")
	fmt.Println(msg)

	// load buffered ping
	fmt.Println("Buffered messages")
	fmt.Println(<-buffmessages)
	fmt.Println(<-buffmessages)
	fmt.Println(<-buffmessages)

	// go doubleping()

	// this code-block will error if we run msg2 before msg1
	// msg2 := <-messages2
	// fmt.Println(msg2)

	// this code-block run well
	fmt.Println("multi channel: ")
	msg1 := <-messages1
	fmt.Println(msg1)
	msg2 := <-messages2
	fmt.Println(msg2)

	done := make(chan bool)
	// we can make channel be a pointer too
	// but it no use because channel auto use references, not create a new var
	go goBlock(done)

	// channel also can make goroutine run sync
	fmt.Println("channel syncing: ")
	fmt.Println(<-done)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go sendChan(pings, "Hellow")
	go recChan(pings, pongs)

	fmt.Println("Directional channel: ")
	fmt.Println(<-pongs)
}
