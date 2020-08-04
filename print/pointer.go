package main

import "fmt"

func main() {
	s := make([]byte, 3)
	fmt.Printf("%p\n %p\n %p\n %p\n", &s, &s[0], &s[1], &s[2])
}
