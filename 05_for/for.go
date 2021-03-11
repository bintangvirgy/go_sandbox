package main

import "fmt"

func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	count := 20
	for j := 0; j < count; j++ {
		if j <= 5 {
			continue
		}
		fmt.Println(j)
		if j >= 10 {
			break
		}
	}
}
