package main

import "fmt"

func main() {
	s := make([]int, 5)
	s[0] = 1
	fmt.Println(s)

	s = append(s, 2, 3, 4, 5)
	fmt.Println(len(s), s)

	s2 := s[:0]
	fmt.Println(s2)
}
