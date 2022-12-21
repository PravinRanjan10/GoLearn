package main

import "fmt"

func binarySearch(arr []int, low, high, key int) int {

	if high < low {
		return -1
	}
	mid := (low + high) / 2

	if arr[mid] == key {
		return mid
	}

	if arr[mid] > key {
		return binarySearch(arr, low, mid-1, key)
	}
	return binarySearch(arr, mid+1, high, key)
}

func main() {

	arr := []int{1, 2, 4, 5, 7, 8, 9, 13, 16}
	lengthArray := len(arr)
	key := 16
	fmt.Printf("element found at index:%d\n", binarySearch(arr, 0, lengthArray, key))
}
