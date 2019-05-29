package main

import "fmt"

func main() {
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	person := "Ann"

	if attended[person] {
		fmt.Println(person, "was at the meeting")
	}
}
