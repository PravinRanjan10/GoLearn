// Input:  [1, 2, 3, 4, 5]
// Output: [5, 4, 3, 2, 1]

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("arr:", arr)

	s := len(arr)

	l := 0
	r := s - 1

	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l += 1
		r -= 1
	}
	fmt.Println("Reversed arr:", arr)
}
