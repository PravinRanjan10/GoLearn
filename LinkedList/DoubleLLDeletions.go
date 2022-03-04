package main

import "fmt"

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

func CreateDLL(head *Node, data int) *Node {
	newNode := getNode(data)
	if head == nil {
		head = newNode
	} else {
		temp := head
		for temp.right != nil {
			temp = temp.right
		}
		temp.right = newNode
		newNode.left = temp
	}
	return head
}

func Println(head *Node) {
	for head != nil {
		fmt.Println("data:", head.data)
		head = head.right
	}
}

func Delete1stNode(head *Node) *Node {
	if head == nil {
		return head
	}
	head = head.right
	return head
}

func DeleteLastNode(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.right == nil {
		return nil
	}
	temp := head
	for temp.right.right != nil {
		temp = temp.right
	}
	temp.right.left = nil
	temp.right = nil
	return head
}

func DeleteKthNode(head *Node, pos int) *Node {
	if pos < 1 {
		fmt.Println("Invalid pos!!!")
	}
	if pos == 1 {
		head = Delete1stNode(head)
	} else {
		k := 2
		temp := head
		for pos > k && temp.right != nil {
			temp = temp.right
			k += 1
		}
		if temp.right != nil {
			temp.right = temp.right.right
			temp.right.left = temp
		} else {
			fmt.Println("Pos out of index!!")
		}
	}

	return head
}

func main() {
	var head *Node

	for i := 0; i < 10; i++ {
		head = CreateDLL(head, i)
	}

	Println(head)
	fmt.Println("******")
	head = Delete1stNode(head)
	Println(head)
	fmt.Println("****DeleteLastNode****")
	head = DeleteLastNode(head)
	Println(head)
	fmt.Println("****DeleteKthNoe****")
	head = DeleteKthNode(head, 2)
	Println(head)

}
