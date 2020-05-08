package main

import (
	"testing"
)


func TestGetLeftIndex(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1},
	}

	if heap.GetLeftIndex(0) != 1 {
		t.Error(`GetLeftIndex(0) should return 1`)
	}
	if heap.GetLeftIndex(1) != 3 {
		t.Error(`GetLeftIndex(1) should return 3`)
	}
	if heap.GetLeftIndex(2) != 5 {
		t.Error(`GetLeftIndex(2) should return 5`)
	}
}

func TestGetRightIndex(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1},
	}

	if heap.GetRightIndex(0) != 2 {
		t.Error(`GetRightIndex(0) should return 2`)
	}
	if heap.GetRightIndex(1) != 4 {
		t.Error(`GetRightIndex(1) should return 4`)
	}
	if heap.GetRightIndex(2) != 6 {
		t.Error(`GetRightIndex(2) should return 6`)
	}
}

func TestGetParentIndex(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1, 0},
	}

	if heap.GetParentIndex(0) != -1 {
		t.Error(`GetParentIndex(0) should return -1`)
	}
	if heap.GetParentIndex(1) != 0 {
		t.Error(`GetParentIndex(1) should return 0`)
	}
	if heap.GetParentIndex(2) != 0 {
		t.Error(`GetParentIndex(2) should return 0`)
	}
	if heap.GetParentIndex(3) != 1 {
		t.Error(`GetParentIndex(3) should return 1`)
	}
	if heap.GetParentIndex(4) != 1 {
		t.Error(`GetParentIndex(4) should return 1`)
	}
	if heap.GetParentIndex(51) != 25 {
		t.Error(`GetParentIndex(51) should return 25`)
	}
}

func TestHasLeftChild(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1},
	}

	if heap.HasLeftChild(0) != true {
		t.Error(`HasLeftChild(0) should return true`)
	}
	if heap.HasLeftChild(1) != true {
		t.Error(`HasLeftChild(1) should return true`)
	}
	if heap.HasLeftChild(2) != false {
		t.Error(`HasLeftChild(2) should return false`)
	}
	if heap.HasLeftChild(99) != false {
		t.Error(`HasLeftChild(99) should return false`)
	}
}

func TestHasRightChild(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1},
	}

	if heap.HasRightChild(0) != true {
		t.Error(`HasRightChild(0) should return true`)
	}
	if heap.HasRightChild(1) != false {
		t.Error(`HasRightChild(1) should return false`)
	}
	if heap.HasRightChild(99) != false {
		t.Error(`HasRightChild(99) should return false`)
	}
}

func TestHasParent(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1, 0},
	}

	if heap.HasParent(0) != false {
		t.Error(`HasParent(0) should return false`)
	}
	if heap.HasParent(1) != true {
		t.Error(`HasParent(1) should return true`)
	}
	if heap.HasParent(4) != true {
		t.Error(`HasParent(4) should return true`)
	}
	if heap.HasParent(5) != false {
		t.Error(`HasParent(5) should return false`)
	}
}