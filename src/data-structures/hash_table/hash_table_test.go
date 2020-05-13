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

func TestListContainsKey(t *testing.T) {
	el := &list.Element{
		Value: 1,
	}
	if listContainsKey(2, el) != false {
		t.Error(`listContainsKey should return false if list does not contain duplicates`)
	}
	if listContainsKey(1, el) != true {
		t.Error(`listContainsKey should return true if list contains duplicates`)
	}
	el = nil
	if listContainsKey(1, el) != false {
		t.Error(`listContainsKey should return false if list is empty`)
	}
}

func TestGetValueFromList(t *testing.T) {
	var val interface{}
	var err error
	el := &list.Element{
		Value: 1,
	}
	val, err = getValueFromList(1, el)
	if val != 1 {
		t.Errorf(`getValueFromList should return 1, returned %+v instead`, val)
	}
	if err != nil {
		t.Errorf(`getValueFromList returned error %+v`, err)
	}
	val, err = getValueFromList(2, el)
	if val != nil {
		t.Errorf(`getValueFromList should return nil, returned %+v instead`, val)
	}
	if err == nil {
		t.Errorf(`getValueFromList returned nil instead of error`)
	}
}

func TestPutInt(t *testing.T) {
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

func TestGetInt(t *testing.T) {
	var val interface{}
	var err error
	ht := New()
	ht.Put(1)

	val, err = ht.Get(1)
	if err != nil {
		t.Errorf(`Get(1) returned error %+v`, err)
	}
	if val != 1 {
		t.Errorf(`Get(1) should return 1, returned %+v instead`, val)
	}
	val, err = ht.Get(2)
	if err == nil {
		t.Errorf(`Get(2) should return error, returned nil instead`)
	}
	if val != nil {
		t.Errorf(`Get(2) should return nil, returned %+v instead`, val)
	}
}

func TestPutString(t *testing.T) {
	var err error
	ht := New()
	// "a" = 49
	// len(ht.array) = 10
	// 49 % 10 = 9
	err = ht.Put("a")

	if err != nil {
		t.Errorf(`Put("a") should not return error;\n%+v`, err)
	}
	if ht.array[9] == nil && ht.array[9].Front().Value != "a" {
		t.Errorf(`Array should have "a" stored at index 9`)
	}

	// Collision
	// "ab" = 49 + 50 = 99
	// len(ht.array) = 10
	// 99 % 10 = 9
	err = ht.Put("ab")
	if err != nil {
		t.Errorf(`Put("ab") should not return error;\n%+v`, err)
	}
	length := ht.array[9].Len()
	if length != 2 {
		t.Errorf(`List at index 9 should be 2 elements long; actual: %d`, length)
	}
	val := ht.array[9].Back().Value
	if val != "ab" {
		t.Errorf(`Element should be "ab"; actual: %+v`, val)
	}
}

func TestGetString(t *testing.T) {
	var val interface{}
	var err error
	ht := New()
	ht.Put("a")
	ht.Put("ab")

	val, err = ht.Get("a")
	if err != nil {
		t.Errorf(`Get("a") returned error when should return nil;\n%+v`, err)
	}
	if val != "a" {
		t.Errorf(`fuck`)
	}
}