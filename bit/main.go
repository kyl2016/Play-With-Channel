package main

import "fmt"

func main() {
	a := 1
	b := 2

	fmt.Printf("a=%d, b=%d\n", a, b)
	fmt.Printf("a&b=%d\n", a&b)
	fmt.Printf("a|b=%d\n", a|b)
	fmt.Printf("a^b=%d\n", a^b)

	b = 1

	fmt.Printf("a&b=%d\n", a&b)
	fmt.Printf("a|b=%d\n", a|b)
	fmt.Printf("a^b=%d\n", a^b)
}
