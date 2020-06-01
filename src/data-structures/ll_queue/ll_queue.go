package main


type Node struct {
	data			interface{}
	next			*Node
}

func (n *Node) HasNext() bool {
	return n.next != nil
}

type LLQueue struct {
	first		*Node
	last		*Node
	size		int
}

func (llq *LLQueue) IsEmpty() bool {
	return llq.first == nil
}

func (llq *LLQueue) Size() int {
	return llq.size
}

func (llq *LLQueue) Enqueue(data interface{}) *Node {
	newNode := &Node{
		data,
		nil,
	}
	if llq.first == nil {
		llq.first = newNode
	}
	if llq.last != nil {
		llq.last.next = newNode
	}
	llq.last = newNode
	llq.size += 1
	// return
	return newNode
}

func (llq *LLQueue) Dequeue() interface{} {
	// Get first node if exists
	if llq.IsEmpty() {
		return nil
	}
	first := llq.first
	if first.HasNext() {
		llq.first = first.next
	}
	llq.size -= 1
	if first == llq.last {
		llq.first = nil
		llq.last = nil
	}
	return first.data
}