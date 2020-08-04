package main

import (
	"fmt"
	"github.com/kyl2016/Play-With-Golang/basic/plugin/dataStruct"
)

func init() {
	fmt.Println("world")
}

func Hello(s string) {
	fmt.Println("hello", s)
}

func Print(p *dataStruct.Person) {
	fmt.Println(p.Name)
}

func Update(p *dataStruct.Person) {
	p.Age++
}

// go build -buildmode=plugin -o ./pkg.so
