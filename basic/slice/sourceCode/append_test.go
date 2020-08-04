package sourceCode

import (
	"fmt"
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

// cap > doublecap
func Test1(t *testing.T) {
	println(unix.SizeofInt)
	println(4 << (^uintptr(0) >> 63))

	//cap := 5
	newcap := 5
	//lenmem := uintptr(2) * 4
	//newlenmem := uintptr(cap) * 4
	capmem := uintptr(newcap) * 4
	//overflow = uintptr(newcap) > maxSliceCap(4)
	newcap = int(capmem / 4)

	si := unsafe.Sizeof(4)
	println(si)
	println(isPowerOfTwo(si))

	s := []int{1, 2}
	print(s)
	s = append(s, []int{3}...)
	print(s)

	s = []int{1, 2}
	s = append(s, 3, 4, 5)
	//s = append(s, []int{3, 4}...)
	//s = append(s, 5)
	print(s)

	s = []int{1, 2}
	s = append(s, []int{3, 4, 5}...)
	print(s)

	s = []int{1, 2}
	s = append(s, []int{3, 4, 5, 6}...)
	print(s)

	s = []int{1, 2}
	s = append(s, []int{3, 4, 5, 6, 7}...)
	print(s)

	s = []int{1, 2}
	s = append(s, []int{3, 4, 5, 6, 7, 8}...)
	print(s)

	s = []int{1, 2}
	s = append(s, []int{3, 4, 5, 6, 7, 8, 9}...)
	print(s)
}

func isPowerOfTwo(x uintptr) bool {
	return x&(x-1) == 0
}

func print(s []int) {
	fmt.Printf("len:%d, cap:%d, value:%v\n", len(s), cap(s), s)
}
