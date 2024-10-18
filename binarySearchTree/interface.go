package binarySearchTree

type bst interface {
	Search(data int) *TreeNode

	Insert(data int)

	Delete(data int)

	InorderHelper()
}