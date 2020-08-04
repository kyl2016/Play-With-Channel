package slice

import (
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
	s := make([]int, 2, 2)
	s[0] = 1

	fmt.Println("s2 = s")
	s2 := s
	fmt.Printf("s  cap: %d, value: %v\n", cap(s), s)
	fmt.Printf("s2 cap: %d, value: %v\n", cap(s2), s2)
	fmt.Println("")

	s2[1] = 2
	fmt.Println("执行 s2[1] = 2, 会影响 s")
	fmt.Printf("s  cap: %d, value: %v\n", cap(s), s)
	fmt.Printf("s2 cap: %d, value: %v\n", cap(s2), s2)
	fmt.Println("")

	fmt.Println("执行 s2 = append(s2, 2)，触发扩容，s2 内部引用的 array 与 s 已经不同")
	s2 = append(s2, 2)
	fmt.Printf("s  cap: %d, value: %v\n", cap(s), s)
	fmt.Printf("s2 cap: %d, value: %v\n", cap(s2), s2)
	fmt.Println("")

	fmt.Println("执行 s2[1] = 22，s2 与 s 没有任何关系")
	s2[1] = 22
	fmt.Printf("s  cap: %d, value: %v\n", cap(s), s)
	fmt.Printf("s2 cap: %d, value: %v\n", cap(s2), s2)
	fmt.Println("")

	fmt.Println("s3 := &s; *s3 = append(*s3, 3)")
	s3 := &s
	*s3 = append(*s3, 3)
	fmt.Printf("s  cap: %d, value: %v\n", cap(s), s)
	fmt.Printf("s3 cap: %d, value: %v\n", cap(*s3), *s3)
	fmt.Println("")

	fmt.Println("append 4,5 to s3")
	*s3 = append(*s3, 4)
	*s3 = append(*s3, 5)
	fmt.Printf("s  cap: %d, value: %v\n", cap(s), s)
	fmt.Printf("s3 cap: %d, value: %v\n", cap(*s3), *s3)

	fmt.Println("")
	fmt.Printf("s  : %p\ns2 : %p\n*s3: %p\ns3 : %p\n", s, s2, *s3, s3)
}
func TestSlice2(t *testing.T) {
	s := make([]int, 2, 4)
	s[0] = 1

	fmt.Println("s2 = s")
	s2 := s
	fmt.Printf("s  cap: %d, len: %d, value: %v\n", cap(s), len(s), s)
	fmt.Printf("s2 cap: %d, len: %d, value: %v\n", cap(s2), len(s2), s2)
	fmt.Println("")

	s2[1] = 2
	fmt.Println("执行 s2[1] = 2, 会影响 s")
	fmt.Printf("s  cap: %d, len: %d, value: %v\n", cap(s), len(s), s)
	fmt.Printf("s2 cap: %d, len: %d, value: %v\n", cap(s2), len(s2), s2)
	fmt.Println("")

	fmt.Println("执行 s2 = append(s2, 2)，未触发扩容，s2 内部引用的 array 与 s 已经不同？NO，相同！只是 len(s)=2，而 len(s2)=3")
	s2 = append(s2, 2)
	fmt.Printf("s  cap: %d, len: %d, value: %v\n", cap(s), len(s), s)
	fmt.Printf("s2 cap: %d, len: %d, value: %v\n", cap(s2), len(s2), s2)
	fmt.Println("")

	s = append(s, 9)
	fmt.Println("执行 s = append(s, 9)")
	fmt.Printf("s  cap: %d, len: %d, value: %v\n", cap(s), len(s), s)
	fmt.Printf("s2 cap: %d, len: %d, value: %v\n", cap(s2), len(s2), s2)
	fmt.Println("")

	fmt.Println("执行 s2[1] = 22，s2 与 s 同时改变了")
	s2[1] = 22
	fmt.Printf("s  cap: %d, len: %d, value: %v\n", cap(s), len(s), s)
	fmt.Printf("s2 cap: %d, len: %d, value: %v\n", cap(s2), len(s2), s2)
	fmt.Println("")

	fmt.Println("s3 := &s; *s3 = append(*s3, 3)")
	s3 := &s
	*s3 = append(*s3, 3)
	fmt.Printf("s  cap: %d, len: %d, value: %v\n", cap(s), len(s), s)
	fmt.Printf("s3 cap: %d, len: %d, value: %v\n", cap(*s3), len(*s3), *s3)
	fmt.Println("")

	fmt.Println("append 4,5 to s3")
	*s3 = append(*s3, 4)
	*s3 = append(*s3, 5)
	fmt.Printf("s  cap: %d, len: %d, value: %v\n", cap(s), len(s), s)
	fmt.Printf("s3 cap: %d, len: %d, value: %v\n", cap(*s3), len(*s3), *s3)

	fmt.Println("")
	fmt.Printf("s  : %p\ns2 : %p\n*s3: %p\ns3 : %p\n", s, s2, *s3, s3)
}
