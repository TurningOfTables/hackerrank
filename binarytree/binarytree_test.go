package main

import (
	"testing"
)

func TestCreateNode(t *testing.T) {
	res := CreateNode(5)
	expectedVal := 5

	if res.Val != expectedVal {
		t.Errorf("expected %v got %v", expectedVal, res)
	}

	if res.Left != nil || res.Right != nil {
		t.Errorf("expected nil got %v", res)
	}
}

func TestInitTree(t *testing.T) {
	res := InitTree()
	expectedRootVal := 4

	// Doesn't test all nodes, just the root
	if res.Val != expectedRootVal {
		t.Errorf("expected %v got %v", expectedRootVal, res)
	}
}

func ExampleTreeNode_PreOrder() {
	res := InitTree()
	res.PreOrder()
	// Output: 4 2 1 3 7 6 9
}

func ExampleTreeNode_PostOrder() {
	res := InitTree()
	res.PostOrder()
	// Output: 1 3 2 6 9 7 4
}

func ExampleTreeNode_MidOrder() {
	res := InitTree()
	res.MidOrder()
	// Output: 1 2 3 4 6 7 9
}

func ExampleInvertTree() {
	tree := InitTree()
	reversed := InvertTree(tree)
	reversed.PreOrder()
	// Output: 4 7 9 6 2 3 1
}
