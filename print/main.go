package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	fs := []float64{9.90, 21.1, 12.23423423, 234.23423423, 42332.12312332432213, 4345225442542452332.12312332432213156988992}

	for _, f := range fs {
		g := fmt.Sprintf("%g", f)
		ff := fmt.Sprintf("%f", f)
		F := fmt.Sprintf("%F", f)
		G := fmt.Sprintf("%G", f)
		fmt.Printf("g:%30s f:%30s F:%30s G:%30s\n", g, ff, F, G)
	}
	//var f float64 = 21.13342424
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

	fmt.Println(fmt.Errorf("error info is %s", "id must be int type"))

	err := errors.New("error info")
	fmt.Println(fmt.Sprintf("couldn't get a downloader: %s", err))
}
