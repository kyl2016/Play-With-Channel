package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	print(s1)

	s2 := make([]int, 5, 8)
	print(s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6] // s3 和 s4 对应相同的底层数组，s4 切片的窗口可以向右扩展，可以看到数组的末尾
	print(s4)

	s4 = append(s4, 9)

	s44 := s4[1:5]
	print(s44)

	fmt.Println(s3)

	s5 := make([]int, 100)
	s5 = append(s5, 1)
	print(s5)

}

func print(s []int) {
	fmt.Printf("The length of s: %d\n", len(s))
	fmt.Printf("The capacity of s: %d\n", cap(s))
	fmt.Printf("The value of s: %d\n", s)
	fmt.Println()
}
