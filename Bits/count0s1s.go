package main

import "fmt"

func count0sAnd1s(n int) {
	count0, count1 := 0, 0
	for n != 0 {
		if n&1 == 1 {
			count1++
		} else {
			count0++
		}
		n >>= 1
	}
	fmt.Printf("count0: %d and count1: %d\n", count0, count1)

}

func main() {
	fmt.Println("Count set bet in number:")
	var n int
	fmt.Scanln(&n)
	count0sAnd1s(n)
}

