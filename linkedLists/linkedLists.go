package main

import "fmt"

type ListNode struct{
	Data int
	Next *ListNode
}

// linkedlists接收ListNode,定义Inster方法
func (linkedlists *ListNode) Inster(data int){
	// 如果输入值为nil，则退出Inster
	if linkedlists == nil{
		return
	}
	//使用newNode开辟新的后续节点
	newNode := &ListNode{Data : data}
	// cur指向节点
	cur := linkedlists
	// 如果cur.Next的值不为空，则指向下一个节点
	for cur.Next != nil{
		cur = cur.Next
	}
	// 将创建的新的后续节点赋给cur.Next
	cur.Next = newNode
}

func main(){
	node := &ListNode{Data : 10}
	node.Inster(20)
	node.Inster(30)
}