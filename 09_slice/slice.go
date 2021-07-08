package main

import "fmt"

func main() {
	// create slice (like flexible array)
	s := make([]string, 3)

	// manual insertion, can't add value on index more than length
	fmt.Println(s)
	s[0] = "a"
	s[2] = "b"
	s[1] = "c"
	fmt.Println(s)
	fmt.Println(len(s))

	// to add new value more than length, use append
	s = append(s, "d", "e", "g")
	fmt.Println(s)
	fmt.Println(len(s))

	c := make([]string, len(s))
	// copy created slice to new slice
	copy(c, s)
	fmt.Println(c)
	fmt.Println(len(c))

	ap := make([]string, 0)
	// other way to copy, use spread operator (...)
	ap = append(ap, s...)
	ap = append(ap, c...)
	fmt.Println(ap)
	fmt.Println(len(ap))

	// create new variable based on created slice
	// slice operator just like slice in python
	newap := ap[2:5]
	fmt.Println(newap)
	// change value on new variable
	newap[0] = "ssnew"
	fmt.Println(newap)
	// it will change the original value
	fmt.Println(ap)

	// like initialize array but without length
	t := []string{"test", "123", "new", "slice"}

	fmt.Println(t)

	num := []int{1, 3, 2, 4, 2}

	fmt.Println(num)

	num = append(num, 2, 1, 3, 2, 4)
	fmt.Println(num)
}
