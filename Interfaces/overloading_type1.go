/*
Function/Method Overloading means that that the same function/method name can be used with a different number and types of parameters
We can workaround Method/Function overloading in GO using

Variadic Function – A Variadic Function is a function that accepts a variable number of arguments
Empty Interface – It is an interface without any methods.
*/

// 1. Type1 overloading: A Variadic Function is a function that accepts a variable number of arguments

package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))
}

func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
