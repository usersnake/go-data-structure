package lists

type listNode[T comparable] struct {
	Data T
	Pre *listNode[T]
	Next *listNode[T]
}

type LinkedLists[T comparable] struct {
	HeadNode *listNode[T]
	TailNode *listNode[T]
	Hash map[T]*listNode[T]
}

func NewList[T comparable]() *LinkedLists[T] {
	return &LinkedLists[T]{
		Hash: make(map[T]*listNode[T]),
	}
}

func (l *LinkedLists[T]) AddNode(val T) {
	newNode := &listNode[T]{Data: val}
	if l.HeadNode == nil {
		l.HeadNode = newNode
		l.TailNode = newNode
	}else {
		l.TailNode.Next = newNode
		newNode.Pre = l.TailNode
		l.TailNode = newNode
	}
	l.Hash[val] = newNode
}

func (l *LinkedLists[T]) DeleteNode(val T) {
	thisNode, ok := l.Hash[val]
	if !ok {
		return
	}
	delete(l.Hash, val)
	if thisNode == l.HeadNode && l.HeadNode != nil {
		l.HeadNode = thisNode.Next
		l.HeadNode.Pre = nil
	}else if thisNode.Pre != nil {
		thisNode.Pre.Next = thisNode.Next
	}
	if thisNode == l.TailNode && l.TailNode != nil {
		l.TailNode = thisNode.Pre
		l.TailNode.Next = nil
	}else if thisNode.Next != nil{
		thisNode.Next.Pre = thisNode.Pre
	}
}

func (l *LinkedLists[T]) SetData(olddata, newdata T) {
	thisNode, ok := l.Hash[olddata]
	if !ok {
		return
	}
	thisNode.Data = newdata
	delete(l.Hash, olddata)
	l.Hash[newdata] = thisNode
}

func (l *LinkedLists[T]) Find(data T) (*listNode[T], bool) {
	thisNode, ok := l.Hash[data]
	return thisNode, ok
}