//Given a value N, if we want to make change for N cents, and we have infinite supply of
// each of S = { S1, S2, .. , Sm} valued coins, how many ways can we make the change?
// The order of coins doesn’t matter.
// For example, for N = 4 and S = {1,2,3}, there are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
// So output should be 4. For N = 10 and S = {2, 5, 3, 6}, there are five
// solutions: {2,2,2,2,2}, {2,2,3,3}, {2,2,6}, {2,3,5} and {5,5}. So the output should be 5.

package main

import (
	"fmt"
	"math"
)

func minCoins(arr1 []int, n int, memory []int) int {

	if n == 0 {
		return 0
	}

	currMin := math.MaxInt

	for i := 0; i < len(arr1); i++ {

		if n-arr1[i] >= 0 {
			subAns := 0
			if memory[n-arr1[i]] != -1 {
				subAns = 1 + memory[n-arr1[i]]
			} else {
				subAns = 1 + minCoins(arr1, n-arr1[i], memory)
			}
			if subAns != math.MaxInt && subAns < currMin {
				currMin = subAns
			}
		}

	}
	memory[n] = currMin
	return memory[n]

}

func main() {

	n := 18
	var memory = make([]int, n+1)

	arr1 := []int{1, 5, 7}

	for i := 0; i < len(memory); i++ {
		memory[i] = -1
	}

	memory[0] = 0

	minCoin := minCoins(arr1, n, memory)
	fmt.Println("The min coin change required:", minCoin)

}
