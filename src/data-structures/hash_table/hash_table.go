package main

import (
	"container/list"
	"errors"
	"fmt"
)


func main() {
	ht := New()
	ht.Put("a", 1)
	fmt.Printf("%+v\n", ht.array[9].Front().Value)
	v, _ := stringHash("ab", 10)
	fmt.Println("string hash", v)
}

type HashTable struct {
	array	[]*list.List
}

func New() *HashTable {
	return &HashTable{
		make([]*list.List, 10),
	}
}

func (ht *HashTable) Put(key interface{}, value interface{}) error {
	var linkedList *list.List
	hash, err := ht.HashFunction(key)
	if err != nil {
		return err
	}
	length := len(ht.array)
	if length < hash {
		return errors.New("TODO: implement extending array")
	}

	linkedList = ht.array[hash]
	if linkedList == nil {
		linkedList = list.New()
		ht.array[hash] = linkedList
	}
	newEntry := [2]interface{}{key, value}
	entry := getElementWithTheKey(key, linkedList.Front())
	// Override entry
	if entry != nil {
		linkedList.Remove(entry)
	}
	linkedList.PushBack(newEntry)
	return nil
}

func (ht *HashTable) Get(key interface{}) (interface{}, error) {
	var linkedList *list.List
	hash, err := ht.HashFunction(key)
	if err != nil {
		return nil, err
	}
	length := len(ht.array)
	if length < hash {
		return nil, errors.New("TODO: implement extending array")
	}

	linkedList = ht.array[hash]
	if linkedList == nil {
		return nil, errors.New(fmt.Sprintf("List not present at index %d", hash))
	}
	return getValueFromList(key, linkedList.Front())
}

func (ht *HashTable) HashFunction(key interface{}) (int, error) {
	length := len(ht.array)
	if data, ok := key.(string); ok {
		return stringHash(data, length)
	} else if data, ok := key.(int); ok {
		return integerHash(data, length)
	} else {
		return -1, errors.New("Unknown type")
	}
}

func listContainsKey(key interface{}, element *list.Element) bool {
	for element != nil {
		elementKey := getElementOfArray(element.Value, 0)
		if elementKey == key {
			return true
		}
		element = element.Next()
	}
	return false
}

func getElementWithTheKey(key interface{}, element *list.Element) *list.Element {
	for element != nil {
		elementKey := getElementOfArray(element.Value, 0)
		if elementKey == key {
			return element
		}
		element = element.Next()
	}
	return nil
}

func getElementOfArray(entry interface{}, index int) interface{} {
	if ent, ok := entry.([2]interface{}); ok {
		return ent[index]
	}
	return nil
}

func getValueFromList(key interface{}, element *list.Element) (interface{}, error) {
	entry := getElementWithTheKey(key, element)
	if entry == nil {
		return nil, errors.New("Element not present in the list")
	}
	ret := getElementOfArray(entry.Value, 1)
	return ret, nil
}

func integerHash(key int, arrayLength int) (int, error) {
	if (arrayLength == 0) {
		return -1, errors.New("Zero division error")
	}
	return key % arrayLength, nil
}

func stringHash(key string, arrayLength int) (int, error) {
	if len(key) == 0 {
		return -1, errors.New("Key cannot be empty string")
	}
	sum := 0
	for _, char := range key {
		sum += int(char - '0')
	}
	return sum % arrayLength, nil
}