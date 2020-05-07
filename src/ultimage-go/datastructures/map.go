package main

import (
	"fmt"
)

type user struct {
	name		string
	surname	string
}

func main() {
	users1 := make(map[string]user)

	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	for key, value := range users1 {
		fmt.Println(key, value)
	}

	// Map literal
	users2 := map[string]user{
		"Roy": user{"Rob", "Roy"},
		"Ford": user{"Henry", "Ford"},
		"Mouse": user{"Mickey", "Mouse"},
		"Jackson": user{"Michael", "Jackson"},
	}
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// Delete key
	delete(users2, "Roy")

	// Find key
	// If ok is false, u1 will hold zero value
	u1, ok1 := users2["Roy"]
	u2, ok2 := users2["Ford"]

	fmt.Println("Roy", ok1, u1)
	fmt.Println("Ford", ok2, u2)


	// Keys of a map must be made of comparable values. Struct won't work
}