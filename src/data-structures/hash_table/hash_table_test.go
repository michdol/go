package main

import (
	"container/list"
	"testing"
)


func TestIntegerHash(t *testing.T) {
	var ret int
	ret, _ = integerHash(2, 2)
	if ret != 0 {
		t.Error(`integerHash(2, 2) != 0`)
	}
	ret, _ = integerHash(3, 2)
	if ret != 1 {
		t.Error(`integerHash(3, 2) != 1`)
	}
	ret, _ = integerHash(5, 3)
	if ret != 2 {
		t.Error(`integerHash(5, 3) != 2`)
	}
	ret, err := integerHash(1, 0)
	if err == nil {
		t.Error(`integerHash(1, 0) should return zero division error`)
	}
}

func TestContainsKey(t *testing.T) {
	el := &list.Element{
		Value: 1,
	}
	if containsKey(2, el) != false {
		t.Error(`containsKey should return false if list does not contain duplicates`)
	}
	if containsKey(1, el) != true {
		t.Error(`containsKey should return true if list contains duplicates`)
	}
	el = nil
	if containsKey(1, el) != false {
		t.Error(`containsKey should return false if list is empty`)
	}
}

func TestPut(t *testing.T) {
	var err error
	ht := New()
	err = ht.Put(1)
	if err != nil {
		t.Error(`ht.Put(1) returned error`, err)
	}
	if ht.array[1].Back().Value != 1 {
		t.Error(`ht.array[1] != 1`)
	}
	err = ht.Put(1)
	if err != nil {
		t.Error(`Second call to ht.Put(1) returned error`)
	}
	if ht.array[1].Len() != 1 {
		t.Error(`ht.array[1].len() != 1 - array should not contain duplicates`)
	}
}
