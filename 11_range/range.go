package main

import "fmt"

func main() {
	rSlices := []int{1, 4, 2, 3, 6, 2}
	sum := 0
	fmt.Println(rSlices)

	for _, val := range rSlices {
		sum += val
	}
	fmt.Println(sum)
}
