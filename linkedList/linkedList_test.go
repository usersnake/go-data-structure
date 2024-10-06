package linkedList

import (
	"fmt"
	"testing"
)

func printList[T comparable](l *List[T]) {
	fmt.Print("List: ")
	for node := l.HeadNode; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Data)
	}
	fmt.Println()
}

func TestListOperations(t *testing.T) {

	l := InitList[int]()

	l.AddNode(10)
	l.AddNode(20)
	l.AddNode(30)

	fmt.Println("Initial list state:")
	printList(l)

	l.DeleteNode(10)

	fmt.Println("List state after deleting 10:")
	printList(l)

	l.SetData(20, 25)

	fmt.Println("List state after setting data 20 to 25:")
	printList(l)

	node, found := l.Find(25)
	if found {
		fmt.Printf("Node with data 25 found: %v\n", node.Data)
	} else {
		t.Errorf("Node with data 25 not found")
	}
}
