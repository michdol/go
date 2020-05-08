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