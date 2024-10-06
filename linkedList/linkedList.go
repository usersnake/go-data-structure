package linkedList

type ListNode[T comparable] struct {
	Data T
	Pre *ListNode[T]
	Next *ListNode[T]
}

type List[T comparable] struct {
	HeadNode *ListNode[T]
	TailNode *ListNode[T]
	Hash map[T]*ListNode[T]
}

func InitList[T comparable]() *List[T] {
	return &List[T]{
		Hash: make(map[T]*ListNode[T]),
	}
}

func (l *List[T]) AddNode(val T) {
	newNode := &ListNode[T]{Data: val}
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

func (l *List[T]) DeleteNode(val T) {
	thisNode, ok := l.Hash[val]
	if !ok {
		return
	}else {
		delete(l.Hash, val)
		if thisNode == l.HeadNode {
			l.HeadNode = thisNode.Next
			if l.HeadNode != nil {
				l.HeadNode.Pre = nil
			}
		}else if thisNode.Pre != nil {
			thisNode.Pre.Next = thisNode.Next
		}
		if thisNode == l.TailNode {
			l.TailNode = thisNode.Pre
			if l.TailNode != nil {
				l.TailNode.Next = nil
			}
		}else if thisNode.Next != nil{
			thisNode.Next.Pre = thisNode.Pre
		}
	}
}

func (l *List[T]) SetData(olddata, newdata T) {
	thisNode, ok := l.Hash[olddata]
	if !ok {
		return
	}
	thisNode.Data = newdata
	delete(l.Hash, olddata)
	l.Hash[newdata] = thisNode
}

func (l *List[T]) Find(data T) (*ListNode[T], bool) {
	thisNode, ok := l.Hash[data]
	return thisNode, ok
}