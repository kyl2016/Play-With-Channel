package main

import "fmt"

func main() {
	a := []int{7, 8, 9}
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
	ap(a)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)

	ap2(&a)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
}

func ap(a []int) {
	a = append(a, 10)
}

func ap2(a *[]int) {
	*a = append(*a, 10)
}
