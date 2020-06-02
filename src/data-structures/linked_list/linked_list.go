package main

import (
	"errors"
	"fmt"
)


type Node struct {
	data		interface{}
	next		*Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Data() interface{} {
	return n.data
}

func (self *Node) LessThan(other interface{}) (bool, error) {
	switch v := self.data.(type) {
	case int:
		otherValue, ok := other.(int)
		if !ok {
			return false, errors.New(fmt.Sprintf("Invalid data type %+v", otherValue))
		}
		return v < otherValue, nil
	// TODO: other cases
	default:
		fmt.Println(other)
	}
	return true, nil
}

type LinkedList struct {
	head		*Node
	size		int
}

func (self *LinkedList) Head() *Node {
	return self.head
}

func (self *LinkedList) InsertAfter(first *Node, second *Node) {
	current := self.GetNodeIfIsAMember(first)
	if current != nil {
		second.next = current.next
		current.next = second
		self.IncreaseSize()
	}
}

func (self *LinkedList) RemoveKey(key interface{}) {
	var previous *Node
	current := self.head
	for current != nil {
		if current.Data() == key {
			if current == self.head {
				self.head = current.next
				current.next = nil
				current = self.head
				self.DecreaseSize()
				continue
			}
			previous.next = current.next
			current.next = nil
			self.DecreaseSize()
		}
		previous = current
		current = current.next
	}
}

func (l *LinkedList) PushBack(data interface{}) {
	if l.head == nil {
		l.head = &Node{
			data,
			nil,
		}
		l.IncreaseSize()
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{
		data,
		nil,
	}
	l.IncreaseSize()
}

func (l *LinkedList) PopBack() interface{} {
	var previous *Node
	var current *Node
	previous = l.head
	current = l.head
	for current.next != nil {
		previous = current
		current = current.next
	}
	if previous == nil {
		return nil
	}
	// Head only
	if previous == current {
		l.DecreaseSize()
		l.head = nil
		return current.data
	}
	previous.next = nil
	l.DecreaseSize()
	return current.data
}

func (l *LinkedList) DeleteKth(k int) error {
	if k > l.size {
		return errors.New(fmt.Sprintf("k = %d exceeds size of the list", k))
	}
	if k == 1 {
		l.DecreaseSize()
		l.head = l.head.next
		return nil
	}

	var previous *Node
	var current *Node
	previous = l.head
	current = l.head
	count := 1
	for count != k {
		previous = current
		current = current.next
		count += 1
	}
	previous.next = current.next
	l.DecreaseSize()
	return nil
}

func (self *LinkedList) RemoveAfter(node *Node) bool {
	/*
		Removes a node after the given node if the node belongs to the list
		Returns: true if a node has been removed, false otherwise
	*/
	current := self.GetNodeIfIsAMember(node)
	if current != nil && current.next != nil {
		current.next = current.next.next
		self.DecreaseSize()
		return true
	}
	return false
}

func (self *LinkedList) GetNodeIfIsAMember(node *Node) *Node {
	current := self.head
	for current != nil {
		if current == node {
			return current
		}
		current = current.next
	}
	return nil
}

func (self *LinkedList) Contains(data interface{}) bool {
	current := self.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func (self *LinkedList) Find(data interface{}) *Node {
	current := self.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil
}

func (self *LinkedList) IncreaseSize() {
	self.size += 1
}

func (self *LinkedList) DecreaseSize() {
	self.size -= 1
}

func (self *LinkedList) Size() int {
	return self.size
}

func (self *LinkedList) Max() (interface{}, error) {
	var max interface{}
	max = 0
	current := self.head
	for current != nil {
		less, err := current.LessThan(max)
		if err != nil {
			return false, err
		}
		if !less {
			max = current.Data()
		}
		current = current.next
	}
	return max, nil
}

func (self *LinkedList) MaxRecursive(current *Node, max interface{}) (interface{}, error) {
	if current == nil {
		return max, nil
	}
	less, err := current.LessThan(max)
	if err != nil {
		return false, err
	}
	if !less {
		max = current.Data()
	}
	return self.MaxRecursive(current.next, max)
}

func (self *LinkedList) Reverse() {
	first := self.head
	var second *Node
	var reverse *Node

	for first != nil {
		second = first.next
		first.next = reverse
		reverse = first
		first = second
	}
	self.head = reverse
}