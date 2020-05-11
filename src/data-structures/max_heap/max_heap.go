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

func (h *MaxHeap) Heapify(n int) {
	// Heapify starts at position 'n' and goes to the bottom if Swap was called
	parent := h.Items[n]
	hasLeft := h.HasLeftChild(n)
	hasRight := h.HasRightChild(n)
	leftIdx := h.GetLeftIndex(n)
	rightIdx := h.GetRightIndex(n)
	if hasLeft && hasRight {
		left := h.Items[leftIdx]
		right := h.Items[rightIdx]
		if left < parent && right < parent {
			return
		}
		if left > parent && left > right {
			h.Swap(n, leftIdx)
			h.Heapify(leftIdx)
		} else if right > parent && right > left {
			h.Swap(n, rightIdx)
			h.Heapify(rightIdx)
		}
	} else if hasLeft {
		left := h.Items[leftIdx]
		if left > parent {
			h.Swap(n, leftIdx)
			h.Heapify(leftIdx)
		}
	} else if hasRight {
		right := h.Items[rightIdx]
		if right > parent {
			h.Swap(n, rightIdx)
			h.Heapify(rightIdx)
		}
	}
}

func (h *MaxHeap) Insert(data int) {
	// Insert bottoms-up until reaching root or parent node is greater than current node
	h.Items = append(h.Items, data)
	var parentIdx, currentIdx int
	currentIdx = len(h.Items) - 1
	for h.HasParent(currentIdx) {
		parentIdx = h.GetParentIndex(currentIdx)
		if h.Items[currentIdx] > h.Items[parentIdx] {
			h.Swap(parentIdx, currentIdx)
			currentIdx = parentIdx
		} else {
			break
		}
	}
}

func (h *MaxHeap) BuildHeap() {
	for i := int(len(h.Items) / 2 - 1); i >= 0; i-- {
		h.Heapify(i)
	}
}

// https://fodor.org/blog/go-heap/
// https://www.youtube.com/watch?v=5iBUTMWGtIQ