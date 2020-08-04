package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))

	re2 := regexp.MustCompile(": \"(.*?)\",")
	r := re2.Find([]byte("        \"Path\": \"cloud.google.com/go\",\n"))
	fmt.Printf("%s\n", r)
	fmt.Printf("%s\n", re2.FindAllSubmatch([]byte("        \"Path\": \"cloud.google.com/go\",\n"), 1)[0][1])
	fmt.Printf("%s\n", re2.FindAllStringSubmatch("        \"Path\": \"cloud.google.com/go\",\n", 1)[0][1])
	fmt.Printf("%s\n", re2.FindStringSubmatch("        \"Path\": \"cloud.google.com/go\",\n")[1])

	re3 := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re3.FindAllSubmatch([]byte(`seafood fool`), -1))
}
