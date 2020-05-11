package main

import (
	"errors"
	"fmt"
)


func main() {

}

type MaxHeap struct {
	Items		[]int
}

func (h *MaxHeap) GetLeftIndex(n int) int {
	return 2 * n + 1
}

func (h *MaxHeap) GetRightIndex(n int) int {
	return 2 * n + 2
}

func (h *MaxHeap) GetParentIndex(n int) int {
	if n == 0 {
		return -1
	}
	return (n - 1) / 2
}

func (h *MaxHeap) HasLeftChild(n int) bool {
	return h.GetLeftIndex(n) < len(h.Items)
}

func (h *MaxHeap) HasRightChild(n int) bool {
	return h.GetRightIndex(n) < len(h.Items)
}

func (h *MaxHeap) HasParent(n int) bool {
	return h.GetParentIndex(n) >= 0 && len(h.Items) >= n + 1
}

func (h *MaxHeap) Left(n int) (int, error) {
	if h.HasLeftChild(n) {
		return h.Items[h.GetLeftIndex(n)], nil
	}
	return -1, errors.New(fmt.Sprintf("Node at index %d has not left child", n))
}

func (h *MaxHeap) Right(n int) (int, error) {
	if h.HasRightChild(n) {
		return h.Items[h.GetRightIndex(n)], nil
	}
	return -1, errors.New(fmt.Sprintf("Node at index %d has not right child", n))
}

func (h *MaxHeap) Parent(n int) (int, error) {
	if h.HasParent(n) {
		return h.Items[h.GetParentIndex(n)], nil
	}
	return -1, errors.New(fmt.Sprintf("Node at index %d has not parent", n))
}

func (h *MaxHeap) Swap(i int, j int) error {
	lastIdx := len(h.Items) - 1
	if i > lastIdx || j > lastIdx || j < 0 || i < 0 {
		return errors.New("Index out of range")
	}
	left := h.Items[i]
	right := h.Items[j]
	h.Items[i] = right
	h.Items[j] = left
	return nil
}

func (h *MaxHeap) MaxHeapify(n int) {
	if !h.HasParent(n) {
		return
	}
	parentIdx := h.GetParentIndex(n)
	parent := h.Items[parentIdx]
	element := h.Items[n]
	if element > parent {
		h.Swap(n, parentIdx)
		h.Heapify(parentIdx)
	}
}

func (h *MaxHeap) Heapify(n int) {
	var left, right, parent int
	var leftIdx, rightIdx int
	parent = h.Items[n]
	fmt.Printf("Enter parent: %d\n", parent)
	leftIdx = h.GetLeftIndex(n)
	rightIdx = h.GetRightIndex(n)
	hasLeft := h.HasLeftChild(n)
	hasRight := h.HasRightChild(n)
	if hasLeft && hasRight {
		left = h.Items[leftIdx]
		right = h.Items[rightIdx]
		fmt.Printf("Parent: %d; Left: %d; Right %d; PIdx: %d\n", parent, left, right, n)
		if left > parent && right > parent {
			if left > right {
				fmt.Printf("Swapping left %d with %d\n", left, parent)
				h.Swap(n, leftIdx)
			} else {
				fmt.Printf("Swapping right %d with %d\n", right, parent)
				h.Swap(n, rightIdx)
			}
		} else if left > parent && right <= parent {
			fmt.Printf("Swapping left %d with %d\n", left, parent)
			h.Swap(n, leftIdx)
		} else {
			fmt.Printf("Swapping right %d with %d\n", right, parent)
			h.Swap(n, rightIdx)
		}
	} else if hasLeft {
		left = h.Items[leftIdx]
		if left > parent {
			fmt.Printf("Swapping left %d with %d\n", left, parent)
			h.Swap(n, leftIdx)
		}
	} else if hasRight {
		right = h.Items[rightIdx]
		if right > parent {
			fmt.Printf("Swapping right %d with %d\n", right, parent)
			h.Swap(n, rightIdx)
		}
	}
	/*
	if h.HasParent(n) {
		fmt.Printf("Moving to parent %d\n", h.Items[h.GetParentIndex(n)])
		h.Heapify(h.GetParentIndex(n))
	}*/
}

func (h *MaxHeap) Insert(data int) {
	h.Items = append(h.Items, data)
	lastIdx := len(h.Items) - 1
	h.Heapify(lastIdx)
}

func (h *MaxHeap) BuildHeap() {
	for i := int(len(h.Items) / 2 - 1); i >= 0; i-- {
		fmt.Println(i)
		h.Heapify(i)
	}
}

// https://fodor.org/blog/go-heap/
// https://www.youtube.com/watch?v=5iBUTMWGtIQ