package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	fileName := "new.md"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	println(file.Name())

	file2, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	println(file2.Name())

	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}

	file3, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_EXCL, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file3.Close()
	println(file3.Name())

	n, err := file3.WriteString("123")
	if err != nil {
		println(err.Error())
	}

	file4, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file4.Close()
	n, err  = file4.WriteString("456")
	if err != nil {
		panic(err)
	}
	println("WriteString ", strconv.Itoa(n))

	buffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	println(string(buffer))

	file5, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, os.ModePerm)
	n, err = file5.WriteString("789")
	if err != nil {
		panic(err)
	}
	defer file5.Close()
	println("WriteString", strconv.Itoa(n))

	buffer, err = ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	println(string(buffer))
}
