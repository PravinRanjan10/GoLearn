//2. Different number of parameters and of different types
// This case can be handled using both variadic function and empty interface

package main

import (
	"fmt"
	"reflect"
)

func main() {
	handle(1, "abc")
	handle("abc", "xyz", 3)
	handle(1, 2, 3, 4)
}

func handle(params ...interface{}) {
	fmt.Println("Handle func called with parameters:")
	for _, param := range params {
		fmt.Println("param and it's type:", param, reflect.TypeOf(param))
	}
}
