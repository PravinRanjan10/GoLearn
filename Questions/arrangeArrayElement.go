//Given an array of elements of length N, ranging from 0 to N – 1. All elements
// may not be present in the array. If the element is not present then there will
// be -1 present in the array. Rearrange the array such that A[i] = i and if i is
// not present, display -1 at that place
//*
// Input : arr = {-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
// Output : [-1, 1, 2, 3, 4, -1, 6, -1, -1, 9]

// Input : arr = {19, 7, 0, 3, 18, 15, 12, 6, 1, 8,
//               11, 10, 9, 5, 13, 16, 2, 14, 17, 4}
// Output : [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
//          11, 12, 13, 14, 15, 16, 17, 18, 19]

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello..")
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}

	for i := 0; i < len(arr); i++ {

		if arr[i] > 0 && arr[i] != i {
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		}
	}
	fmt.Println("arr:", arr)
}
