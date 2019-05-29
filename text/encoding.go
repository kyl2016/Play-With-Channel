package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len(s))
	fmt.Println(utf8.ValidString(s))

	s = "A\xa1\xa1\xb0\xa1\x81\x40\x81\x80\xaa\x40\xaa\x80\xa8\x40\xa8\x80Z\x80"
	fmt.Println(s)
	fmt.Println(utf8.ValidString(s))
}
