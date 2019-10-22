package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	message := []byte("Hello")
	filename := "/tmp/test"
	err := ioutil.WriteFile(filename, message, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file, err := os.Stat(filename)
		if err != nil {
			log.Fatal(err)
		}
		println(file.Name())
	}()
	defer os.Remove(filename)
}
