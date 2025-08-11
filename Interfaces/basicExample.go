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

	//********* 1st Approach *********
	var t tank
	t = rectangle{5, 6, 10}
	fmt.Println("rectangle area:", t.Area())

	t = circle{5}
	fmt.Println("Circle area:", t.Area())

	// ******** 2nd Approach **********
	c := circle{
		radius: 4.5,
	}
	fmt.Println("rectangle area:", c.Area())

	r := rectangle{
		length: 4,
		breath: 3,
		height: 1,
	}
	fmt.Println("Circle area:", r.Area())

}
