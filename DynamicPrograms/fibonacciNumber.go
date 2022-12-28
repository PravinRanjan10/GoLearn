/*
Fibonacci numbers

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
Fibonacci sequence, such that each number is the sum of the two preceding ones,
starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).
*/

package main

import "fmt"

func nThfibonacciNumber(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return nThfibonacciNumber(n-1) + nThfibonacciNumber(n-2)
}

func main() {
	fibNumber := nThfibonacciNumber(10)
	fmt.Println("The nth fibonacci number:", fibNumber)
}

// Output:
// The nth fibonacci number: 55
