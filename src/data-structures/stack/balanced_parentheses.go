package main

import (
	"fmt"
)


func main() {
	balanced := "[]{}()[()]([{}]())"
	unbalanced := "[]{}()[()([{}]())"
	unbalanced2 := "[]}}()[()([{}]())"
	fmt.Println(isBalanced(balanced))
	fmt.Println(isBalanced(unbalanced))
	fmt.Println(isBalanced(unbalanced2))
}

var openingCharacter = map[string]bool{
	"(": true,
	"[": true,
	"{": true,
	"<": true,
}

var closingCharacter = map[string]bool{
	")": true,
	"]": true,
	"}": true,
	">": true,
}

func isBalanced(input string) bool {
	s := new(Stack)
	for _, character := range input {
		characterStr := string(character)
		if _, ok := openingCharacter[characterStr]; ok {
			s.push(1)
		}
		if _, ok := closingCharacter[characterStr]; ok {
			_, err := s.pop()
			if err != nil {
				return false
			}
		}
	}
	return s.isEmpty()
}