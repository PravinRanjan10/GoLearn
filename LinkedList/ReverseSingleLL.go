package main

import "fmt"

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

func InsertAtEnd(head *Node, data int) *Node {
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

func ReverseLLRecursive(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return head
	}
	first := ReverseLLRecursive(head.next)
	head.next.next = head
	head.next = nil
	return first
}

func Reverse(head *Node) *Node {
	if head == nil {
		return nil
	}
	var prev *Node
	prev = nil
	curr := head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	head = prev
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
		head = InsertAtEnd(head, i)
	}
	Println(head)
	fmt.Println("*****Recursive*****")
	head = ReverseLLRecursive(head)
	Println(head)
	fmt.Println("*****Non-Recursive*****")
	head = Reverse(head)
	Println(head)
}
