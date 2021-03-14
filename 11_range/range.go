package main

import "fmt"

func main() {
	rSlices := []int{1, 4, 2, 3, 6, 2}
	sum := 0
	fmt.Println(rSlices)

	// in this case, we don't need key
	for _, val := range rSlices {
		sum += val
	}
	fmt.Println(sum)

	for key, val := range rSlices {
		if key%2 == 0 {
			fmt.Println("Even key :", val)
		} else {
			fmt.Println("Odd key :", val)
		}
	}

	fruits := map[string]string{"a": "Apple", "b": "Banana", "c": "Cherry"}

	for key, val := range fruits {
		fmt.Println("Key :", key, "=> Val:", val)
	}

	// also work with unicode
	for _, v := range "testing" {
		fmt.Println(v)
	}
}
