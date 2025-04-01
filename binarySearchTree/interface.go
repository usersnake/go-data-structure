package binarySearchTree

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type bst interface {
	Search(data int) *TreeNode

	Insert(data int)

	Delete(data int)

	InorderHelper()
}
