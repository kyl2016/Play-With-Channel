package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		fmt.Println(x)
	}

	r := rand.New(rand.NewSource(50))
	fmt.Println(r.Int31())
	fmt.Println(r.Float32())
	fmt.Println(r.Int31n(10))
	fmt.Println(r.Int31n(100))
	fmt.Println(r.Int31n(100))

	fmt.Println(rand.Int31n(50))
	fmt.Println(rand.Int31n(50))
	fmt.Println(rand.Int31n(50))
}
