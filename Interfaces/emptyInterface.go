package main

import "fmt"

func main() {
	var i, j interface{}
	println("******Int operation******")
	i = 42
	j = 42
	describe(i, j)
	println("******String operation******")
	i = "hello"
	j = "bye"
	describe(i, j)
	println("******multiple operation******")
	fmt.Println("AddOne", AddOne(10.5))
}

func describe(i, j interface{}) {
	fmt.Println("print value of i and j:", i, j)

	x, a := i.(int)
	y, b := j.(int)
	fmt.Println("operation on i and j:", x+1, y+5, a, b)

}

func AddOne(n interface{}) interface{} {
	var empty interface{}
	switch nn := n.(type) {
	case int:
		nn += 1
		return nn
	case int8:
		nn += 1
		return nn
		// ... etc...
	case float64:
		nn += 1
		return nn
	}
	return empty
}

/*
Output:
******Int operation******
print value of i and j: 42 42
operation on i and j: 43 47 true true
******String operation******
print value of i and j: hello bye
operation on i and j: 1 5 false false
******multiple operation******
AddOne 11.5
*/
