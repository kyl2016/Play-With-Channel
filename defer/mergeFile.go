package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	folder := "/home/lynxi/tmp/"

	f, _ := os.Open(folder + "file1.txt")
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file1.txt err %v\n", err)
			} else {
				fmt.Println("close file1.txt success")
			}
		}(f)
	}

	f, _ = os.Open(folder + "file2.txt")
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file2.txt err %v\n", err)
			} else {
				fmt.Println("close file2.txt success")
			}
		}(f)
	}
}
