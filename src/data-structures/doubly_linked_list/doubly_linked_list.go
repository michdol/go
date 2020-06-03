package main

// import (
// 	"errors"
// 	"fmt"
// )


type DoubleNode struct {
	data			interface{}
	next			*DoubleNode
	previous 	*DoubleNode
}

func (self *DoubleNode) Next() *DoubleNode {
	return self.next
}

func (self *DoubleNode) Previous() *DoubleNode {
	return self.previous
}

func (self *DoubleNode) Data() interface{} {
	return self.data
}

func NewNode(data interface{}, next, previous *DoubleNode) *DoubleNode {
	return &DoubleNode{
		data,
		next,
		previous,
	}
}

type DoublyLinkedList struct {
	front		*DoubleNode
	back		*DoubleNode
	size		int
}

func (self *DoublyLinkedList) Front() *DoubleNode {
	return self.front
}

func (self *DoublyLinkedList) Back() *DoubleNode {
	return self.back
}

func (self *DoublyLinkedList) Size() int {
	return self.size
}

func (self *DoublyLinkedList) IsEmpty() bool {
	return self.size == 0
}

func (self *DoublyLinkedList) IncreaseSize() {
	self.size += 1
}

func (self *DoublyLinkedList) DecreaseSize() {
	self.size -= 1
}

func (self *DoublyLinkedList) PushFront(data interface{}) {
	newNode := &DoubleNode{
		data: data,
		next: self.front,
		previous: nil,
	}
	if self.IsEmpty() {
		self.back = newNode
	} else {
		self.front.previous = newNode
	}
	self.front = newNode
	self.IncreaseSize()
}

// PushBack
func (self *DoublyLinkedList) PushBack(data interface{}) {
	newNode := &DoubleNode{
		data: data,
		next: nil,
		previous: self.back,
	}
	if self.IsEmpty() {
		self.front = newNode
	} else {
		self.back.next = newNode
	}
	self.back = newNode
	self.IncreaseSize()
}

// RemoveFront
// RemoveBack
// InsertBefore
// InsertAfter
// RemoveNode

func (self *DoublyLinkedList) GetNodeIfIsAMember(node *DoubleNode) *DoubleNode {
	current := self.front
	for current != nil {
		if current == node {
			return current
		}
		current = current.next
	}
	return nil
}
