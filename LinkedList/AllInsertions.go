// In SingleList: Insert a new node at end
// 1. check head is nil then head will be new node
// 2. Assign a new pointer as head and traverse till nil
// 3. Insert at last

// In SingleList: Insert a new node at begining
//	1. check head is nil, then head will be new node
//  2. Insert node at beg

// In SingleList: Insert a new node at Kth position
//  1. If pos < 1: return
//  2. If pos ==1 : Insert at beg
//  3. check pos > k && temp != nil
//	4. Insert at kth pos
//  5. if temp == nil, meaning, pos is out of index

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
