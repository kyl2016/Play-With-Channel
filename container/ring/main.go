package main

import (
	"container/ring"
	"fmt"
)

func main(){
	r := ring.New(10)

	r.Do(func(i interface{}) {
		fmt.Println("do")
	})



	fmt.Println(r.Value)
}
