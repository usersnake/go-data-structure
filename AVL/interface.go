package avl

type avl interface {
	Search(data int) *TreeNode

	Remote(data int)

	Insert(data int)
}
