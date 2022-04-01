// Print all Permutations of string
// Ex: S="abc", O/P:{abc, acb, bac, bca, cba, cab}

package main

import (
	"fmt"
)

func Permute(s string, l, r int) {
	if l == r {
		fmt.Println(s)
	} else {
		s1 := []rune(s)
		for i := l; i < r; i++ {
			s1[l], s1[i] = s1[i], s1[l] // swap
			Permute(string(s1), l+1, r) // backtracking
			s1[l], s1[i] = s1[i], s1[l] // swap again
		}
	}

}

func main() {

	s := "abc"
	Permute(s, 0, len(s))
}

// Note: This solution, print all permutation(with repeative). For non-repeative, use map/set during print
// TimeComplexity: O(n**n), i.e n factorial
