package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("a12")
	fmt.Println("i:",i,"\nerr:", err)
}
