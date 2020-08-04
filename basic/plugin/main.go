package main

import (
	"fmt"
	"plugin"

	"github.com/kyl2016/Play-With-Golang/basic/plugin/dataStruct"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main")

	p, err := plugin.Open("pkg.so")
	if err != nil {
		panic(err)
	}

	fmt.Println("opened")

	f, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	fmt.Println("Loopup")

	fmt.Println(p)

	f.(func(s string))("test")

	update, err := p.Lookup("Update")
	if err != nil {
		panic(err)
	}

	person := dataStruct.Person{ID: 1, Name: "Kitty", Age: 12}
	update.(func(_p *dataStruct.Person))(&person)
	fmt.Println(person.Age)
}

// go run cpu.go
