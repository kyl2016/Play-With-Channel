package main

import "math"

func main() {
	println(3%4, -3%4, -3%2, -1%6, -3%-5)
	println(math.Mod(-3, -4))
	println(math.Mod(-3, 4))
}
