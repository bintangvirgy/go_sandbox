package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// same multiple datatype variable
func plusPlus(a, b, c int) int {
	return a + b + c
}

// multiple return
func plusMinus(a, b, c, d int) (int, int) {
	return a + b + c + d, a - b - c - d
}

//variadic function, use spread variable ( ... )
func sumAll(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

//closure function, function return another function
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	res := plus(3, 6)
	fmt.Println(res)

	res1 := plusPlus(3, 5, 213)
	fmt.Println(res1)

	res21, res22 := plusMinus(1, 65, 34, 43)
	fmt.Println(res21, res22)

	res21, _ = plusMinus(12, 13, 24, 33)
	fmt.Println(res21)

	ressum1 := sumAll(1, 2, 3, 2, 1, 2, 3)
	fmt.Println(ressum1)

	ressum2 := sumAll(1, 3)
	fmt.Println(ressum2)

	numMap := []int{3, 3, 3, 3, 3}
	ressum3 := sumAll(numMap...)
	fmt.Println(ressum3)

	// state new closure function
	nextInt := intSeq()

	fmt.Println(nextInt)
	//call closure function
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// different closure function
	nextInts := intSeq()

	fmt.Println(nextInts())

	// factorial
	fmt.Println(factorial(7))
}
