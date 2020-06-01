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

func TestDequeue(t *testing.T) {
	var value interface{}
	llq := LLQueue{}
	llq.Enqueue(1)
	llq.Enqueue(2)
	llq.Enqueue(3)

	value = llq.Dequeue()
	if value != 1 {
		t.Error("Dequeue() should return 1")
	}
	if llq.Size() != 2 {
		t.Error("Size() should return 2")
	}
	if llq.first.data != 2 {
		t.Error("First in queue should be 2")
	}

	value = llq.Dequeue()
	if value != 2 {
		t.Error("Dequeue() should return 2")
	}
	if llq.Size() != 1 {
		t.Error("Size() should return 1")
	}

	value = llq.Dequeue()
	if value != 3 {
		t.Error("Dequeue() should return 3")
	}
	if llq.Size() != 0 {
		t.Error("Size() should return 0")
	}
	if llq.IsEmpty() != true {
		t.Error("IsEmpty() should return true")
	}
	if llq.first != nil || llq.last != nil {
		t.Error("Queue should be empty")
	}
}