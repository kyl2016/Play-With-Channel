package main

import (
	"fmt"
)

func main() {
	var s []int

	set(s)
	fmt.Printf("%v\n", s)

	set1(s)
	fmt.Printf("%v\n", s)

	set2(&s)
	fmt.Printf("%v\n", s)

	s1 := make([]int, 2)

	set(s1)
	fmt.Printf("%v\n", s1)

	set1(s1)
	fmt.Printf("%v\n", s1)

	clear(s1)
	fmt.Printf("%v\n", s1)

	clear2(&s1)
	fmt.Printf("%v\n", s1)
}

func set(s []int) {
	s = []int{}
	s = append(s, 0)
	fmt.Printf("inner %v\n", s)
}

func set1(s []int) {
	s = append(s, 1)
	s[0] = 1
	fmt.Printf("inner %v\n", s)
}

func set2(s *[]int) {
	*s = append(*s, 2)
	fmt.Printf("inner %v\n", s)
}

func clear(s []int) {
	s = s[:0]
}

func clear2(s *[]int) {
	*s = (*s)[:0]
}
