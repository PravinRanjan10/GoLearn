/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.

Example 1:

Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
*/

// solutions:
/*
 It's buttom-up approach. At every stage calculate the minCost. basically,
 it would be it's own cost + min(of last two)cost
 Cost[i] = Cost[i] + min(Cost[i-1], Cost[i-2])
*/
package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2 {
		return 0
	}

	for i := 2; i < n; i++ {
		cost[i] = cost[i] + min(cost[i-1], cost[i-2])
	}
	return min(cost[n-1], cost[n-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	cost := []int{10, 15, 20}
	fmt.Println("minCostClimbingStairs:", minCostClimbingStairs(cost))
}

// output:
// minCostClimbingStairs: 15
