package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello....")
	arr1 := []int32{1, 2, 3}
	arr2 := []int32{4, 5, 6}
	arr3 := []int32{}

	// 1. Append Example
	arr3 = append(arr3, arr1...) // can't append more than two slices at once
	arr3 = append(arr3, arr2...)
	fmt.Println("arr3:", arr3) // [1 2 3 4 5 6]

	// 2. Sort the slice in ascending
	arr4 := []int{2, 0, 9, 5, 6, 8, 1, 3}
	sort.Ints(arr4)
	fmt.Println("arr4:", arr4)

	// 3. To Reverse or Sort in descending order
	sort.Slice(arr4, func(i, j int) bool { return arr4[i] > arr4[j] })
	fmt.Println("arr4 inreverese:", arr4)

	// 4.Sort the slice
	arr5 := []int32{3, 4, 0, 9, 7, 8}
	sort.Slice(arr5, func(i, j int) bool { return arr5[i] < arr5[j] })
	fmt.Println("arr5:", arr5)

	// 5. Search element in sorted slice, if found return Index else return length(arr1)
	arr7 := []int{1, 2, 4, 8, 9, 10}
	fmt.Println("search element:", sort.SearchInts(arr7, 9))

	// 6. Sort the strings in ascending order
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println("ssss:", s)

	// 7. Iterate the slice
	arr6 := []int32{1, 3, 0, 9, 5, 8, 7, 4}

	for i, key := range arr6 {
		fmt.Println("Key:", i, key)
	}

}
