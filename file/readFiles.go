package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "/home/lynxi/Documents/faces/lfw/lfw_raw.zip"
	path = "/home/lynxi/Documents/faces/6万人/6万人.zip"

	read, err := zip.OpenReader(path)
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	log.Println(len(read.File))

	count := 0

	for _, file := range read.File {
		count++
		if err := listFiles(file); err != nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}

	log.Println(count)

	//file, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}
	//
	//bufio.NewReader(file).ReadLine()
}

func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	fmt.Fprintf(os.Stdout, "%s:", file.Name)

	if err != nil {
		msg := "Failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	fmt.Println()

	return nil
}
