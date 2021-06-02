package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct{
	data int
	next *Node
}

func getNode(data int) *Node{
	var node Node
	node.data = data
	node.next = nil
	return &node
}

func InsertLL(head *Node, data, i int) *Node{
	fmt.Printf("Create node:%v, data: %v\n", i, data)
	newNode := getNode(data)
	if head == nil{
		head = newNode
	}else{
		temp := head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
	return head
}

func PrintLL(head *Node){
	for head != nil{
		fmt.Println(head.data)
		head = head.next
	}
}

func main() {
	var head *Node = nil
	var i int
	fmt.Println("Create 10 node with random data")
	for i = 0; i < 10; i++ {
		head = InsertLL(head, rand.Intn(time.Now().Second()), i)
	}
    fmt.Println("********************************")
	fmt.Println("Print the all node of LinkedList")
	PrintLL(head)
}
