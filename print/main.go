package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Println(os.Stdout, "Hello ", 23, "\n")
	fmt.Println(os.Stderr, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	type T struct {
		a int
		b float64
		c string
	}
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("v : %v\n", t)
	fmt.Printf("+v: %+v\n", t)
	fmt.Printf("#v: %#v\n", t)
	fmt.Printf("T : %T\n", t)
	fmt.Printf("q : %q\n", "abc")
	fmt.Printf("dgq:%d%g%q\n", t.a, t.b, t.c)

	fmt.Printf("%s\n", []byte("abc"))

	fmt.Printf("%v\n", time.Now())
net.Error()r
	fmt.Println(fmt.Errorf("error info is %s", "id must be int type"))
}
