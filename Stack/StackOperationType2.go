// stack implementation type2: using slice array
package main

import "fmt"

var stack []int32

func IsStackEmpty() bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

func Push(data int32) {
	stack = append(stack, data)
}

func Pop() {
	if IsStackEmpty() {
		fmt.Println("Stack Underflow!!")
		return
	}
	n := len(stack)
	stack = stack[:n-1]
}

func main() {
	Push(5)
	Push(2)
	Push(8)
	Push(9)
	fmt.Println("Current Stack:", stack)
	Pop()
	Pop()
	fmt.Println("Current Stack:", stack)
	Push(20)
	Push(30)
	fmt.Println("Current Stack:", stack)
}
