package main

import "fmt"

func main() {
	s2 := make([]byte, 2)
	fmt.Printf("&s2[0]=%p, len=%d, cap=%d\n", &s2[0], len(s2), cap(s2))

	appendToValue2(s2)

	fmt.Printf("&s2[0]=%p, len=%d, cap=%d\n", &s2[0], len(s2), cap(s2))
	fmt.Println("")
}

func appendToValue2(s []byte) {
	fmt.Println("------------ appendToValue2...")
	fmt.Printf("------------ s 对应数组的首地址：%p\n", &s[0])
	s = append(s, 10)
	fmt.Println("------------ appendToValue", len(s), s)
	fmt.Printf("------------ s 对应数组的首地址：%p\n", &s[0])
}
