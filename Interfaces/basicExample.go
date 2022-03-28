package main

import (
	"fmt"
)

type tank interface {
	Area() int
}

type rectangle struct {
	length int
	breath int
	height int
}

type circle struct {
	radius float64
}

func (c circle) Area() int {
	return int((2 * 3.14 * c.radius * c.radius))
}

func (r rectangle) Area() int {
	area := r.length * r.breath * r.height
	return area
}

func main() {
	fmt.Println("I am in interfces!!")
	var t tank
	t = rectangle{5, 6, 10}
	fmt.Println("rectangle area:", t.Area())

	t = circle{5}
	fmt.Println("Circle area:", t.Area())
}
