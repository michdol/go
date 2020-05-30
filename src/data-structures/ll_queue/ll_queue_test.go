package main

import (
	"testing"
)


func TestNewQueue(t *testing.T) {
	llq := LLQueue{}

	if llq.IsEmpty() != true {
		t.Error("IsEmpty() should return true")
	}
	if llq.Size() != 0 {
		t.Error("Size() should return 0")
	}
}

func TestEnqueue(t *testing.T) {
	var node *Node
	llq := LLQueue{}

	node = llq.Enqueue(1)
	if node.data != 1 {
		t.Error("Node's data should be 1")
	}
	if node.next != nil {
		t.Error("Node's next should be nil")
	}
	if llq.Size() != 1 {
		t.Error("Size() should return 1")
	}
	if llq.IsEmpty() != false {
		t.Error("IsEmpty() should return false")
	}
	if llq.first != node || llq.last != node {
		t.Error("Queue's first and last should be equal to returned node")
	}

	// ====== Second Enqueue
	secondNode := llq.Enqueue(2)
	if node.next != secondNode {
		t.Error("First node's next should be second node")
	}
	if llq.last != secondNode {
		t.Error("Queue's last should be second node")
	}
	if llq.first != node {
		t.Error("Queue's first should be first node")
	}
	if llq.Size() != 2 {
		t.Error("Size() should return 2")
	}

	// ====== Third Enqueue
	thirdNode := llq.Enqueue(3)
	if secondNode.next != thirdNode {
		t.Error("Second node's next should be third node")
	}
	if node.next != secondNode {
		t.Error("First node's next should be second node")
	}
	if llq.last != thirdNode {
		t.Error("Queue's last should be third node")
	}
	if llq.first != node {
		t.Error("Queue's first should be first node")
	}
	if llq.Size() != 3 {
		t.Error("Size() should return 3")
	}
}