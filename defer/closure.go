package main

import "fmt"

func main() {
	var whatever [3]struct{}

	for i, _ := range whatever {
		fmt.Printf("index %d pointer is %p\n", i, &i)
		defer func() {
			fmt.Println(i)
		}()
	}
}
