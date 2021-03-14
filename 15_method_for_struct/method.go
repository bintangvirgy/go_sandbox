package main

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

// this function will become method of struct in parameter
// best practice : use pointer from struct
func (r *rectangle) area() float64 {
	return r.width * r.height
}

// but value can be used too
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	rect1 := rectangle{
		width:  20,
		height: 50,
	}

	fmt.Println(rect1)
	fmt.Println(rect1.area())
	fmt.Println(rect1.perimeter())

	rect2 := &rectangle{
		width:  10,
		height: 7,
	}

	fmt.Println(rect2)
	fmt.Println(rect2.area())
	fmt.Println(rect2.perimeter())

	rect3 := &rect1
	rect3.height = 21
	fmt.Println(rect3)
	fmt.Println(rect3.area())
	fmt.Println(rect3.perimeter())

	fmt.Println(rect1)
	fmt.Println(rect1.area())
	fmt.Println(rect1.perimeter())

	rect4 := rect1
	rect4.height = 28
	fmt.Println(rect4)
	fmt.Println(rect4.area())
	fmt.Println(rect4.perimeter())

	fmt.Println(rect1)
	fmt.Println(rect1.area())
	fmt.Println(rect1.perimeter())

}
