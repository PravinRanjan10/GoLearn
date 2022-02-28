// In SingleList: Insert a new node at end
// In SingleList: Insert a new node at begining
// In SingleList: Insert a new node at Kth position

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

func InsertAtBegining(head *Node, data int) *Node {
	node := getNode(data)
	if head == nil {
		head = node
	} else {
		node.next = head
		head = node
	}
	return head

}

func InsertAtKthPos(head *Node, pos, data int) *Node {
	if pos < 1 {
		fmt.Println("Invalid Position!")
	}
	node := getNode(data)
	if pos == 1 {
		node.next = head
		head = node
	} else {
		k := 2
		temp := head
		for pos > k && temp != nil {
			k += 1
			temp = temp.next
		}
		if temp != nil {
			node.next = temp.next
			temp.next = node
		} else {
			fmt.Println("Pos is out of range!!")
		}

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
		head = InsertAtEnd(head, i)
	}
	Println(head)
	fmt.Println("**********")
	head = InsertAtBegining(head, 200)
	Println(head)

	fmt.Println("**********")
	head = InsertAtKthPos(head, 5, 300)
	Println(head)

}
