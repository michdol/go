package main

import (
    "fmt"
)

func main() {
    tree := new(Tree);
    tree.insert(1)
    fmt.Printf("%+v", tree.root)
}

type Node struct {
    data        int
    left        *Node
    right   *Node
    parent  *Node
}

type Tree struct {
    root        *Node
}

func (t *Tree) insert(data int) {
    newNode := &Node{data: data}
    if t.root == nil {
        t.root = newNode
        return
    }
    var current, parent *Node
    current = t.root
    for {
        parent = current
        if (data < parent.data) {
            current = current.left
            if current == nil {
                parent.left = newNode
                return
            }
        } else {
            current = current.right
            if current == nil {
                parent.right = newNode
                return
            }
        }
    }
}

func (t *Tree) delete(data int) bool {
    var current, parent *Node
    current = t.root
    parent = current
    for current != nil && current.data != data {
        parent = current
        if data < current.data {
            current = current.left
        } else {
            current = current.right
        }
    }
    if current == nil {
        return false
    }
    if current.left != nil && current.right != nil {
        t.switchMinimumValue(data, current)
    } else if current.left != nil && current.right == nil {
        if parent.left == current {
            parent.left = current.left
        } else {
            parent.right = current.left
        }
    } else if current.right != nil && current.left == nil {
        if parent.left == current {
            parent.left = current.right
        } else {
            parent.right = current.right
        }
    } else {
        if parent.left == current {
            parent.left = nil
        } else {
            parent.right = nil
        }
    }
    return true
}

func (t *Tree) switchMinimumValue(data int, node *Node) {
    var current, previous *Node
    previous = node.left
    current = node.left
    for current != nil {
        previous = current
        if current.left != nil {
            current = current.left
        } else {
            current = current.right
        }
    }
    node.data = current.data
    if previous.left == current {
        previous.left = nil
    } else {
        previous.right = nil
    }
}

func (t *Tree) search(data int) *Node {
    var current *Node
    current = t.root
    if current == nil {
        return nil
    }
    for {
        if current.data == data {
            return current
        }
        if data < current.data {
            current = current.left
        } else {
            current = current.right
        }
        if current == nil {
            return nil
        }
    }
}

func (t *Tree) preOrder(current *Node) {
    fmt.Println("Current: ", current.data)

    if current.left != nil {
        t.preOrder(current.left)
    }

    if current.right != nil {
        t.preOrder(current.right)
    }
}

func (t *Tree) inOrder(current *Node) {
    if current.left != nil {
        t.inOrder(current.left)
    }

    fmt.Println("Current: ", current.data)

    if current.right != nil {
        t.inOrder(current.right)
    } 
}

func (t *Tree) postOrder(current *Node) {
    if current.left != nil {
        t.postOrder(current.left)
    }

    if current.right != nil {
        t.postOrder(current.right)
    } 

    fmt.Println("Current: ", current.data)
}