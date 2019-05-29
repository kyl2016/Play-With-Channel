package business

import (
	//_ "github.com/kyl2016/Play-With-Golang/link/linker"
	_ "unsafe"
)

func init(){
	register(clean)
}

func register(clean func())

func clean() {
	println("clean")
}
