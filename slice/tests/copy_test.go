package tests2

import (
	"fmt"
	"testing"
)

func Test_Copy(t *testing.T) {
	s1 := []int{1,2,3}
	var s2 []int
	s2 = make([]int, 2)
	copy(s2, s1)
	s2[0] = 5
	fmt.Println(s1)
	fmt.Println(s2)
}