package main

import (
	"fmt"
	"testing"
)

func TestMap_Append(t *testing.T) {
	m := map[int]int{}
	add22(m, 1, 1)
	add22(m, 2, 2)
	fmt.Println(m)
}

func add22(m map[int]int, k, v int) {
	m[k] = v
}
