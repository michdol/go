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
        if current.data > data {
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