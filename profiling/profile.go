package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

func main() {
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatal(err)
	}
	fullPath, err := filepath.Abs("cpu_profile")
	fmt.Println(fullPath)

	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	memf, err := os.Create("mem_profile")
	if err != nil {
		log.Fatal("could not create memory profile:", err)
	}
	if err := pprof.WriteHeapProfile(memf); err != nil {
		log.Fatal("could not write memory profile:", err)
	}
	defer memf.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	test22(ctx)

	time.Sleep(time.Second * 3)
}

func test22(c context.Context) {
	i := 0
	j := 0
	go func() {
		m := map[int]int{}
		for {
			i++
			m[i] = i
		}
	}()
	go func() {
		m := map[int]int{}
		for {
			j++
			m[i] = i
		}
	}()
	select {
	case <-c.Done():
		fmt.Println("done, i", i, "j", j)
		return
	}
}
