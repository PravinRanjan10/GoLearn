/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?



Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

/*
solution:
very similar to fibonacci number.
dp[i] = dp[i-1] + dp[i-1] , because either you can climb 1 steps or 2 steps.
*/
package main

import "fmt"

func climbStairs(n int) int {

	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func main() {
	fmt.Println("total distict way to climb top:", climbStairs(5))
}

// output:
// total distict way to climb top: 8
