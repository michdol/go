package main

import (
	"testing"
)


func TestNewStack(t *testing.T) {
	lls := LLStack{}
	if lls.top != nil {
		t.Error("Top should be nil")
	}
	if lls.Size() != 0 {
		t.Error("Size() should return 0")
	}
	if lls.IsEmpty() != true {
		t.Error("IsEmpty() should return true")
	}
}

func TestPush(t *testing.T) {
	var node *Node
	lls := LLStack{}

	node = lls.Push(1)
	if node.data != 1 {
		t.Error("node's data should be 1")
	}
	if node.next != nil {
		t.Error("Node's next should be nil")
	}

	if lls.top != node {
		t.Error("Returned node should be stack's Top")
	}
	if lls.Size() != 1 {
		t.Error("Size() should return 1")
	}
	if lls.IsEmpty() != false {
		t.Error("IsEmpty() should return false")
	}

	// ====== second Push
	secondNode := lls.Push(1)
	if secondNode.data != 1 {
		t.Error("node's data should be 1")
	}
	if secondNode.next != node {
		t.Error("Node's next should be previous node")
	}

	if lls.top != secondNode {
		t.Error("Returned node should be stack's Top")
	}
	if lls.Size() != 2 {
		t.Error("Size() should return 2")
	}
}

func TestPopFromEmptyStack(t *testing.T) {
	lls := LLStack{}

	val, err := lls.Pop()
	if val != nil {
		t.Errorf(`Pop returned %+v, nil expected`, val)
	}

	if err == nil || err.Error() != "Can't pop from empty stack" {
		t.Errorf("Expected error, got %+v instead", err)
	}
}

func TestPop(t *testing.T) {
	var val interface{}
	var err error
	lls := LLStack{}
	lls.Push(1)
	lls.Push(2)

	val, err = lls.Pop()
	if val != 2 {
		t.Errorf("Expected val to be 2, got %d instead", val)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, got %+v instead", err)
	}
	if lls.Size() != 1 {
		t.Errorf("Size() should return 1")
	}
	if lls.top.data != 1 {
		t.Error("Stack's top data should be 1")
	}

	val, err = lls.Pop()
	if val != 1 {
		t.Errorf("Expected val to be 2, got %d instead", val)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, got %+v instead", err)
	}
	if lls.Size() != 0 {
		t.Error("Size() should return 0")
	}
	if lls.top != nil {
		t.Error("Stack's top should be nil")
	}
}