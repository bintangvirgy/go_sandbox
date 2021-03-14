package main

import "fmt"

// pointer will reference and dereference memoryaddress of a variable

// non pointer
func strNon(valStr string) {
	valStr = "Hello World non pointer"
}

// use pointer
func strPtr(valStr *string) {
	// valStr as memory address
	// *valStr to get value inside memory address
	*valStr = "Hello World pointer"
}

func main() {
	str := "Hello"
	fmt.Println(str)

	strNon(str)
	fmt.Println(str)

	// &str to pass memory address
	strPtr(&str)
	fmt.Println(str)

	// pointer address
	fmt.Println(&str)

}
