package main

import "fmt"

func main() {
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Before [%s] : ", five[1])

	for i := range five {
		five[1] = "Jack"

		if i == 1 {
			fmt.Printf("After [%s]\n", five[1])
		}
	}

	five = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Before [%s] : ", five[1])

	for i, v := range five {
		five[1] = "Jack"

		if i == 1 {
			fmt.Printf("After [%s]\n", v)
		}
	}
}
