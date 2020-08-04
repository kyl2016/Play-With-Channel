package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println("x cap:",cap(x), "x:",x)

	var z *[]int
	z = &x
	fmt.Println(z)

	x[1] = 11
	fmt.Println("x cap:",cap(x), "x:",x)
	fmt.Println("z:", *z)


	x = append(x, 4, 5, 6)	
	fmt.Println("x cap:",cap(x), "x:",x)
	fmt.Println("z:", *z)

	y := []int{7, 8, 9}
	x = append(x, y...)
	fmt.Println(x)
}
