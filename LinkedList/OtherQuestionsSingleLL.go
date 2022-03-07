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

// Find middle of a LL
func FindMiddleNode(head *Node) *Node {
	ptr1 := head
	ptr2 := head

	for ptr2 != nil && ptr2.next != nil {
		ptr1 = ptr1.next
		ptr2 = ptr2.next.next
	}
	return ptr1
}

// Display a LL from the end (just print no need to reverse)
func PrintInReverse(head *Node) {
	if head == nil {
		return
	}

	PrintInReverse(head.next)
	fmt.Println("data:", head.data)
}

// 2nd approach, using two pointers. Move 1st pointer n-node ahead and then
// move both pointer together
func FindNthNodeFromEnd2(head *Node, n int) {
	if head == nil {
		return
	}
	ptr1 := head
	ptr2 := head
	k := 1
	for ptr1.next != nil && n > k {
		ptr1 = ptr1.next
		k += 1
	}
	for ptr1.next != nil {
		ptr1 = ptr1.next
		ptr2 = ptr2.next
	}

	fmt.Printf("%+v: from last is: %+v\n", n, ptr2.data)
}

// Simpleway. Find length and then subtract lenght - n.
func FindNthNodeFromEnd1(head *Node, n int) *Node {
	if n < 1 {
		fmt.Println("Invalid Input")
		return head
	}
	if n == 1 {
		return head
	}
	count := 1
	temp := head
	for temp != nil {
		temp = temp.next
		count += 1
	}
	temp = head
	for i := 1; i < count-n; i++ {
		temp = temp.next
	}
	return temp
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

	// fmt.Println("****PrintMiddleNode****")
	// mnode := FindMiddleNode(head)
	// fmt.Println("Middle Node data:", mnode.data)

	// fmt.Println("****PrintInReverse****")
	// PrintInReverse(head)

	// fmt.Println("****FindKthNodeFromEnd****")
	// Println(head)

	// kthNode := FindNthNodeFromEnd(head, 3)
	// fmt.Println("Kth data:", kthNode.data)
	FindNthNodeFromEnd2(head, 3)

}
