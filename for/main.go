package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	//procs := runtime.GOMAXPROCS(1)
	//fmt.Println("procs:", procs)
	//fmt.Println("NumCPU:", runtime.NumCPU())
	//
	//procs = runtime.GOMAXPROCS(1)
	//fmt.Println("procs:", procs)
	//fmt.Println("NumCPU:", runtime.NumCPU())

	//err := runtime.StartTrace()
	//if err != nil {
	//	panic(err)
	//}

	println(os.Getpid())

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGUSR1)
	//
	//go func() {
	//	println("waiting...")
	//	for range c {
	//		DumpStacks()
	//	}
	//}()

	go func() { fmt.Println(1) }()
	go func() { println(2) }()
	go func() { println(3) }()

	println(4)

	for {
	}
}

func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("--- BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump", buf)
}
