package main

import (
	"io/ioutil"
	"os"
	"syscall"
)

func main() {
	file := "/home/lynxi/Documents/faces/6万人/6万人.zip"
	in, err := OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	buffer, err := ioutil.ReadAll(in)
	println(len(buffer))

	fi, err := in.Stat()
	if err != nil {
		panic(err)
	}

	println(fi.Name())
}

func OpenFileSkipPageCache(name string, flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(name, syscall.O_DIRECT|flag, perm)
}

func OpenFile(name string, flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(name, flag, perm)
}
