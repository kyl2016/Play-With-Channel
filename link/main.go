package main

import (
	"github.com/kyl2016/Play-With-Golang/link/linker"
	_ "github.com/kyl2016/Play-With-Golang/link/business"
)

func main() {
	println("begin")

	linker.Test()

	println("end")
}
