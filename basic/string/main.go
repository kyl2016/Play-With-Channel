package main

import (
	"fmt"
)

func main() {
	bytes1 := []byte("5c3ed3444844b428a65dc0b6")
	println(len(bytes1))

	buffer2 := []byte("5c3ed3444844b428a65dc0b61")
	if string(bytes1) == string(buffer2) && string(bytes1) == string(buffer2) {
		fmt.Println("equal")		
	}

	// bytes1 = []byte("gc3ed3444844b428a65dc0b6")
	// println(len(bytes1))

	// s := "Go语言"
	// println(s[2:5])
	// println(s[2:6])
}
