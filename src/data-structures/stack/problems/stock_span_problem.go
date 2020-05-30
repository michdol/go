package main

import (
	"fmt"
)

func main() {
	s := new(Stack)
	input := []int{100, 80, 60, 70, 60, 75, 85}
	output := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		currentValue := input[i]
		for j := 0; j < i; j++ {
			previousValue := input[j]
			if previousValue <= currentValue {
				s.push(1)
			} else {
				s.clear()
			}
		}
		output[i] = s.length + 1
		s.clear()
	}

	fmt.Println(output)
}

/*
Traverse the input price array.
For every element being visited, traverse elements on left of it and increment
the span value of it while elements on the left side are smaller.
*/