package main

import "fmt"

func main() {
	var d = 72.8
	fmt.Printf("%.1f", d)

	s := fmt.Sprintf("%.1f\n", d)
	fmt.Println(s)

	fmt.Printf("%.2f%%\n", float64(150)/float64(10000)*100)

	var d2 float64  = 1.01212
	fmt.Printf("%f", d2)
}
