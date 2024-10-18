package binarySearchTree

import (
	"testing"
)

func TestBinarySearchTree_InorderTraversal(t *testing.T) {
	// create a newbinarysearchtree without anything
	bst := binarySearchTree{}

	// insert some elements
	values := []int{7, 3, 9, 1, 4, 8, 10}
	for _, v := range values {
		bst.Insert(v)
	}

	// Use the medium order to traverse the elements in the print tree
	var result []int
	appendValues := func(n *TreeNode) {
		if n != nil {
			result = append(result, n.Data)
		}
	}
	bst.inorderHelper(appendValues)

	// Verify that the results are in sort order
	expected := []int{1, 3, 4, 7, 8, 9, 10}
	if len(result) != len(expected) {
		t.Errorf("Got %v, want %v", result, expected)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Got %v, want %v at index %d", result, expected, i)
		}
	}
}

// inorderHelper wrapper to collect values
func (bst binarySearchTree) inorderHelper(visit func(*TreeNode)) {
	inorderHelper(bst.Root, visit)
}

// inorderHelper with visitor function
func inorderHelper(node *TreeNode, visit func(*TreeNode)) {
	if node == nil {
		return
	}
	inorderHelper(node.Left, visit)
	visit(node)
	inorderHelper(node.Right, visit)
}
