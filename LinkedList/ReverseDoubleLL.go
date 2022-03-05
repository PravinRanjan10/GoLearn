package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func getNode(data int) *Node {
	var node Node
	node.data = data
	node.left = nil
	node.right = nil
	return &node
}

func Insert(head *Node, data int) *Node {
	node := getNode(data)
	if head == nil {
		head = node
	} else {
		temp := head
		for temp.right != nil {
			temp = temp.right
		}
		temp.right = node
		node.left = temp
	}
	return head
}

func Reverse(head *Node) *Node {
	if head == nil || head.right == nil {
		return head
	}
	first := Reverse(head.right)
	head.right.right = head
	head.right = nil
	return first
}

func Println(head *Node) {
	for head != nil {
		fmt.Println("Data:", head.data)
		head = head.right
	}
}

func main() {
	var head *Node

	for i := 0; i < 10; i++ {
		head = Insert(head, i)
	}
	Println(head)
	fmt.Println("******")
	head = Reverse(head)
	Println(head)

}
