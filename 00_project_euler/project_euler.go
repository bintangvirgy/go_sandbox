package main

import (
	"fmt"
	"math"
)

// this is code for solving project euler in projecteuler.net

func main() {

	// problem 1
	// var vsum int
	// vsum = 0
	// prob1(&vsum)
	// fmt.Println(vsum)

	// problem 2
	// vsumevent := 0
	// prob2(&vsumevent)
	// fmt.Println(vsumevent)

	// problem 3
	// vnum := 600851475143
	// vres := prob3(&vnum)
	// fmt.Println(vres)

	// problem 4
	vpalindrome := prob4()
	fmt.Println(vpalindrome)

}

func prob4() int {
	res := 0

	return res

}

func prob3(num *int) int {
	cur := *num
	res := 0
	for i := 2; i < int(math.Sqrt(float64(*num))); i++ {

		for cur%i == 0 {
			cur /= i
			res = i
			if cur == 1 {
				break
			}
		}

		if cur == 1 {
			break
		}
	}

	return res

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
