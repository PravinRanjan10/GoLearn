// IMplement a queue usig librabry Linked lIst

package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)
	fmt.Println("The current queue length:", queue.Len())

	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Println("Thw queue value:", temp.Value)
	}

	queue.Remove(queue.Front())
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(".....Thw queue value:", temp.Value)
	}
}
