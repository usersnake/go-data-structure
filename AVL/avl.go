package avl

type TreeNode struct {
	Data   int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

type AvlTree struct {
	Root *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (a *AvlTree) height(node *TreeNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}

func (a *AvlTree) setHeight(node *TreeNode) {
	leftheight := a.height(node.Left)
	rightheight := a.height(node.Right)

	if leftheight > rightheight {
		node.Height = leftheight + 1
	} else {
		node.Height = rightheight + 1
	}
}

func (a *AvlTree) balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return a.height(node.Left) - a.height(node.Right)
}

// dextrorotation
func (a *AvlTree) rightrotate(node *TreeNode) *TreeNode {
	// node is the first unbalanced node of the subtree
	child := node.Left
	granchild := child.Right
	child.Right = node
	node.Left = granchild
	// After the right-hand rotation, the node changes, and the height of the changed node needs to be updated
	a.setHeight(node)
	a.setHeight(child)
	return child
}

// levorotation
func (a *AvlTree) leftrotate(node *TreeNode) *TreeNode {
	child := node.Right
	granchild := child.Left
	child.Left = node
	node.Right = granchild

	a.setHeight(node)
	a.setHeight(child)
	return child
}

func (a *AvlTree) rotate(node *TreeNode) *TreeNode {
	balanceFactor := a.balanceFactor(node)
	if balanceFactor > 1 {
		if a.balanceFactor(node.Left) >= 0 {
			return a.rightrotate(node)
		} else {
			node.Left = a.leftrotate(node.Left)
			return a.rightrotate(node)
		}
	}
	if balanceFactor < -1 {
		if a.balanceFactor(node.Right) <= 0 {
			return a.leftrotate(node)
		} else {
			node.Right = a.rightrotate(node.Right)
			return a.leftrotate(node)
		}
	}
	return node
}

func (a *AvlTree) insertHelper(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return NewTreeNode(data)
	}
	if data < node.Data {
		node.Left = a.insertHelper(node.Left, data)
	} else if data > node.Data {
		node.Right = a.insertHelper(node.Right, data)
	} else {
		return node
	}
	a.setHeight(node)
	node = a.rotate(node)
	return node
}

func (a *AvlTree) Insert(val int) {
	a.Root = a.insertHelper(a.Root, val)
}

func (a *AvlTree) remoteHelper(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return nil
	}
	if data < node.Data {
		node.Left = a.remoteHelper(node.Left, data)
	} else if data > node.Data {
		node.Right = a.remoteHelper(node.Right, data)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				return nil
			} else {
				node = child
			}
		} else {
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = a.remoteHelper(node.Right, temp.Data)
			node.Data = temp.Data
		}
	}
	a.setHeight(node)
	node = a.rotate(node)
	return node
}

func (a *AvlTree) Remote(data int) {
	a.Root = a.remoteHelper(a.Root, data)
}

func (a *AvlTree) Search(data int) *TreeNode {
	thisNode := a.Root
	for thisNode != nil {
		if thisNode.Data == data {
			break
		} else if thisNode.Data < data {
			thisNode = thisNode.Right
		} else {
			thisNode = thisNode.Left
		}
	}
	return thisNode
}
