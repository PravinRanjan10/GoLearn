package main

import "fmt"

// golbal variable for memoization
var memo [][]int

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
func lcsMemoization(s1, s2 string, n, m int) int {
	if memo[n][m] != -1 {
		return memo[n][m]
	}

	if m == 0 || n == 0 {
		memo[n][m] = 0
	} else {
		if s1[n-1] == s2[m-1] {
			memo[n][m] = 1 + lcsMemoization(s1, s2, n-1, m-1)
		} else {
			memo[n][m] = max(lcsMemoization(s1, s2, n, m-1), lcsMemoization(s1, s2, n-1, m))
		}
	}
	return memo[n][m]
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
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println("dpp==:", dp)
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

	memo = make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	fmt.Println(lcsMemoization(s1, s2, len(s1), len(s2))) // outtput: 4, GTAB
	fmt.Println(lcsBottomUp(s1, s2))
}
