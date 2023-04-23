package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack(profit, weight []int, W, n int) int {
	// base condition
	if n == 0 || W == 0 {
		return 0
	}

	// If, weight of item is less than W, then it can be either included or not
	if weight[n-1] <= W {
		return max(profit[n-1]+knapsack(profit, weight, W-weight[n-1], n-1), knapsack(profit, weight, W, n-1))
	} else {
		return knapsack(profit, weight, W, n-1)
	}
}

func main() {
	profit := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	W := 50
	fmt.Println(knapsack(profit, weight, W, len(weight)))

}

// output: 220
