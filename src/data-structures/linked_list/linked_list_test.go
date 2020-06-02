package main

import (
	"testing"
)


func TestPushBack(t *testing.T) {
	var size int
	l := LinkedList{}

	size = l.Size()
	if size != 0 {
		t.Fatalf("Size() should return 0, got: %d", size)
	}

	l.PushBack(1)
	size = l.Size()
	if size != 1 {
		t.Fatalf("Size() should return 1, got %d", size)
	}
	if l.head.data != 1 {
		t.Fatalf("Head's data should be 1, got %d", l.head.data)
	}

	l.PushBack(2)
	size = l.Size()
	if size != 2 {
		t.Fatalf("Size() should return 2, got %d", size)
	}
	if l.head.Next().Next() != nil {
		t.Fatalf("Second node should be the last node")
	}
	if l.head.Next().data != 2 {
		t.Fatalf("Last node's data should be 2, got %d", l.head.Next().data)
	}
}

func TestPopBack(t *testing.T) {
	var value interface{}
	var size int
	l := LinkedList{}

	l.PushBack(1)
	value = l.PopBack()
	if value != 1 {
		t.Fatalf("Returned value should be 1, got %+v", value)
	}
	size = l.Size()
	if size != 0 {
		t.Fatalf("Size() should return 0, got %d", size)
	}

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	value = l.PopBack()
	if value != 3 {
		t.Fatalf("Returned value should be 3, got %+v", value)
	}
	size = l.Size()
	if size != 2 {
		t.Fatalf("Size() should return 2, got %d", size)
	}
	if l.head.next.next != nil {
		t.Fatalf("Second node should not have next node")
	}
}

func TestDeleteKth(t *testing.T) {
	var err error
	l := LinkedList{}

	l.PushBack(1)
	err = l.DeleteKth(1)
	if err != nil {
		t.Fatalf("DeleteKth() returned error: %s", err.Error())
	}
	if l.Contains(1) != false {
		t.Fatal("List should not contain 1")
	}

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	// 3 means 3rd element not value of 3
	err = l.DeleteKth(3)
	if err != nil {
		t.Fatalf("DeleteKth() returned error: %s", err.Error())
	}
	if l.Contains(3) != false {
		t.Fatal("List should not contain 3")
	}

	err = l.DeleteKth(4)
	if err == nil || err.Error() != "k = 4 exceeds size of the list" {
		t.Fatal("Expected error with specific message")
	}
}

func TestContains(t *testing.T) {
	l := LinkedList{}

	if l.Contains(1) != false {
		t.Fatal("List should not contain 1")
	}

	l.PushBack(1)
	if l.Contains(1) != true {
		t.Fatal("List should contain 1")
	}

	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	if l.Contains(3) != true {
		t.Fatal("List should contain 3")
	}

	if l.Contains(6) != false {
		t.Fatal("List should not contain 6")
	}
}

func TestFind(t *testing.T) {
	var node *Node
	l := LinkedList{}

	node = l.Find(1)
	if node != nil {
		t.Fatal("Find(1) on empty list should return nil")
	}

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(5)
	l.PushBack(4)

	node = l.Find(2)
	if node.Data() != 2 {
		t.Fatal("Data() should return 2")
	}
	if node != l.head.next {
		t.Fatal("Expected returned node to be list's second node")
	}

	node = l.Find(4)
	if node.Data() != 4 {
		t.Fatal("Data() should return 4")
	}

	if l.Find(0) != nil {
		t.Fatal("Find(0) should return nil")
	}

	if l.Find(nil) != nil {
		t.Fatal("Find(nil) should return nil")
	}
}

func TestRemoveAfter(t *testing.T) {
	var node *Node
	l := LinkedList{}

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	node = l.Find(2)

	if l.RemoveAfter(node) != true {
		t.Fatal("RemoveAfter() should return true")
	}
	if l.Size() != 3 {
		t.Fatal("Size() should return 3")
	}
	if node.Next().Data() != 4 {
		t.Fatal("node.Next().Data() should return 4")
	}

	if l.RemoveAfter(l.Find(4)) != false {
		t.Fatal("RemoveAfter() should return false")
	}
	if l.Size() != 3 {
		t.Fatal("Size() should return 3")
	}

	// Try to remove node that doesn't belong to the list
	node = &Node{
		1,
		l.Find(2),
	}

	if l.RemoveAfter(node) != false {
		t.Fatal("RemoveAfter() should return false")
	}
	if l.Size() != 3 {
		t.Fatal("Size() should return 3")
	}

	if l.RemoveAfter(nil) != false {
		t.Fatal("RemoveAfter(nil) should return false")
	}
}

func TestGetNodeIfIsAMember(t *testing.T) {
	var arg *Node
	var ret *Node
	l := LinkedList{}

	arg = &Node{
		1,
		nil,
	}

	ret = l.GetNodeIfIsAMember(arg)
	if ret != nil {
		t.Fatal("GetNodeIfIsAMember() should return nil")
	}

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	arg = l.Find(3)

	ret = l.GetNodeIfIsAMember(arg)
	if ret != arg {
		t.Fatal("GetNodeIfIsAMember() should return a node")
	}
}

func TestInsertAfter(t *testing.T) {
	var first, second *Node
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	first = l.Find(2)
	second = &Node{
		4,
		nil,
	}

	l.InsertAfter(first, second)
	if l.Find(4) != second {
		t.Fatal("List should contain the 'second' node")
	}
	if second.Next().Data() != 3 {
		t.Fatal("'second' next node should contain 3")
	}
}

func TestInsertAfterNodeIsNotAMember(t *testing.T) {
	var first, second *Node
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	first = &Node{1, l.Find(2)}
	second = &Node{5, nil}

	l.InsertAfter(first, second)
	if l.Size() != 3 {
		t.Fatal("Size() should return 3")
	}
	if l.Find(1).Next() != l.Find(2) {
		t.Fatal("Second node should not change")
	}
}

func TestRemoveKeyOnlyHead(t *testing.T) {
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(2)

	l.RemoveKey(1)
	if l.head.Data() != 2 {
		t.Fatal("Head should be 2")
	}
	if l.Size() != 1 {
		t.Fatal("Size() should return 1")
	}
}

func TestRemoveKeyWholeList(t *testing.T) {
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(1)
	l.PushBack(1)
	l.PushBack(1)

	l.RemoveKey(1)
	if l.head != nil {
		t.Fatal("Head should be nil")
	}
	if l.Size() != 0 {
		t.Fatal("Size() should return 0")
	}
}

func TestRemoveKeyMiddleElement(t *testing.T) {
	l := LinkedList{}
	l.PushBack(2)
	l.PushBack(1)
	l.PushBack(2)

	l.RemoveKey(1)
	if l.head.Data() != 2 {
		t.Fatal("Head should be nil")
	}
	if l.head.Next().Data() != 2 {
		t.Fatal("Head should be nil")
	}
	if l.Size() != 2 {
		t.Fatal("Size() should return 2")
	}
}

func TestRemoveKeyOnlyTail(t *testing.T) {
	l := LinkedList{}
	l.PushBack(2)
	l.PushBack(2)
	l.PushBack(1)

	l.RemoveKey(1)
	if l.head.Data() != 2 {
		t.Fatal("Head should be nil")
	}
	if l.head.Next().Data() != 2 {
		t.Fatal("Head should be nil")
	}
	if l.Size() != 2 {
		t.Fatal("Size() should return 2")
	}
}

func TestMaxMiddle(t *testing.T) {
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(1)

	ret, err := l.Max()
	if err != nil {
		t.Fatalf("Max() returned err: %s", err.Error())
	}
	if ret != 3 {
		t.Fatalf("Max() should return 3")
	}
}

func TestMaxFirst(t *testing.T) {
	l := LinkedList{}
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(1)
	l.PushBack(1)

	ret, err := l.Max()
	if err != nil {
		t.Fatalf("Max() returned err: %s", err.Error())
	}
	if ret != 3 {
		t.Fatalf("Max() should return 3")
	}
}

func TestMaxLast(t *testing.T) {
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(1)
	l.PushBack(3)

	ret, err := l.Max()
	if err != nil {
		t.Fatalf("Max() returned err: %s", err.Error())
	}
	if ret != 3 {
		t.Fatalf("Max() should return 3")
	}
}

func TestMaxEmpty(t *testing.T) {
	l := LinkedList{}

	ret, err := l.Max()
	if err != nil {
		t.Fatalf("Max() returned err: %s", err.Error())
	}
	if ret != 0 {
		t.Fatalf("Max() should return 0")
	}
}

func TestMaxRecursiverLast(t *testing.T) {
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(1)
	l.PushBack(3)
	var max interface{}
	max = 0

	ret, err := l.MaxRecursive(l.Head(), max)
	if err != nil {
		t.Fatalf("Max() returned err: %s", err.Error())
	}
	if ret != 3 {
		t.Fatalf("Max() should return 3")
	}
}

func TestReverse(t *testing.T) {
	l := LinkedList{}
	// Check that reversing empty list does't panic
	l.Reverse()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Reverse()

	if l.head.Data() != 3 {
		t.Fatal("New head should be 3")
	}
	if l.head.next.data != 2 {
		t.Fatal("New second node should be 2")
	}
	if l.head.next.next.data != 1 {
		t.Fatal("New third node should be 1")
	}
}

func TestReverseTwo(t *testing.T) {
	l := LinkedList{}
	l.PushBack(1)
	l.PushBack(2)

	l.Reverse()

	if l.head.Data() != 2 {
		t.Fatal("New head should be 2")
	}
	if l.head.next.data != 1 {
		t.Fatal("New second node should be 1")
	}
}

func TestReverseOne(t *testing.T) {
	l := LinkedList{}
	l.PushBack(1)

	l.Reverse()

	if l.head.Data() != 1 {
		t.Fatal("New head should be 1")
	}
}
