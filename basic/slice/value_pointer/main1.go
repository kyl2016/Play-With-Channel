package main

import "fmt"

func main() {
	s1 := make([]byte, 2)
	fmt.Println("s1", len(s1), s1)
	write(s1)
	fmt.Println("update s1[0]")
	fmt.Println("s1", len(s1), s1)
	fmt.Println("")

	s2 := make([]byte, 2)
	fmt.Println("s2", len(s2), s2)
	appendToValue(s2)
	fmt.Println("append to s2")
	fmt.Println("s2", len(s2), s2)
	fmt.Println("")

	s3 := make([]byte, 2)
	fmt.Println("s3", len(s3), s3)
	appendToPointer(&s3)
	fmt.Println("append to &s3")
	fmt.Println("s3", len(s3), s3)
	fmt.Println("")
}

func write(s []byte) {
	s[0] = 'a'
}

func appendToValue(s []byte) {
	s = append(s, 10)
	fmt.Println("------------ appendToValue", len(s), s)
}

func appendToPointer(s *[]byte) {
	*s = append(*s, 10)
	fmt.Println("------------ appendToPointer", len(*s), *s)
}
