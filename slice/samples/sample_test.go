package samples

import (
	"fmt"
	"testing"
)

// s 和 s2 引用相同的数组
func Test1(t *testing.T) {
	s := make([]int, 2)
	fmt.Println(len(s), cap(s), s)
	s2 := s[:1]
	fmt.Println(len(s2), cap(s2), s2)
	s[0] = 1
	fmt.Println(len(s), cap(s), s)
	fmt.Println(len(s2), cap(s2), s2)
	//2 2 [0 0]
	//1 2 [0]
	//2 2 [1 0]
	//1 2 [1]
}

// s1 扩容，不影响 s2 引用的数组
func Test2(t *testing.T) {
	s1 := make([]int, 2)
	s2 := s1[:1]
	s1 = append(s1, 3)
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)
	//3 4 [0 0 3]
	//1 2 [0]
}
