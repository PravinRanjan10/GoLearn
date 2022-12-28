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

// recursive way to find nth fibonacci number
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// generate nth fibonacci number using DP
func fibUsingDP(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fibNumber := fib(5)
	fmt.Println("The nth fibonacci number:", fibNumber)
	fibNumber = fibUsingDP(5)
	fmt.Println("The nth fibonacci number using DP:", fibNumber)
}

// Output:
// The nth fibonacci number: 55
// The nth fibonacci number using DP: 5
