package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 250
	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println(len(a))

	b := [5]int{1, 4, 2, 3, 12}

	fmt.Println(b)

	twoD := [2][3]int{{2, 3, 1}, {31, 2, 1}}

	fmt.Println(twoD[0])
}
