package main

import (
	"fmt"
	"math"
)

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// to add all struct to interface, the struct must implement all function within interface
type geometry interface {
	perimeter() float64
	area() float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	rect1 := rect{
		width:  20,
		height: 5,
	}

	cir1 := circle{
		radius: 7,
	}

	measure(&rect1)
	measure(&cir1)
}
