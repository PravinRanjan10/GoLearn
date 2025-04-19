package main

import "fmt"

// recursive approach of finding LCS
// Reference: https://www.youtube.com/watch?v=sSno9rV8Rhg
func lcsRecursive(s1, s2 string, i, j int) int {
	if i == 0 || j == 0 {
		return 0
	}

	if s1[i-1] == s2[j-1] {
		return 1 + lcsRecursive(s1, s2, i-1, j-1)
	} else {
		return max(lcsRecursive(s1, s2, i-1, j), lcsRecursive(s1, s2, i, j-1))
	}
}

// Memoization is also know as top down approach
func lcsMemoization(s1, s2 string) int {
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return lcsHelper(s1, s2, 0, 0, memo)
}

func lcsHelper(s1, s2 string, i, j int, memo [][]int) int {
	if i == len(s1) || j == len(s2) {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if s1[i] == s2[j] {
		memo[i][j] = 1 + lcsHelper(s1, s2, i+1, j+1, memo)
	} else {
		memo[i][j] = max(lcsHelper(s1, s2, i+1, j, memo), lcsHelper(s1, s2, i, j+1, memo))
	}
	return memo[i][j]
}

// Bottom-up, efficient solution
func lcsBottomUp(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	// Create a 2D slice (m+1 x n+1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := "AGGTAB"
	s2 := "GXTXAYB"
	fmt.Println(lcsRecursive(s1, s2, len(s1), len(s2))) // outtput: 4, GTAB
	fmt.Println(lcsMemoization(s1, s2))                 // outtput: 4, GTAB
	fmt.Println(lcsBottomUp(s1, s2))
}
