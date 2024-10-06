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
	// 初始化链表
	l := InitList[int]()

	// 添加节点
	l.AddNode(10)
	l.AddNode(20)
	l.AddNode(30)

	// 打印链表状态
	fmt.Println("Initial list state:")
	printList(l)

	// 删除节点 10
	l.DeleteNode(10)

	// 打印链表状态
	fmt.Println("List state after deleting 10:")
	printList(l)

	// 设置数据
	l.SetData(20, 25)

	// 打印链表状态
	fmt.Println("List state after setting data 20 to 25:")
	printList(l)

	// 尝试查找节点
	node, found := l.Find(25)
	if found {
		fmt.Printf("Node with data 25 found: %v\n", node.Data)
	} else {
		t.Errorf("Node with data 25 not found")
	}
}
