package main

import "fmt"

type user struct {
	name string
	email string
}


// Using the value receiver, the method operates on
// its own copy of the value that is used to make the call.
func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}

// Using the pointed receiver, method operates on shared access.
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("bill@hotmail.com")

	michal := &user{"Michal", "michal@email.com"}
	michal.notify()
	michal.changeEmail("michal@hotmail.com")

	// Calling notify on michal (*user) is still correct.
	// As long as we deal with the type user, Go can adjust to make the call.
	
	// Behind the scenes we have something like (*michal).notify()

	// Similarily calling changeEmail on bill (user), Go will take address of bill
	// &bill.changeEmail(email)

	// Slice of users
	users := []user{
		{"bill", "bill@email"},
		{"michal", "michal@email"},
	}

	// Ranging over this slice of values, makes a copy of each value and call to notify
	// makes another copy.
	for _, u := range users {
		u.notify()
	}

	// Bad practice
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	// Email did not change
	for _, u := range users {
		u.notify()
	}
}