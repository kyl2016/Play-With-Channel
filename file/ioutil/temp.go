package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	name, err := ioutil.TempDir("", "temp_")
	if err != nil {
		log.Fatal(err)
	}

	println(name)

	name, err = ioutil.TempDir("/home/lynxi/Documents", "temp_")
	if err != nil {
		log.Fatal(err)
	}

	println(name)

	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	println(tmpfile.Name())
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if err = tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
