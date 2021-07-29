package main

import "fmt"

// this is code for solving project euler in projecteuler.net

func main() {

	// problem 1
	var vsum int
	vsum = 0
	prob1(&vsum)
	fmt.Println(vsum)

	// problem 2
	// vsumevent := 0

	// fmt.Println(vsumevent)

	// problem 3

}

func prob2(vsumevent *int) {
	vold := 1
	vnew := 2
	vnext := 0

	for vnew <= 4000000 {
		vnext = vold + vnew
		if vnew%2 == 0 {
			*vsumevent += vnew
		}
		vold = vnew
		vnew = vnext
	}
}

func prob1(vsum *int) {
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			*vsum += i
		}
	}
}
