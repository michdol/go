package main

import (
	"errors"
)


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
func (self *DoublyLinkedList) RemoveFront() (interface{}, error) {
	if self.IsEmpty() {
		return nil, self.emptyListError()
	}
	self.DecreaseSize()
	value := self.front.data
	// In case there is only 1 item before the function is called
	// self.front is set to nil
	self.front = self.front.next
	if self.size == 0 {
		self.back = nil
	} else {
		self.front.previous = nil
	}
	return value, nil
}

func (self *DoublyLinkedList) RemoveBack() (interface{}, error) {
	if self.IsEmpty() {
		return nil, self.emptyListError()
	}
	self.DecreaseSize()
	value := self.back.data
	// In case there is only 1 item before the function is called
	// self.back is set to nil
	self.back = self.back.previous
	if self.size == 0 {
		self.front = nil
	} else {
		self.back.next = nil
	}
	return value, nil
}

func (self *DoublyLinkedList) emptyListError() error {
	return errors.New("List is empty")
}

func (self *DoublyLinkedList) InsertBefore(node *DoubleNode, data interface{}) error {
	targetNode, err := self.GetNodeIfIsAMember(node)
	if err != nil {
		return err
	}
	newNode := &DoubleNode{
		data: data,
		next: targetNode,
		previous: targetNode.previous,
	}
	if targetNode == self.front {
		self.front = newNode
	}
	oldPrevious := targetNode.previous
	if oldPrevious != nil {
		oldPrevious.next = newNode
	}
	targetNode.previous = newNode
	self.IncreaseSize()
	return nil
}

func (self *DoublyLinkedList) GetNodeIfIsAMember(node *DoubleNode) (*DoubleNode, error) {
	current := self.front
	for current != nil {
		if current == node {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("Passed node does not belong to the list")
}

func (self *DoublyLinkedList) InsertAfter(node *DoubleNode, data interface{}) error {
	targetNode, err := self.GetNodeIfIsAMember(node)
	if err != nil {
		return err
	}
	newNode := &DoubleNode{
		data: data,
		next: targetNode.next,
		previous: targetNode,
	}
	if targetNode == self.back {
		self.back = newNode
	}
	oldNext := targetNode.next
	if oldNext != nil {
		oldNext.previous = newNode
	}
	targetNode.next = newNode
	self.IncreaseSize()
	return nil
}

func (self *DoublyLinkedList) RemoveNode(node *DoubleNode) error {
	targetNode, err := self.GetNodeIfIsAMember(node)
	if err != nil {
		return err
	}
	self.DecreaseSize()
	if targetNode == self.front {
		self.front = targetNode.next
	}
	if targetNode == self.back {
		self.back = targetNode.previous
	}
	previous := targetNode.previous
	next := targetNode.next
	if previous != nil {
		previous.next = next
	}
	if next != nil {
		next.previous = previous
	}
	return nil
}