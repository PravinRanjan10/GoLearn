package main

import "fmt"

type Stack []int32

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(data int32) {
	*s = append(*s, data) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int32, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var stack Stack // create a stack variable of type Stack

	stack.Push(5)
	stack.Push(9)
	stack.Push(4)
	fmt.Println("Current Stack:", stack)
	stack.Pop()
	fmt.Println("Current Stack:", stack)

}
