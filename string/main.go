package main

import (
	"bytes"
)

func main() {
	bytes1 := []byte("5c3ed3444844b428a65dc0b6")
	println(len(bytes1))

	bytes1 = []byte("gc3ed3444844b428a65dc0b6")
	println(len(bytes1))

	s := "Go语言"
	println(s[2:5])
	println(s[2:6])

	var buffer bytes.Buffer

}
