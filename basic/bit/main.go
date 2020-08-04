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

	fmt.Printf("a<<1=%d\n", a<<1)
	fmt.Printf("b=%d b>>1=%d\n", b, b>>1)
	fmt.Printf("a>>1=%d\n", a>>1)

	c := -1
	fmt.Printf("c=%d c>>1=%d\n", c, c>>1)
	fmt.Printf("c=%d c<<1=%d\n", c, c<<1)
}
