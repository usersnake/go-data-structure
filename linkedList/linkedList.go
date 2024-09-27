package linkedList

import(
	"fmt"
)

type Object interface {}

type Node struct {
	Key int
	Value Object
	Next *Node
}

type HashTable struct {
	Buckets []*Node
	HashLen int
}

func NewHashTable(hashLen int) *HashTable {
	buckets := make([]*Node, hashLen)
	for i := 0; i < hashLen; i++ {
		buckets[i] = &Node{}
	}
	return &HashTable{Buckets: buckets,HashLen: hashLen}
}

func (h *HashTable) Insert(key int, value Object) {
	index := key % h.HashLen
	newHead := &Node{Key: key, Value: value, Next: h.Buckets[index].Next}
	h.Buckets[index].Next = newHead
}

func (h *HashTable) Find(key int) (Object, bool) {
	index := key % h.HashLen
	for node := h.Buckets[index].Next; node != nil; node = node.Next {
		if node.Key == key {
			return node.Value, true
		}
	}
	return nil, false
}
func (h *HashTable) Change(key int,value Object) bool {
	index := key % h.HashLen
	node := h.Buckets[index].Next
	for node != nil {
		if node.Key == key {
			node.Value = value
			return true
		}
		node = node.Next
	}
	return false
}


func (h *HashTable) Delete(key int) bool {
	index := key % h.HashLen
	node := h.Buckets[index].Next
	pre := h.Buckets[index]
	node = pre.Next
	for node != nil {
		if node.Key == key {
			pre.Next = node.Next
			return true
		}
		pre = node
		node = node.Next
	}
	return false
}

func (ht *HashTable) Print() {
    for i, bucket := range ht.Buckets {
        fmt.Printf("Bucket %d: ", i)
        for node := bucket.Next; node != nil; node = node.Next {
            fmt.Printf("(%d, %v) -> ", node.Key, node.Value)
        }
        fmt.Println("nil")
    }
}