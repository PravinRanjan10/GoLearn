// Implementation of queue using slice
package main

import "fmt"

var queue []int32

func Enqueue(data int32) {
	queue = append(queue, data)
}

func Dequeue() {
	queue = queue[1:]
}

func main() {
	Enqueue(5)
	Enqueue(8)
	Enqueue(9)
	fmt.Println("Current queue:", queue)
	Dequeue()
	fmt.Println("Current queue:", queue)
}
