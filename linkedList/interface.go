package lists

type linkedLists[T comparable] interface {
	AddNode(val T)

	DeleteNode(val T)

	SetData(olddata, newdata T)

	Find(data T)
}