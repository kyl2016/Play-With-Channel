package main

import "fmt"

func main() {
	var d = 72.1234567
	fmt.Printf("%.1f\n", d)
	fmt.Printf("%.2f\n", d)
	fmt.Printf("%.3f\n", d)
	fmt.Printf("%.4f\n", d)
	fmt.Printf("%.5f\n", d)

	d = 72.5
	fmt.Printf("d: %d\n", int(d))

	fmt.Printf("%.3f\n", float64(150)/float64(10000)*100)

	var d2 float64 = 1.01212
	fmt.Printf("%f\n", d2)
}
