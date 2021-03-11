package main

import (
	"fmt"
	"math"
)

const s string = "constant s"

func main() {
	fmt.Println(s)

	const n float64 = 50000

	const d float64 = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
