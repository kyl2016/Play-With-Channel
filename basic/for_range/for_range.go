package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	users := []User{
		{"Bill", "bill@email.com"},
		{"Lisa", "lisa@email.com"},
		{"Paul", "paul@email.com"},
	}

	// Two times copy
	for i, u := range users { // copy user value to u
		fmt.Println(i, u) // Println creates a second copy
	}

	// The code inside the loop is no longer operating on its own copy, instead it is operating on the origin user value stored inside the slice.
	for i := range users {
		fmt.Println(i, users[i]) // The call to Println is still using value semantics and is being passed a copy.
	}

	// No value copy
	for i := range users {
		fmt.Println(i, &users[i])
	}
}
