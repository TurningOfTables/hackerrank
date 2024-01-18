package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateNode returns a new TreeNode with value v and empty left/right branches
func CreateNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

// InitTree creates a new tree with arbitrary values to test with
func InitTree() *TreeNode {
	root := CreateNode(4)
	root.Left = CreateNode(2)
	root.Right = CreateNode(7)
	root.Left.Left = CreateNode(1)
	root.Left.Right = CreateNode(3)
	root.Right.Left = CreateNode(6)
	root.Right.Right = CreateNode(9)
	return root
}

// PreOrder recursively traverses and prints a tree in the order Node, Left, Right
// See https://en.wikipedia.org/wiki/Tree_traversal#Pre-order,_NLR
func (root *TreeNode) PreOrder() {
	if root != nil {
		fmt.Print(root.Val, " ")
		root.Left.PreOrder()
		root.Right.PreOrder()
	}
}

// PostOrder recursively traverses and prints a tree in the order Left, Right, Node
// See https://en.wikipedia.org/wiki/Tree_traversal#Post-order,_LRN
func (root *TreeNode) PostOrder() {
	if root != nil {
		root.Left.PostOrder()
		root.Right.PostOrder()
		fmt.Print(root.Val, " ")
	}
}

// MidOrder recursively traverses and prints a tree in the order Left, Node, Right
// See https://en.wikipedia.org/wiki/Tree_traversal#In-order,_LNR
func (root *TreeNode) MidOrder() {
	if root != nil {
		root.Left.MidOrder()
		fmt.Print(root.Val, " ")
		root.Right.MidOrder()
	}
}

// InvertTree recursively reverses the left/right nodes of a binary tree
func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left

		InvertTree(root.Left)
		InvertTree(root.Right)
	}

	return root
}

func main() {
	tree := InitTree()

	tree.PreOrder()
	fmt.Println()

	inverted := InvertTree(tree)
	inverted.PreOrder()
}
