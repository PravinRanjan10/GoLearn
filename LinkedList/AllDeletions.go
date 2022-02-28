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

func Delete1stNode(head *Node) *Node {
	if head == nil {
		fmt.Println("Nothing to delete!!")
		return head
	}
	newHead := head.next
	head.next = nil

	return newHead
}

func DeleteLastNode(head *Node) *Node {
	if head == nil {
		fmt.Println("Nothing to delete!")
		return head
	}
	if head.next == nil {
		head = nil
		return head
	}
	temp := head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	return head
}

func DeleteKthNode(head *Node, pos int) *Node {
	if head == nil {
		fmt.Println("Nothing to delete!")
		return head
	}
	if pos < 1 {
		fmt.Println("Invalid postion!")
		return head
	}
	if pos == 1 {
		head = head.next
	} else {
		k := 2
		temp := head
		for pos > k && temp.next != nil {
			k += 1
			temp = temp.next
		}
		if temp.next != nil {
			temp.next = temp.next.next
		} else {
			fmt.Println("Pos is out of range!")
		}
	}

	return head
}

func Println(head *Node) {
	for head != nil {
		fmt.Println("data:", head.data)
		head = head.next
	}
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

func main() {

	var head *Node
	for i := 0; i < 10; i++ {
		head = Insert(head, i)
	}
	Println(head)
	fmt.Println("*****")
	head = Delete1stNode(head)
	Println(head)

	fmt.Println("*****")
	head = DeleteLastNode(head)
	Println(head)

	fmt.Println("*****")
	head = DeleteKthNode(head, 4)
	Println(head)

}
