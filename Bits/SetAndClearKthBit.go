// set the kth position of a given number
// clear the kth position

package main

import "fmt"

func setkthBit(n, k int) int {

	// logic of setting kth position.

	x := (1 << k) | n
	return x
}

func clearKthBit(n, k int) int {

	// logic of clear
	y := n &^ (1 << k)

	return y
}

func main() {
	fmt.Println("The number after set kth position is:", setkthBit(21, 3))
	fmt.Println("The number after clear kth position is:", clearKthBit(21, 2))

}
