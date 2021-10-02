package main
import (
	"fmt"
	"sort"
)
func main()  {
	var arr []int32
	arr = []int32{0, 4 , 3, 8, 7, 9, 1, 10}
	fmt.Println("arr before sort:", arr)
	//sort.Ints(arr) // it sorts only []int
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] }) // it sort other than []int
	fmt.Println("arr after sort:", arr)
}