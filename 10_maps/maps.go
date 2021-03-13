package main

import "fmt"

func main() {
	// maps is like dictionary in python or associative array in PHP

	// map created by map[key-datatype]value-datatype
	m := make(map[string]int)
	fmt.Println(m)
	m["key1"] = 13
	m["key2"] = 50
	fmt.Println(m)

	var1 := m["key2"]
	fmt.Println(var1)

	// if key not found, first will return default value of datatype, and second variable will return boolean is key exist
	var2, key2 := m["key5"]
	fmt.Println(var2)
	fmt.Println(key2)

	// we can use blank identifier ( _ ) to ignore variable
	_, isKeyExist := m["key1"]
	fmt.Println(isKeyExist)

	delete(m, "key2")
	delete(m, "key5")
	fmt.Println(m)

}
