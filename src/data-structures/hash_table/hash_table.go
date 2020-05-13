package main

import (
	"container/list"
	"errors"
	"fmt"
)


func main() {
	ht := New()
	ht.Put("a")
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

func (ht *HashTable) Put(key interface{}) error {
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
	if !listContainsKey(key, linkedList.Front()) {
		ht.array[hash].PushBack(key)
	}
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
		if element.Value == key {
			return true
		}
		element = element.Next()
	}
	return false
}

func getValueFromList(key interface{}, element *list.Element) (interface{}, error) {
	for element != nil {
		if element.Value == key {
			return element.Value, nil
		}
		element = element.Next()
	}
	return nil, errors.New("Element not present in the list")
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