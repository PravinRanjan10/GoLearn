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

func InsertAtBegining(head *Node, data int) *Node {

	node := getNode(data)
	if head == nil {
		head = node
	} else {
		node.right = head
		head.left = node
		head = node
	}

	return head
}

func InsertAtEnd(head *Node, data int) *Node {
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

func InsertAtKthPos(head *Node, data, pos int) *Node {
	if pos < 1 {
		fmt.Println("Invalid Position!!")
	}
	node := getNode(data)
	if pos == 1 {
		node.right = head
		head.left = node
		head = node
	} else {
		k := 2
		temp := head
		for pos > k && temp != nil {
			k += 1
			temp = temp.right
		}
		if temp != nil {
			node.right = temp.right
			temp.right.left = node
			temp.right = node
			node.left = temp
		} else {
			fmt.Println("Pos is out of range!")
		}

	}
	return head
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
		head = InsertAtBegining(head, i)
	}
	Println(head)
	head = InsertAtEnd(head, 25)
	fmt.Println("*******")
	Println(head)
	fmt.Println("*******")
	head = InsertAtKthPos(head, 100, 5)
	Println(head)

}
