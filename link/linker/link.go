package linker

import _ "unsafe"

func Test(){
	if clean2 != nil {
		clean2()
	}
}

var clean2 func()

//go:linkname register github.com/kyl2016/Play-With-Golang/link/business.register
func register(clean func()) {
	clean2 = clean
}
