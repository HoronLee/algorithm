package main

import "fmt"

// Node 泛型双向链表节点
type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

// DoublyLinkedList 泛型双向链表
type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// Append 尾插
func (dll *DoublyLinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.Next = newNode
		newNode.Prev = dll.tail
		dll.tail = newNode
	}
	dll.size++
}

// Prepend 头插
func (dll *DoublyLinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.Next = dll.head
		dll.head.Prev = newNode
		dll.head = newNode
	}
	dll.size++
}

// Insert 指定位置插入
func (dll *DoublyLinkedList[T]) Insert(index int, value T) {
	if index <= 0 {
		dll.Prepend(value)
		return
	}
	if index >= dll.size {
		dll.Append(value)
		return
	}

	newNode := &Node[T]{Value: value}
	current := dll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	prev := current.Prev
	prev.Next = newNode
	newNode.Prev = prev
	newNode.Next = current
	current.Prev = newNode
	dll.size++
}

// RemoveByValue 按值删除
func (dll *DoublyLinkedList[T]) RemoveByValue(value T, equal func(a, b T) bool) {
	current := dll.head
	for current != nil {
		if equal(current.Value, value) {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				dll.head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				dll.tail = current.Prev
			}
			dll.size--
			return
		}
		current = current.Next
	}
}

// RemoveByIndex 按索引删除
func (dll *DoublyLinkedList[T]) RemoveByIndex(index int) {
	if index < 0 || index >= dll.size {
		fmt.Println("索引越界")
		return
	}
	current := dll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	if current.Prev != nil {
		current.Prev.Next = current.Next
	} else {
		dll.head = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
	} else {
		dll.tail = current.Prev
	}
	dll.size--
}

// Find 查找节点
func (dll *DoublyLinkedList[T]) Find(value T, equal func(a, b T) bool) *Node[T] {
	current := dll.head
	for current != nil {
		if equal(current.Value, value) {
			return current
		}
		current = current.Next
	}
	return nil
}

func main() {
	// 创建泛型双向链表，存储 int
	dllInt := &DoublyLinkedList[int]{}

	dllInt.Append(1)
	dllInt.Append(2)
	dllInt.Append(3)
	dllInt.Prepend(0)
	dllInt.Insert(2, 99)
}
