package tests2

import (
	"fmt"
	"testing"
)

func TestReuse(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 1)
	fmt.Println(s)

	s = s[:0]
	s = append(s, 2)
	fmt.Println(s)
}
