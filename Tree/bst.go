package main

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func getNode(data int) *TreeNode {

	var node TreeNode
	node.data = data
	node.left = nil
	node.right = nil

	return &node
}

func InsertBST(root *TreeNode, data int) *TreeNode {

	newNode := getNode(data)
	if root == nil {
		root = newNode
		return root
	} else {

		if data >= root.data {
			root.right = InsertBST(root.right, data)
		} else {
			root.left = InsertBST(root.left, data)
		}
	}

	return root
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	preOrder(root.left)
	preOrder(root.right)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Println(root.data)
	InOrder(root.right)
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.left)
	InOrder(root.right)
	fmt.Println(root.data)
}

func main() {
	var root *TreeNode

	for i := 0; i < 10; i++ {
		root = InsertBST(root, i*rand.Intn(20))
	}
	fmt.Println("PreOrder----\n")
	preOrder(root)

	fmt.Println("\nInOrder----\n")
	InOrder(root)

	fmt.Println("\nPostOrder----\n")
	PostOrder(root)
}
