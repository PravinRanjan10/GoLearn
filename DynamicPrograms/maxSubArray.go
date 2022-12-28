/*
53. Maximum Subarray

Given an integer array nums, find the subarray
which has the largest sum and return its sum.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
*/

// solution:
/*
	1. Find the currSum at every index, by adding curSum and current value.
	2. If currSum > maxSum:
		maxSum = curSum
	3. if currSum < 0, means, adding current value is not meaningful, then make
		curSum = 0, and proceed for next element
*/

package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	maxSum := nums[0]
	currSum := 0
	for i := 0; i < n; i++ {
		currSum = currSum + nums[i]

		if currSum > maxSum {
			maxSum = currSum
		}

		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("max subarray sum:", maxSubArray(nums))
}

// output:
// max subarray sum: 6
