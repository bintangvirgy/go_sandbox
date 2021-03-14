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

func main() {
	bob := People{"Bob", 42}
	fmt.Println(bob)

	ann := People{age: 35, name: "Ann"}
	fmt.Println(ann)

	nameonly := People{name: "Namehere"}
	fmt.Println(nameonly)

	ageonly := People{age: 32}
	fmt.Println(ageonly)

	pointerst := &People{
		name: "Pointo",
		age:  33,
	}

	fmt.Println(pointerst)
	// pointer of struct can auto dereferenced
	fmt.Println(pointerst.age)

	initPeople := newPeople("Johnny")
	fmt.Println(initPeople)

	combinePeople(initPeople, &bob)
	fmt.Println(initPeople)

	combinePeople(&bob, initPeople)
	fmt.Println(bob)

}
