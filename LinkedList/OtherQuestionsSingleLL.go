package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func getNode(data int) *Node {
	var node Node
	node.data = data
	node.next = nil

	return &node
}

func FindMiddleNode(head *Node) *Node {
	ptr1 := head
	ptr2 := head

	for ptr2 != nil && ptr2.next != nil {
		ptr1 = ptr1.next
		ptr2 = ptr2.next.next
	}
	return ptr1
}

func Insert(head *Node, data int) *Node {
	node := getNode(data)
	if head == nil {
		head = node
	} else {
		temp := head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	}
	return head

}

func Println(head *Node) {
	for head != nil {
		fmt.Println("Head-->:", head.data)
		head = head.next
	}
}

func main() {

	var head *Node
	for i := 0; i < 10; i++ {
		head = Insert(head, i)
	}
	Println(head)

	mnode := FindMiddleNode(head)
	fmt.Println("Middle Node data:", mnode.data)
}
