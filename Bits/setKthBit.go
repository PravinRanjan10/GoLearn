package main

import "fmt"

func setKthBit(n, k int) int {
	return (1 << k) | n
}

func main() {
	n, k := 10, 2
	fmt.Println(setKthBit(n, k))
}
