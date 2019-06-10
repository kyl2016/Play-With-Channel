package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f1 float32 = 9.90
	fmt.Println(f1 * 100)
	var f2 float64 = 9.90
	fmt.Println(f2 * 100)

	s := strconv.FormatFloat(1234.5678, 'g', 8, 32)
	fmt.Println(s)
	fmt.Println(strconv.ParseFloat(s, 32))
	s = strconv.FormatFloat(1234.5678, 'g', 8, 64)
	fmt.Println(s)
	fmt.Println(strconv.ParseFloat(s, 64))

	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}

	fmt.Println(n)

	n = 0
	for i := 0; i < 1000; i++ {
		n += .01 * 100
	}

	fmt.Println(n / 100)
}
