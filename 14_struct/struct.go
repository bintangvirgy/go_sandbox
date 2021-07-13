package main

import "fmt"

//People : create struct type, type which hold multiple data
type People struct {
	name string
	age  int
}

func newPeople(name string) *People {
	p := People{
		name: name,
		age:  35,
	}
	return &p
}

func combinePeople(pp1, pp2 *People) {
	pp1.age = pp1.age + pp2.age
	pp1.name = pp1.name + pp2.name
}

func manipulatePeople(pp *People) {
	pp.age = pp.age + 15
}

func manipulatenonPeople(pp People) {
	pp.age = pp.age + 15
}

func main() {
	bob := People{"Bob", 42}
	fmt.Println("Bob", bob)

	ann := People{age: 35, name: "Ann"}
	fmt.Println("Ann", ann)
	manipulatenonPeople(ann)
	fmt.Println("Ann", ann)

	nameonly := People{name: "Namehere"}
	fmt.Println("Nameonly", nameonly)

	ageonly := People{age: 32}
	fmt.Println("Age only", ageonly)

	pointerst := &People{
		name: "Pointo",
		age:  33,
	}

	fmt.Println("Pointer struct", pointerst)
	// pointer of struct can auto dereferenced
	fmt.Println("Pointer struct age", pointerst.age)

	initPeople := newPeople("Johnny")
	fmt.Println("Init use pointer", initPeople)
	fmt.Println("Init use pointer", initPeople.age)
	manipulatePeople(initPeople)
	fmt.Println("Init use pointer", initPeople.age)
	manipulatePeople(initPeople)
	fmt.Println("Init use pointer", initPeople.age)

	combinePeople(initPeople, &bob)
	fmt.Println("Combine", initPeople)

	combinePeople(&bob, initPeople)
	fmt.Println("Combine", bob)

}
