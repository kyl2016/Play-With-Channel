package main

import "fmt"

// Here is an example that is a complete horror show.

type User struct {
	Name  string
	Likes int
}

func (u *User) Notify() {
	fmt.Printf("%s has %d likes\n", u.Name, u.Likes)
}

func (u *User) AddLike() {
	u.Likes++
}

func main() {
	users := []User{
		{Name: "bill"},
		{Name: "lisa"},
	}

	for _, u := range users {
		u.AddLike() // u is the copy value, not affect the original user.
	}

	for _, u := range users {
		u.Notify()
	}

	for i, _ := range users {
		users[i].AddLike()
	}

	for _, u := range users {
		u.Notify()
	}
}
