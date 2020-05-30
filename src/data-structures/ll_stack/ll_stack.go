package main

import (
	"errors"
)


type Node struct {
	data			interface{}
	next			*Node
}

func (n *Node) HasNext() bool {
	return n.next != nil
}

type LLStack struct {
	top 		*Node
	size		int
}

func (lls *LLStack) Push(data interface{}) *Node {
	top := lls.top
	newTop := &Node{
		data,
		top,
	}
	lls.top = newTop
	lls.size += 1
	return lls.top
}

func (lls *LLStack) Pop() (interface{}, error) {
	if lls.IsEmpty() {
		return nil, errors.New("Can't pop from empty stack")
	}
	top := lls.top
	if top.HasNext() {
		lls.top = top.next
	} else {
		lls.top = nil
	}
	lls.size -= 1
	return top.data, nil
}

func (lls *LLStack) Size() int {
	return lls.size
}

func (lls *LLStack) IsEmpty() bool {
	return lls.size == 0
}