package test

import "fmt"

func ExampleFmt() {
	fmt.Printf("%15s\t%s\n", "abc", "def")
	fmt.Printf("%15d\n", 123)
	fmt.Printf("%t\n", true)

	// Output:
	// abc def
	//             123
	// true
}
