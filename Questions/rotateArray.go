package main

import (
	"fmt"
)

// Time: O(n), Space: O(n)
func basicWay1(arr1 []int, d int) []int {
	arr2 := []int{}
	arr2 = append(arr2, arr1[d:]...)
	arr2 = append(arr2, arr1[:d]...)
	return arr2
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}

// Time: O(n), Space: O(1)
// The Idea is:
// 1. Find gcd of len(arr) and d
// 2. Then every time

func AdvanceWay(arr []int, d int) []int {

	gcd := Gcd(len(arr), d)
	fmt.Println("gcd:", gcd)
	// TODO:

	return arr

}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	var d int
	fmt.Scanln(&d)

	fmt.Println("Output for basicWay:", basicWay1(arr1, d))
	fmt.Println("Output for AdvanceWay:", AdvanceWay(arr1, d))

}
