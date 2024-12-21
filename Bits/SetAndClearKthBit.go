package main

import (
	"fmt"
	"strconv"
)

func setKthBit(n, k int) int {
	return (1 << k) | n // left shift then OR operation
}

func clearBits(n, k int) int {
	return n & ^(1 << k) // invert then perform and operation on number
}

func PrintIntToBinary(n int) {
	fmt.Println("Binary of Number after clearBits:", strconv.FormatInt(int64(n), 2))
}

func main() {
	n, k := 63, 4
	fmt.Println("Number after setBit:", setKthBit(n, k))
	y := clearBits(n, k)
	fmt.Println("Number after clearBits=:", y)
	PrintIntToBinary(y)
}
