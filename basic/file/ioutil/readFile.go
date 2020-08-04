package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	folder := "/home/lynxi/go/src/github.com/lynxitech/ivs/volumes/169/face"
	folder = "/home/lynxi/Documents/faces/6万人/169/face"
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	count := 0
	for _, file := range files {
		count++
		//fmt.Println(file.Name())
		buf, err := ioutil.ReadFile(folder + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		//println(len(buf))
		_ = len(buf)
	}

	fmt.Println(count, time.Since(start).Seconds())
}
