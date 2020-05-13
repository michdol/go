package main

import (
	"container/list"
	"errors"
	"fmt"
)


func main() {

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
	if !containsKey(key, linkedList.Front()) {
		ht.array[hash].PushBack(key)
	}
	return nil
}

func (ht *HashTable) HashFunction(key interface{}) (int, error) {
	var hash int
	length := len(ht.array)
	if data, ok := key.(string); ok {
		hash = stringHash(data, length)
		fmt.Println(hash)
		// TODO:
		return -1, nil
	} else if data, ok := key.(int); ok {
		return integerHash(data, length)
	} else {
		return -1, errors.New("Unknown type")
	}
}

func containsKey(key interface{}, element *list.Element) bool {
	for element != nil {
		if element.Value == key {
			return true
		}
		element = element.Next()
	}
	return false
}

func integerHash(key int, arrayLength int) (int, error) {
	if (arrayLength == 0) {
		return -1, errors.New("Zero division error")
	}
	return key % arrayLength, nil
}

func stringHash(key string, arrayLength int) int {
	sum := 0
	for _, char := range key {
		fmt.Println(char, sum)
	}
	return 1
}