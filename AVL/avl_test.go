package avl

import (
	"testing"
)

func TestAvlTreeOperations(t *testing.T) {
	tree := &AvlTree{}

	// 测试插入
	values := []int{10, 20, 30, 25, 5, 35}
	for _, val := range values {
		tree.Insert(val)
	}

	// 测试搜索
	for _, val := range values {
		if node := tree.Search(val); node == nil || node.Data != val {
			t.Errorf("Failed to find %d in the AVL tree", val)
		}
	}

	// 测试不在树中的值
	if node := tree.Search(100); node != nil {
		t.Errorf("Search should not have found 100 in the AVL tree, but found %+v", node)
	}

	// 测试删除
	tree.Remote(20) // 删除一个节点并检查平衡
	if node := tree.Search(20); node != nil {
		t.Errorf("Failed to remove 20 from the AVL tree")
	}

	// 测试删除后的搜索
	for _, val := range []int{10, 30, 25, 5, 35} {
		if node := tree.Search(val); node == nil || node.Data != val {
			t.Errorf("Failed to find %d in the AVL tree after removal", val)
		}
	}

	// 测试删除根节点
	tree.Remote(10) // 删除根节点
	if node := tree.Search(10); node != nil {
		t.Errorf("Failed to remove the root node 10 from the AVL tree")
	}

	// 测试删除后的树仍然平衡
	verifyTreeIsBalanced(tree, t)
}

func verifyTreeIsBalanced(tree *AvlTree, t *testing.T) {
	var checkBalance func(node *TreeNode) int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	checkBalance = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := checkBalance(node.Left)
		rightHeight := checkBalance(node.Right)
		balance := leftHeight - rightHeight

		if balance > 1 || balance < -1 {
			t.Errorf("Node %d is unbalanced with a balance factor of %d", node.Data, balance)
			return 0 // 如果不平衡，返回0防止后续计算
		}
		return max(leftHeight, rightHeight) + 1
	}

	height := checkBalance(tree.Root)
	if height == 0 {
		t.Errorf("Tree is not balanced")
	}
}
