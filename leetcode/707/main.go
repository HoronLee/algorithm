package main

import "fmt"

// 链表节点定义
type Node struct {
	Val  int
	Next *Node
}

// 链表定义
type MyLinkedList struct {
	dummyHead *Node
	Size      int
}

// 构造新的链表
func Constructor() MyLinkedList {
	node := &Node{-999, nil}
	return MyLinkedList{
		dummyHead: node,
		Size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	// 判定数据是否有效
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	// 真实头节点
	cur := this.dummyHead.Next
	for range index {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	node.Next = this.dummyHead.Next
	this.dummyHead.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.Size {
		return
	}
	node := &Node{Val: val}
	cur := this.dummyHead
	for range index {
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.dummyHead
	for range index {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	this.Size--
}

func main() {
	list := Constructor()     // 初始化链表
	list.AddAtHead(100)       // 在头部添加元素
	list.AddAtTail(242)       // 在尾部添加元素
	list.AddAtTail(777)       // 在尾部添加元素
	list.AddAtIndex(1, 99999) // 在指定位置添加元素
	list.printLinkedList()    // 打印链表
}

// 打印链表
func (list *MyLinkedList) printLinkedList() {
	cur := list.dummyHead // 设置当前节点为虚拟头节点
	for cur.Next != nil { // 遍历链表
		fmt.Println(cur.Next.Val) // 打印节点值
		cur = cur.Next            // 切换到下一个节点
	}
}
