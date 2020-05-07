package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := new(Tree)
	if tree.root != nil {
		t.Error(`root is not nil`)
	}
	tree.insert(1)
	if tree.root.data != 1 {
		t.Error(`tree.root.data is not 1`)
	}

	tree.insert(2)
	if tree.root.right == nil {
		t.Error(`tree.root.right is not a node`)
	}
	if tree.root.right.data != 2 {
		t.Error(`tree.root.right.data is not 2`)
	}

	tree.insert(0)
	if tree.root.left == nil {
		t.Error(`tree.root.left is not a node`)
	}
	if tree.root.left.data != 0 {
		t.Error(`tree.root.left.data is not 0`)
	}

	tree.insert(4)
	if tree.root.right.right == nil {
		t.Error(`tree.root.right.right is not a node`)
	}
	if tree.root.right.right.data != 4 {
		t.Error(`tree.root.right.right.data is not 4`)
	}
}

func TestSearch(t *testing.T) {
	tree := createTestingTree()
	rootRef := tree.root

	if tree.search(4) != rootRef {
		t.Error(`root's data should be 4`)
	}

	var expectedRef *Node
	expectedRef = tree.root.right
	if tree.search(5) != expectedRef {
		t.Error(`failed to find 5`)
	}
	if tree.search(1) != nil {
		t.Error(`found node when should return nil`)
	}

	expectedRef = tree.root.left
	if tree.search(3) != expectedRef {
		t.Error(`failed to find 3`)
	}
	expectedRef = tree.root.left.left
	if tree.search(2) != expectedRef {
		t.Error(`failed to find 2`)
	}
}

func TestPreOrder(t *testing.T) {
	fmt.Println("TestPreOrder")
	tree := new(Tree)
	tree.insert(4)
	tree.insert(3)
	tree.insert(1)
	tree.insert(5)
	tree.insert(2)
	//tree.preOrder(tree.root)
}

func TestInOrder(t *testing.T) {
	fmt.Println("TestInOrder")
	tree := new(Tree)
	tree.insert(4)
	tree.insert(3)
	tree.insert(1)
	tree.insert(5)
	tree.insert(2)
	//tree.inOrder(tree.root)
}

func TestPostOrder(t *testing.T) {
	/*
				4
			/  \
		 3    5
		/ \
	 1   2
	*/
	fmt.Println("TestPostOrder")
	tree := new(Tree)
	tree.insert(4)
	tree.insert(3)
	tree.insert(1)
	tree.insert(5)
	tree.insert(2)
	//tree.postOrder(tree.root)
}

func TestDeleteNoChildren(t *testing.T) {
	tree := createTestingTree()

	ret := tree.delete(5)
	if ret != true {
		t.Error(`delete operation did not return true`)
	}
	if tree.root.right != nil {
		t.Error(`tree.root.right is not empty`)
	}
}

func TestDeleteOneChild(t *testing.T) {
	tree := createTestingTree()

	ret := tree.delete(3)
	if ret != true {
		t.Error(`delete operation did not return true`)
	}
	if tree.root.left.data != 2 {
		t.Error(`root's left should be 2`)
	}
}

func TestDeleteTwoChildren(t *testing.T) {
	tree := createTestingTree()

	ret := tree.delete(4)
	if ret != true {
		t.Error(`delete operation did not return true`)
	}
	if tree.root.data != 2 {
		t.Error(`root's left should be 2`)
	}
}

func createTestingTree() *Tree {
	tree := new(Tree)
	tree.insert(4)
	tree.insert(3)
	tree.insert(5)
	tree.insert(2)
	return tree
}
/*
			4
		/  \
	 3    5
	/
 2
*/