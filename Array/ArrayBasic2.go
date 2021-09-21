package main
import (
	"fmt"
)
func main(){
	fmt.Printf("Enter size of your array: ")
	var size int
	var arr [10] int
	fmt.Scanln(&size)
	for i:=0; i<size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)
}