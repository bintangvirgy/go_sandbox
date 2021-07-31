package main

import (
	"fmt"
	"math"
	"strconv"
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
	// vpalindrome := prob4()
	// fmt.Println(vpalindrome)

	// problem 5
	vkpk := prob5()
	fmt.Println(vkpk)

}

func prob5() int {
	res := 0
	num := 20
	for i := 1; true; i++ {
		res = num * i
		for j := num; j > 0; j-- {
			if res%j != 0 {
				break
			}

			if j == 1 {
				return res
			}
		}
	}
	return res
}

func prob4() int {
	// largest, so we start from here
	first_num := 999
	res := 0

	// loop backward for first number
	for i := first_num; i > 0; i-- {
		// loop backward too for second number
		// but i make limit by hundred place from first number, so if we still calculate 9xx from first number, we can't make second number lower than 9xx
		for j := i; j/100 >= i/100; j-- {
			res = i * j
			// check palindrome use custom function, if palindrome then break
			if isPalindrome(&res) {
				break
			}
		}
		if isPalindrome(&res) {
			break
		}
	}
	return res

}

func isPalindrome(num *int) bool {
	str := strconv.Itoa(*num)
	len_str := len(str) - 1
	for i := 0; i < len(str); i++ {
		if i >= len_str {
			break
		}

		if string(str[i]) != string(str[len_str]) {
			return false
		}

		len_str--

	}
	return true

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
