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

// cap 计算方法
func Test3(t *testing.T) {
	s1 := make([]int, 5)
	s2 := s1[2:]
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2) // cap(s2) = cap(s1) - 2
}

func TestCopy1(t *testing.T) {
	s1 := make([]int, 5)
	var s2 []int
	c := copy(s2, s1)
	fmt.Println(len(s2), cap(s2), s2, c) // s2 的长度是 0，无法 copy
}

func TestCopy2(t *testing.T) {
	s1 := make([]int, 5)
	var s2 = make([]int, 6)
	s1[0] = 1
	s1[1] = 2
	c := copy(s2, s1)
	s1[2] = 3
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)
	fmt.Println(c) // Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
}

func TestRemove(t *testing.T) {
	s := make([]int, 5)
	s[0] = 1
	fmt.Println(len(s), cap(s))
	s = s[:0]
	fmt.Println(len(s), cap(s))
	//5 5
	//0 5
	//fmt.Println(s[0]) panic: runtime error: index out of range [0] with length 0 [recovered]
}

func TestRemove2(t *testing.T) {
	s := make([]int, 5)
	s[0] = 1
	fmt.Println(len(s), cap(s))
	s = []int{}
	fmt.Println(len(s), cap(s))
	//5 5
	//0 0
	//fmt.Println(s[0]) panic: runtime error: index out of range [0] with length 0 [recovered]
}
