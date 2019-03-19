package main

import "fmt"

func main() {
	// slice
	fmt.Println("\nslice samples\n")

	slice := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for _, v := range slice {
		slice = slice[:2]

		fmt.Printf("v [%s]\n", v)
	}

	slice = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	// The for range took the length of the slice before iterating, but during the loop that length changed. Now on the third iteration, the loop attempts to access an element that is no longer associated with the slice's length.
	for i := range slice {
		slice = slice[:2]

		fmt.Printf("%d [%s]\n", i, slice[i]) // panic
	}
}
