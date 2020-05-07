package main

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := new(Stack)

	stack.push(0)
	if stack.arr[0] != 0 {
		t.Error(`Stack's first item should be 0`)
	}

	stack.push(1)
	stack.push(2)
	stack.push(3)
	if stack.arr[1] != 1 {
		t.Error(`Stack's second item should be 1`)
	}
	if stack.arr[2] != 2 {
		t.Error(`Stack's third item should be 2`)
	}
	if stack.arr[3] != 3 {
		t.Error(`Stack's fourth item should be 3`)
	}
}

func TestStackPop(t *testing.T) {
	stack := new(Stack)
	stack.push(0)

	ret, err := stack.pop()
	if ret != 0 {
		t.Error(`Stack's pop should return 0`)
	}
	if err != nil {
		t.Error(`Pop should not return error`)
	}

	ret, err = stack.pop()
	if ret != -1 {
		t.Error(`Returned value should be -1`)
	}
	if err == nil {
		t.Error(`Pop should return error when popping from empty stack`)
	}
}

func TestStackTop(t *testing.T) {
	stack := new(Stack)

	ret, err := stack.top()
	if ret != -1 {
		t.Error(`Returned value should be -1`)
	}
	if err == nil {
		t.Error(`Should return error`)
	}

	stack.push(1)

	ret, err = stack.top()
	if ret != 1 {
		t.Error(`Stack's top should return 1`)
	}
	if err != nil {
		t.Error(`Pop should not return error`)
	}
	if len(stack.arr) != 1 && stack.arr[0] != 1 {
		t.Error(`Stack's array should not change after top operation`)
	}
}

func TestStackLen(t *testing.T) {
	stack := new(Stack)

	if stack.len() != 0 {
		t.Error(`Stack len operation on empty stack should return 0`)
	}
	stack.push(1)
	if stack.len() != 1 {
		t.Error(`Stack len operation on stack with 1 element should return 1`)
	}
	stack.push(1)
	if stack.len() != 2 {
		t.Error(`Stack len operation on stack with 2 elements should return 2`)
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := new(Stack)

	if stack.isEmpty() != true {
		t.Error(`isEmpty should return true on stack with no elements`)
	}
	stack.push(1)
	if stack.isEmpty() != false {
		t.Error(`isEmpty should return false on stack with 1 element`)
	}
}