package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not found")
	}

	currentDay := time.Now().Weekday()
	switch currentDay {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I know, it's a boolean")
		case int:
			fmt.Println("I know, it's a number")
		default:
			fmt.Printf("What is %T\n", t)

		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("Test")
	whatAmI(2.22)
}
