package main

import "fmt"

func main() {
	list := []string{"a", "bcd", "ef", "g", "hij"}
	res := reduce(Map(list, length), sum)
	fmt.Println(res)
}

func length(s string) int {
	return len(s)
}

func Map(list []string, fn func(string) int) []int {
	res := make([]int, len(list))
	for i, elem := range list {
		res[i] = fn(elem)
	}

	return res
}

func reduce(list []int, fn func(int, int) int) (res int) {
	for _, elem := range list {
		res = fn(res, elem)
	}
	return res
}

func sum(a, b int) int {
	return a + b
}
