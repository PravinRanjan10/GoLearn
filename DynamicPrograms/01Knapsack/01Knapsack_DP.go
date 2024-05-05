package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Table [][]int

func kp(profit, w []int, W, n int) int {
	if W == 0 || n == 0 {
		return 0
	}

	if Table[n-1][W] != -1 {
		return Table[n-1][W]
	}

	if w[n-1] <= W {
		Table[n-1][W] = max(profit[n-1]+kp(profit, w, W-w[n-1], n-1), kp(profit, w, W, n-1))
		return Table[n-1][W]
	} else {
		return kp(profit, w, W, n-1)
	}
}

func main() {
	var W int
	profit := []int{60, 100, 120}
	w := []int{50, 20, 30}
	W = 50

	n := len(profit)

	Table = make([][]int, n+1)
	for i := range Table {
		Table[i] = make([]int, W+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			Table[i][j] = -1
		}
	}

	fmt.Println("Max Profit:", kp(profit, w, W, n))
}

