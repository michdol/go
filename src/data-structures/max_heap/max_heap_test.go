package main

import (
	"fmt"
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

func TestLeft(t *testing.T) {
	var ret int
	var err error
	heap := MaxHeap{
		[]int{4, 3, 2, 1, 0},
	}

	ret, err = heap.Left(0)
	if ret != 3 || err != nil {
		t.Error(`Left(0) should return 3 without error`)
	}

	ret, err = heap.Left(1)
	if ret != 1 || err != nil {
		t.Error(`Left(1) should return 1 without error`)
	}

	ret, err = heap.Left(2)
	if ret != -1 {
		t.Error(`Left(2) should return nil`)
	}
	if err == nil {
		t.Error(`Left(2) should return an error`)
	}
}

func TestRight(t *testing.T) {
	var ret int
	var err error
	heap := MaxHeap{
		[]int{4, 3, 2, 1, 0},
	}

	ret, err = heap.Right(0)
	if ret != 2 || err != nil {
		t.Error(`Right(0) should return 2 without error`)
	}

	ret, err = heap.Right(1)
	if ret != 0 || err != nil {
		t.Error(`Right(1) should return 0 without error`)
	}

	ret, err = heap.Right(2)
	if ret != -1 {
		t.Error(`Right(2) should return nil`)
	}
	if err == nil {
		t.Error(`Right(2) should return an error`)
	}
}

func TestParent(t *testing.T) {
	var ret int
	var err error
	heap := MaxHeap{
		[]int{4, 3, 2, 1, 0},
	}

	ret, err = heap.Parent(1)
	if ret != 4 || err != nil {
		t.Error(`Parent(1) should return 4 without error`)
	}
	ret, err = heap.Parent(2)
	if ret != 4 || err != nil {
		t.Error(`Parent(2) should return 4 without error`)
	}
	ret, err = heap.Parent(5)
	if ret != -1 || err == nil {
		t.Error(`Parent(5) should return -1 and error`)
	}
	ret, err = heap.Parent(0)
	if ret != -1 || err == nil {
		t.Error(`Parent(0) should return -1 and error`)
	}
}

func TestSwap(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1, 0},
	}

	if heap.Swap(0, 0) != nil {
		t.Error(`Swap returned error when it should not`)
	}
	if heap.Items[0] != 4 {
		t.Error(`Items array has changed when it should not`)
	}

	if heap.Swap(0, 1) != nil {
		t.Error(`Swap returned error when it should not`)
	}
	if heap.Items[0] != 3 {
		t.Error(`First item should be 3`)
	}
	if heap.Items[1] != 4 {
		t.Error(`Second item should be 4`)
	}

	if heap.Swap(0, 4) != nil {
		t.Error(`Swap returned error when it should not`)
	}
	if heap.Items[0] != 0 {
		t.Error(`First item should be 0`)
	}
	if heap.Items[4] != 3 {
		t.Error(`Last item should be 3`)
	}

	if heap.Swap(-1, 0) == nil {
		t.Error(`Swap should return error for incorrect indices but it did not`)
	}
	if heap.Swap(1, 5) == nil {
		t.Error(`Swap should return error for incorrect indices but it did not`)
	}
}

func TestHeapify(t *testing.T) {
	var heap MaxHeap
	heap = MaxHeap{
		[]int{4, 3, 2, 1, 5},
	}

	//				4
	//			/		\
	//		3			2
	//	/		\
	// 1		 5
	heap.Heapify(4)

	if heap.Items[0] != 5 {
		t.Error(`Items[0] should be 5`)
	}
	if heap.Items[1] != 4 {
		t.Error(`Items[1] should be 4`)
	}
	if heap.Items[2] != 2 {
		t.Error(`Items[2] should be 2`)
	}
	if heap.Items[3] != 1 {
		t.Error(`Items[3] should be 1`)
	}
	if heap.Items[4] != 3 {
		t.Error(`Items[4] should be 3`)
	}

	heap = MaxHeap{
		[]int{4, 3, 5, 1, 0},
	}

	//				4
	//			/		\
	//		3			5
	//	/		\
	// 1		 0
	heap.Heapify(2)

	if heap.Items[0] != 5 {
		t.Error(`Items[0] should be 5`)
	}
	if heap.Items[1] != 3 {
		t.Error(`Items[1] should be 3`)
	}
	if heap.Items[2] != 4 {
		t.Error(`Items[2] should be 4`)
	}
	if heap.Items[3] != 1 {
		t.Error(`Items[3] should be 1`)
	}
	if heap.Items[4] != 0 {
		t.Error(`Items[4] should be 0`)
	}
}

func TestInsert(t *testing.T) {
	heap := MaxHeap{
		[]int{4, 3, 2, 1, 0},
	}

	//				4
	//			/		\
	//		3			2
	//	/		\
	// 1		 0
	heap.Insert(5)
	if heap.Items[0] != 5 {
		t.Error(`Items[0] should be 5`)
	}
	if heap.Items[1] != 3 {
		t.Error(`Items[1] should be 3`)
	}
	if heap.Items[2] != 4 {
		t.Error(`Items[2] should be 4`)
	}
	if heap.Items[3] != 1 {
		t.Error(`Items[3] should be 1`)
	}
	if heap.Items[4] != 0 {
		t.Error(`Items[4] should be 0`)
	}
	if heap.Items[5] != 2 {
		t.Error(`Items[5] should be 2`)
	}
}

func TestBuildHeap(t *testing.T) {
	arr := []int{1, 16, 5, 30, 27, 17, 20, 2, 57, 3, 90}
	heap := MaxHeap{
		arr,
	}
	heap.BuildHeap()

	fmt.Println(heap.Items)
	expectedArray := []int{90, 57, 20, 30, 27, 17, 5, 2, 1, 3, 16}

	for i, v := range expectedArray {
		if heap.Items[i] != v {
			t.Errorf("Expected i[%d], v[%d], actual %d", i, v, heap.Items[i])
		}
	}
}