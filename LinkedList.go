package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

func InsertLL(head *Node, data, i int) *Node {
	fmt.Printf("Create node:%v, data: %v\n", i, data)
	newNode := getNode(data)
	if head == nil {
		head = newNode
	} else {
		temp := head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
	return head
}

func PrintLL(head *Node) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func PrintLLRecursive(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.data)
	PrintLLRecursive(head.next)
}

func PrintLLRecursiveInReverseOrder(head *Node) {
	if head == nil {
		return
	}
	PrintLLRecursiveInReverseOrder(head.next)
	fmt.Println(head.data)
}

func CountLength(head *Node) int {
	var count int
	for head != nil {
		head = head.next
		count += 1
	}
	return count
}

func CountLengthRecursively(head *Node) int {
	if head == nil {
		return 0
	} else {
		return CountLength(head.next) + 1
	}
}

func FindSumElements(head *Node) int {
	var sum int
	for head != nil {
		sum += head.data
		head = head.next
	}
	return sum
}

func FindSumElementsRecursive(head *Node) int {
	if head == nil {
		return 0
	}
	return FindSumElementsRecursive(head.next) + head.data
}

func FindMaxElement(head *Node) int {
	var max int
	max = math.MinInt32
	if head == nil {
		max = -1
	} else {
		for head != nil {
			if head.data > max {
				max = head.data
			}
			head = head.next
		}
	}
	return max
}

func FindMaxElementRecursive(head *Node) int {
	if head == nil {
		return math.MinInt32
	} else {
		max1 := FindMaxElementRecursive(head.next)
		if head.data > max1 {
			max1 = head.data
		}
		return max1
	}
}

func SearchElementRecursive(head *Node, ele int) bool {
	if head == nil {
		return false
	}
	if head.data == ele {
		return true
	}
	return SearchElement(head.next, ele)
}

func SearchElement(head *Node, ele int) bool {
	if head == nil {
		return false
	}
	for head != nil {
		if head.data == ele {
			return true
		}
		head = head.next
	}
	return false
}

func ImproveSearchElement(head *Node, ele int) bool {
	if head == nil {
		return false
	}

	if head.data == ele {
		return true
	} else {
		var p, q *Node
		q = nil
		p = head
		for p != nil {
			if p.data == ele {
				q.next = p.next
				p.next = head
				head = p
				return true
			}
			q = p
			p = p.next
		}
	}
	return false
}

func main() {
	var head *Node = nil
	var i int
	fmt.Println("Create 10 node with random data")
	for i = 0; i < 10; i++ {
		head = InsertLL(head, rand.Intn(time.Now().Second()), i)
	}
	fmt.Println("********************************")
	fmt.Println("Print the all node of LinkedList:")
	PrintLL(head)

	fmt.Println("Print the all node of LinkedList recursively:")
	PrintLLRecursive(head)

	fmt.Println("Print the all node of LinkedList in reverse order:")
	PrintLLRecursiveInReverseOrder(head)

	fmt.Println("Find the length of LinkedList")
	fmt.Println("count =:", CountLength(head))

	fmt.Println("Find the length of LinkedList recursively")
	fmt.Println("count =:", CountLengthRecursively(head))

	fmt.Println("Find the total sum of LinkedList")
	fmt.Println("Sum =:", FindSumElements(head))

	fmt.Println("Find the total sum of LinkedList recursively")
	fmt.Println("Sum =:", FindSumElementsRecursive(head))

	fmt.Println("Find the max element of LinkedList")
	fmt.Println("Max =:", FindMaxElement(head))

	fmt.Println("Find the max element of LinkedList recursively")
	fmt.Println("Max =:", FindMaxElementRecursive(head))

	fmt.Println("Search an element in LinkedList")
	fmt.Println("found =:", SearchElement(head, 15))

	fmt.Println("Search an element in LinkedList recursively")
	fmt.Println("found =:", SearchElementRecursive(head, 15))

	fmt.Println("Improve searching in LinkedList")
	fmt.Println("found =:", ImproveSearchElement(head, 0))
}
