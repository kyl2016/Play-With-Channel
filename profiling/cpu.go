package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "cpu_profile", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()

		do()
	} else {
		fmt.Println("cpuprofile is null")
	}
}

func do() {
	var c int = 1
	for i := 1; i <= int(math.Pow(2, 32)); i++ {
		c += (i * i)
	}
	fmt.Println(c)
}
