package binarySearchTree

type TreeNode struct {
	Data int
	Left *TreeNode
	Right *TreeNode
}

type binarySearchTree struct {
	Root *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode {
		Data: data,
		Left: nil,
		Right: nil,
	}
}

func (bst *binarySearchTree) Search(data int) *TreeNode {
	thisNode := bst.Root
	for thisNode != nil {
		if thisNode.Data == data {
			break
		}else if thisNode.Data < data {
			thisNode = thisNode.Right
		}else {
			thisNode = thisNode.Left
		}
	}
	return thisNode
}

func (bst *binarySearchTree) Insert(data int) {
	thisNode := bst.Root
	// The current node is initially the root node
	if thisNode == nil {
		bst.Root = NewTreeNode(data)
		// If the root node is empty, the inserted node is the root node
		return
	}
	var preNode *TreeNode = nil
	for thisNode != nil {
		if thisNode.Data == data {
			// If there is a node with the value of num in the binary search tree, the system returns directly
			return
		}
		preNode = thisNode
		if thisNode.Data < data {
			thisNode = thisNode.Right
		}else {
			thisNode = thisNode.Left
		}
	}
	Newnode := NewTreeNode(data)
	if preNode.Data < data {
		preNode.Right = Newnode
	}else {
		preNode.Left = Newnode
	}
}

func (bst *binarySearchTree) Delete(data int) {
	thisNode := bst.Root
	if thisNode == nil {
		return
	}
	var preNode *TreeNode = nil
	for thisNode != nil {
		if thisNode.Data == data {
			// Determine whether the value of the current node is the value to be deleted, and if so, exit the loop directly
			break
		}
		preNode = thisNode
		// If not, step on to the next node
		if thisNode.Data < data {
			thisNode = thisNode.Right
		}else {
			thisNode = thisNode.Left
		}
	}
	if thisNode == nil {
		return
	}
	// deletion
	if thisNode.Left == nil || thisNode.Right == nil {
		var child *TreeNode = nil
		if thisNode.Left != nil {
			child = thisNode.Left
		}else {
			child = thisNode.Right
		}
		// Determines whether the left node or the right node of the current node is not empty, and passes the value that is not empty to the current node
		if thisNode != bst.Root {
			if preNode.Left == thisNode {
				preNode.Left = child
			}else {
				preNode.Right = child
			}
		}else {
			bst.Root = child
		}
		// If the current node is the root node and has only one subtree, the subtree is treated as the new root node
	}else {
		// If the current node has two subtrees (the current node includes the root node), the value of the new current node is determined using the intermediate order traversal
		tmp := thisNode.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		bst.Delete(tmp.Data)
		thisNode.Data = tmp.Data
	}
}

func (bst *binarySearchTree) InorderHelper() {
	InorderHelper(bst.Root)
}

func InorderHelper(node *TreeNode) {
	if node == nil {
		return
	}
	InorderHelper(node.Left)
	InorderHelper(node.Right)
}