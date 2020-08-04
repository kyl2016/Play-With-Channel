package main

/*
#include "adder.h"
*/
import "C"

import "fmt"

// The //export allows C to call the Go func
//export GiveGoTotal
func GiveGoTotal(total C.int) {
	fmt.Printf("Go: go total from C %d\n", total)
}

func main() {
	fmt.Printf("Go: calling C to add numbers... \n")
	C.add_and_give_go_total(30, 2)
}
