package main

import (
	"testing"
)


func TestPushFront(t *testing.T) {
	l := DoublyLinkedList{}

	l.PushFront(1)
	if l.Size() != 1 {
		t.Fatal("Size() should return 1")
	}
	if l.front == nil || l.front.data != 1 {
		t.Fatal("Front should be 1")
	}
	if l.back == nil || l.front != l.back {
		t.Fatal("Front should be Back")
	}
	if l.front.previous != nil {
		t.Fatal("Front's previous should be nil")
	}
	if l.front.next != nil {
		t.Fatal("Front's next should be nil")
	}

	// node		front 	back
	// data		2 			1
	// prev   nil 		front
	// next   back 		nil
	l.PushFront(2)
	if l.Size() != 2 {
		t.Fatal("Size() should return 2")
	}

	if l.front.data != 2 {
		t.Fatal("Front should be 2")
	}
	if l.front.previous != nil {
		t.Fatal("Front's previous should be nil")
	}
	if l.front.next != l.back {
		t.Fatal("Second node should be Back")
	}

	if l.back.data != 1 {
		t.Fatal("Back should be 1")
	}
	if l.back.previous != l.front {
		t.Fatal("Back's previous should be Front")
	}
	if l.back.next != nil {
		t.Fatal("Back's next should be nil")
	}

	// node		front 		node 		back
	// data		3		 			2				1
	// prev   nil   		front   node
	// next   node   		back 		nil
	l.PushFront(3)
	if l.Size() != 3 {
		t.Fatal("Size() should return 3")
	}

	// Check front
	if l.front.data != 3 {
		t.Fatal("Front should be 3")
	}
	if l.front.previous != nil {
		t.Fatal("Front's previous should be nil")
	}

	// Check middle
	node := l.front.next
	if node.data != 2 {
		t.Fatal("Second node should be 2")
	}
	if node.previous != l.front {
		t.Fatal("Second node's previous should be Front")
	}
	if node.next != l.back {
		t.Fatal("Middle node's next should be Back")
	}

	// Check back
	if l.back.data != 1 {
		t.Fatal("Back should be 1")
	}
	if l.back.previous != node {
		t.Fatal("Back's previous should be Middle node")
	}
	if l.back.next != nil {
		t.Fatal("Back's next should be nil")
	}
}

func TestPushBack(t *testing.T) {
	l := DoublyLinkedList{}

	l.PushBack(1)
	// node		front 		back
	// data		1		 			1
	// prev   nil   		nil
	// next   nil   		nil
	if l.Size() != 1 {
		t.Fatal("Size() should return 1")
	}
	if l.front != l.back {
		t.Fatal("Front and Back should point to the same node")
	}
	if l.front.next != nil || l.front.previous != nil {
		t.Fatal("Front's previous and next should be nil")
	}

	l.PushBack(2)
	// node		front 		back
	// data		1		 			2
	// prev   nil   		front
	// next   back   		nil
	if l.Size() != 2 {
		t.Fatal("Size() should return 2")
	}

	if l.front.data != 1 {
		t.Fatal("Front should be 1")
	}
	if l.front.next != l.back {
		t.Fatal("Front's next should be Back")
	}
	if l.front.previous != nil {
		t.Fatal("Front's previous should be nil")
	}

	if l.back.data != 2 {
		t.Fatal("Back should be 2")
	}
	if l.back.next != nil {
		t.Fatal("Back's next should be nil")
	}
	if l.back.previous != l.front {
		t.Fatal("Back's previous should be Front")
	}

	l.PushBack(3)
	// node		front 		back
	// data		1		 			2
	// prev   nil   		front
	// next   back   		nil
	if l.Size() != 3 {
		t.Fatal("Size() should return 3")
	}

	if l.front.data != 1 {
		t.Fatal("Front should be 1")
	}
	if l.front.previous != nil {
		t.Fatal("Front's previous should be nil")
	}

	node := l.front.next
	if node.data != 2 {
		t.Fatal("Middle should be 2")
	}
	if node.previous != l.front {
		t.Fatal("Middle's previous should be Front")
	}
	if node.next != l.back {
		t.Fatal("Middle's next should be Back")
	}

	if l.back.data != 3 {
		t.Fatal("Back should be 3")
	}
	if l.back.previous != node {
		t.Fatal("Back's previous should be Middle")
	}
	if l.back.next != nil {
		t.Fatal("Back's next should be nil")
	}
}

func TestRemoveFrontEmpty(t *testing.T) {
	l := DoublyLinkedList{}

	ret, err := l.RemoveFront()
	if ret != nil {
		t.Fatal("RemoveFront() should return nil")
	}
	if err == nil || err.Error() != "List is empty" {
		t.Fatal("Expected error")
	}
}

func TestRemoveFrontOneItem(t *testing.T) {
	l := DoublyLinkedList{}
	l.PushFront(1)

	ret, err := l.RemoveFront()
	if l.Size() != 0 {
		t.Fatal("Size() should return 0")
	}
	if ret != 1 {
		t.Fatal("RemoveFront() should return 1")
	}
	if err != nil {
		t.Fatal("RemoveFront() should NOT return error")
	}
	if l.front != nil || l.back != nil {
		t.Fatal("Front and Back should be nil")
	}
}

func TestRemoveFrontTwoItems(t *testing.T) {
	var ret interface{}
	var err error
	l := DoublyLinkedList{}
	l.PushFront(1)
	l.PushFront(2)

	ret, err = l.RemoveFront()
	if l.Size() != 1 {
		t.Fatal("Size() should return 1")
	}
	if ret != 2 {
		t.Fatal("RemoveFront() should return 2")
	}
	if err != nil {
		t.Fatal("RemoveFront() should NOT return error")
	}
	if l.front != l.back {
		t.Fatal("Front should be Back")
	}
	if l.front.previous != nil || l.front.next != nil {
		t.Fatal("Front's previous and next should be nil")
	}
	if l.front.data != 1 {
		t.Fatal("Front should be 1")
	}
}

func TestRemoveFrontMoreThanTwoItems(t *testing.T) {
	var ret interface{}
	var err error
	l := DoublyLinkedList{}
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	ret, err = l.RemoveFront()
	if l.Size() != 2 {
		t.Fatal("Size() should return 2")
	}
	if ret != 3 {
		t.Fatal("RemoveFront() should return 3")
	}
	if err != nil {
		t.Fatal("RemoveFront() should NOT return error")
	}
	if l.front.data != 2 {
		t.Fatal("Front should be 2")
	}
	if l.back.data != 1 {
		t.Fatal("Back should be 1")
	}
	if l.front.next != l.back {
		t.Fatal("Front's next should be Back")
	}
	if l.front.previous != nil {
		t.Fatal("Front's previous should be nil")
	}
	if l.back.previous != l.front {
		t.Fatal("Back's previous should be Front")
	}
	if l.back.next != nil {
		t.Fatal("Back's next should be nil")
	}
}

func TestRemoveFrontClearTheList(t *testing.T) {
	var ret interface{}
	var err error
	var currentSize, expectedSize int
	l := DoublyLinkedList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	initialLength := l.Size()
	for i := 0; i < initialLength; i++ {
		ret, err = l.RemoveFront()
		if ret != i + 1 {
			t.Fatalf("Value should be %d, got: %d", i + 1, ret)
		}
		if err != nil {
			t.Fatal("RemoveFront() should not return error")
		}
		expectedSize = initialLength - (i + 1)
		currentSize = l.Size()
		if currentSize != expectedSize {
			t.Fatalf("Size() should return %d, got: %d", expectedSize, currentSize)
		}
	}

	ret, err = l.RemoveFront()
	if ret != nil {
		t.Fatal("RemoveFront() should return nil")
	}
	if err == nil || err.Error() != "List is empty" {
		t.Fatal("Expected error")
	}
}

func TestRemoveBackEmpty(t *testing.T) {
	l := DoublyLinkedList{}

	ret, err := l.RemoveBack()
	if ret != nil {
		t.Fatal("RemoveFront() should return nil")
	}
	if err == nil || err.Error() != "List is empty" {
		t.Fatal("Expected error")
	}
}

func TestRemoveBackOneItem(t *testing.T) {
	l := DoublyLinkedList{}
	l.PushFront(1)

	ret, err := l.RemoveBack()
	if ret != 1 {
		t.Fatal("RemoveBack() should return 1")
	}
	if err != nil {
		t.Fatal("RemoveBack() should NOT return error")
	}
	if l.Size() != 0 {
		t.Fatal("Size() should return 0")
	}
	if l.front != nil || l.back != nil {
		t.Fatal("Front and Back should be nil")
	}
}

func TestRemoveBackTwoItems(t *testing.T) {
	var ret interface{}
	var err error
	l := DoublyLinkedList{}
	l.PushBack(1)
	l.PushBack(2)

	ret, err = l.RemoveBack()
	if l.Size() != 1 {
		t.Fatal("Size() should return 1")
	}
	if ret != 2 {
		t.Fatal("RemoveBack() should return 2")
	}
	if err != nil {
		t.Fatal("RemoveBack() should NOT return error")
	}
	if l.front != l.back {
		t.Fatal("Front should be Back")
	}
	if l.front.previous != nil || l.front.next != nil {
		t.Fatal("Front's previous and next should be nil")
	}
	if l.front.data != 1 {
		t.Fatal("Front should be 1")
	}
}

func TestRemoveBackThreeItems(t *testing.T) {
	var ret interface{}
	var err error
	l := DoublyLinkedList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	ret, err = l.RemoveBack()
	if l.Size() != 2 {
		t.Fatal("Size() should return 2")
	}
	if ret != 3 {
		t.Fatal("RemoveBack() should return 3")
	}
	if err != nil {
		t.Fatal("RemoveBack() should NOT return error")
	}
	if l.front.data != 1 {
		t.Fatal("Front should be 1")
	}
	if l.back.data != 2 {
		t.Fatal("Back should be 2")
	}
	if l.front.next != l.back {
		t.Fatal("Front's next should be Back")
	}
	if l.front.previous != nil {
		t.Fatal("Front's previous should be nil")
	}
	if l.back.previous != l.front {
		t.Fatal("Back's previous should be Front")
	}
	if l.back.next != nil {
		t.Fatal("Back's next should be nil")
	}
}

func TestRemoveBackClearTheList(t *testing.T) {
	var ret interface{}
	var err error
	var currentSize, expectedSize int
	l := DoublyLinkedList{}
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushFront(4)

	initialLength := l.Size()
	for i := 0; i < initialLength; i++ {
		ret, err = l.RemoveBack()
		if ret != i + 1 {
			t.Fatalf("Value should be %d, got: %d", i + 1, ret)
		}
		if err != nil {
			t.Fatal("RemoveBack() should not return error")
		}
		expectedSize = initialLength - (i + 1)
		currentSize = l.Size()
		if currentSize != expectedSize {
			t.Fatalf("Size() should return %d, got: %d", expectedSize, currentSize)
		}
	}

	ret, err = l.RemoveBack()
	if ret != nil {
		t.Fatal("RemoveBack() should return nil")
	}
	if err == nil || err.Error() != "List is empty" {
		t.Fatal("Expected error")
	}
}