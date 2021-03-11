package main

import "fmt"

func main() {
	i := 50

	if i%2 == 0 {
		fmt.Println("Even")
	} else { // else should be inline with {}
		fmt.Println("Odd")
	}

	j := "value"

	if i >= 69 {
		fmt.Println("number")
	} else if j == "value" { // also else if
		fmt.Println("val")
	} else {
		fmt.Println("else")
	}
}
