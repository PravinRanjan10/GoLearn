package main

import "fmt"

func main() {
	str := "abccdba"
	length := len(str)

	l := 0
	r := length - 1
	for l <= r {
		if str[l] != str[r] {
			fmt.Println("Not Palindrome!")
			return
		}
		l += 1
		r -= 1
	}
	fmt.Println("Palindrome!")
}
