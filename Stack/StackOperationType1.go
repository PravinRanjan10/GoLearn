// stack implementation type2: using fixed size array

package main

import "fmt"

var stack [10]int32
var top int

func IsStackFull() bool {
	if len(stack) == 9 {
		return true
	}
	return false
}

func IsStackEmpty() bool {
	if len(stack) == -1 {
		return true
	}
	return false
}

func Push(data int32) {
	if IsStackFull() {
		fmt.Println("Stack Overflow!!")
		return
	}
	top += 1
	stack[top] = data
}

func Pop() {
	if IsStackEmpty() {
		fmt.Println("Stack Underflow!!")
		return
	}
	stack[top] = int32(-1)
	top -= 1
}

func main() {
	top = -1
	Push(5)
	Push(2)
	Push(8)
	Push(9)
	fmt.Println("Current Stack:", stack)
	Pop()
	Pop()
	fmt.Println("Current Stack:", stack)
}
