package main

/*
#include <string.h>
void write(char* jpg) {
	memset(jpg,'c',1);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	jpg := make([]byte, 10)
	//C.write((*C.char)(unsafe.Pointer(&jpg)))
	C.write((*C.char)(unsafe.Pointer(&jpg[0])))

	fmt.Printf("%p %p\n", &jpg, &jpg[0])

	//C.write((*C.char)(C.CBytes(jpg)))

	cs := C.CString("test")
	C.write(cs)
	s := C.GoString(cs)

	r := C.GoBytes(unsafe.Pointer(&jpg), 2)

	fmt.Println(s, r)
}
