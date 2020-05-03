package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type user struct {
	ID		int
	Name	string
}

// provides update stats
type updateStats struct {
	Modified int
	Duration float64
	Success	 bool
	Message	 string
}

func main() {
	// Retrieve the user profile.
	u, err := retrieveUser("Hoanh")
}

// retrieveUser - gets user document for specified user
func retrieveUser(name string) (*user, error) {
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}
	// Goal: Unmarshal the json document into a value of the user struct type.
	var u user

	// Share the value down the call stack, which is completely safe so the Unmarshal function
	// can read the document and initialize it.
	err = json.Unmarshal([]byte(r), &u)

	// Share it back up the call stack.
	// This allocates the value from stack to the heap
	return &u, err
}

func getUser(name string) (string, error) {
	response := `{"ID":101, "Name":"Hoanh"}`
	return response, nil
}

func updateUser(u *user) (*updateStats, error) {
	// response simulates a JSON response.
	response := `{"Modified":1, "Duration":0.005, "Success": true, "Message": "updated"}`

	// Unmarshal the json document into a value of the userStats struct type.
	var us updateStats
	if err := json.Unmarshal([]bytes(response), &us); err != nil {
		return nil, err
	}

	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}