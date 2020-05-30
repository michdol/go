package main

import (
	"errors"
)


type Stack struct {
	arr			[]interface{}
	length	int
}

func (s *Stack) push(data interface{}) {
	s.arr = append(s.arr, data)
	s.length += 1
}

func (s *Stack) pop() (interface{}, error) {
	lastIdx := len(s.arr) - 1
	if lastIdx == -1 {
		return -1, errors.New("Can't pop from empty stack")
	}
	element := s.arr[lastIdx]
	s.arr = s.arr[:lastIdx]
	s.length -= 1
	return element, nil
}

func (s *Stack) top() (interface{}, error) {
	lastIdx := len(s.arr) - 1
	if lastIdx == -1 {
		return -1, errors.New("Stack is empty")
	}
	element := s.arr[lastIdx]
	return element, nil
}

func (s *Stack) len() int {
	return s.length
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) clear() {
	for !s.isEmpty() {
		s.pop()
	}
}

func (s Stack) Iterator() *StackIterator {
	return &StackIterator{
		s.len(),
		s,
	}
}

type StackIterator struct {
	i 		int
	stack Stack
}

func (si *StackIterator) HasNext() bool {
	return si.i > 0
}

func (si *StackIterator) Next() (interface{}, error) {
	si.i -= 1
	return si.stack.pop()
}