package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var file = "/tmp/cacheTest"

func main() {
	buf := make([]byte, 8192)
	for i := 0; i < 20; i++ {
		buf[i] = byte(i)
	}

	//fmt.Println("----------------------------------------")
	//fmt.Println("write without alignment and DIRECT:")
	//writeWithoutAlignmentWithoutDIRECT(buf)
	//fmt.Println("----------------------------------------")
	//fmt.Println("write without alignment but with DIRECT:")
	//writeWithoutAlignmentWithDIRECT(buf)
	//fmt.Println("----------------------------------------")
	//fmt.Println("write with alignment and DIRECT:")
	//writeWithAlignmentWithDIRECT(buf)
	//fmt.Println("----------------------------------------")
	//fmt.Println("read without alignment and DIRECT:")
	//readWithoutAlignmentWithoutDIRECT(buf)
	//fmt.Println("----------------------------------------")
	//fmt.Println("read without alignment but with DIRECT:")
	//readWithoutAlignmentWithDIRECT(buf)
	//fmt.Println("----------------------------------------")
	//fmt.Println("read with alignment and DIRECT:")
	//readWithAlignmentWithDIRECT(buf)

	readWithAlignmentWithDIRECT2(buf)
}

func writeWithoutAlignmentWithoutDIRECT(buf []byte) {
	// open file
	file, err := os.OpenFile("/tmp/sdb",
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer file.Close()

	// write file
	fmt.Println("buffer ", unsafe.Pointer(&buf))
	fmt.Println("buffer[0] ", unsafe.Pointer(&buf[0]))
	buf2 := buf[4:516]
	fmt.Println("write with buffer ", unsafe.Pointer(&buf2[0]))

	_, err = file.WriteAt(buf2, 512)
	if err != nil {
		fmt.Println("write error ", err)
	} else {
		fmt.Println("write succeed")
	}
}
func writeWithoutAlignmentWithDIRECT(buf []byte) {
	// open file
	file, err := os.OpenFile("/tmp/sdc",
		os.O_WRONLY|os.O_CREATE|syscall.O_DIRECT, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer file.Close()

	// write file
	fmt.Println("buffer ", unsafe.Pointer(&buf))
	fmt.Println("buffer[0] ", unsafe.Pointer(&buf[0]))
	buf2 := buf[4:516]
	fmt.Println("write with buffer ", unsafe.Pointer(&buf2[0]))

	_, err = file.WriteAt(buf2, 512)
	if err != nil {
		fmt.Println("write error ", err)
	} else {
		fmt.Println("write succeed")
	}
}
func writeWithAlignmentWithDIRECT(buf []byte) {
	// open file
	file, err := os.OpenFile("/tmp/sdd",
		os.O_WRONLY|os.O_CREATE|syscall.O_DIRECT, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer file.Close()

	// write file
	fmt.Println("buffer ", unsafe.Pointer(&buf))
	fmt.Println("buffer[0] ", unsafe.Pointer(&buf[0]))
	buf2 := buf[512 : 512+512]
	fmt.Println("write with buffer ", unsafe.Pointer(&buf2[0]))

	_, err = file.WriteAt(buf2, 512)
	if err != nil {
		fmt.Println("write error ", err)
	} else {
		fmt.Println("write succeed")
	}
}

func readWithoutAlignmentWithoutDIRECT(buf []byte) {
	// read file
	file, err := os.OpenFile("/tmp/sdb", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("An error occurred whit file ipening.\n")
		return
	}
	defer file.Close()

	buf = buf[2:514]
	fmt.Println("read with buffer ", unsafe.Pointer(&buf[0]))

	_, err = file.ReadAt(buf, 512)
	if err != nil {
		fmt.Println("read error ", err)
	} else {
		fmt.Println("read succeed", buf)
	}
}
func readWithoutAlignmentWithDIRECT(buf []byte) {
	// read file
	file, err := os.OpenFile("/tmp/sdc", os.O_RDONLY|syscall.O_DIRECT, 0666)
	if err != nil {
		fmt.Printf("An error occurred whit file ipening.\n")
		return
	}
	defer file.Close()

	buf = buf[2:514]
	fmt.Println("read with buffer ", unsafe.Pointer(&buf[0]))

	_, err = file.ReadAt(buf, 512)
	if err != nil {
		fmt.Println("read error ", err)
	} else {
		fmt.Println("read succeed", buf)
	}
}
func readWithAlignmentWithDIRECT(buf []byte) {
	// read file
	file, err := os.OpenFile("/tmp/sdd", os.O_RDONLY|syscall.O_DIRECT, 0666)
	if err != nil {
		fmt.Printf("An error occurred whit file ipening.\n")
		return
	}
	defer file.Close()

	buf = buf[512 : 512+512]
	fmt.Println("read with buffer ", unsafe.Pointer(&buf[0]))

	_, err = file.ReadAt(buf, 1024)
	if err != nil {
		fmt.Println("read error ", err)
	} else {
		fmt.Println("read succeed", buf)
	}
}

func readWithAlignmentWithDIRECT2(buf []byte) {
	// read file
	file, err := os.OpenFile("/home/lynxi/Documents/faces/6万人/6万人.zip", os.O_RDONLY|syscall.O_DIRECT, 0666)
	if err != nil {
		fmt.Printf("An error occurred whit file ipening.\n")
		return
	}
	defer file.Close()

	buf2 := make([]byte, 1024)
	fmt.Println("read with buffer ", unsafe.Pointer(&buf2[0]))

	//info, err := file.Stat()
	//if err != nil {
	//	panic(err)
	//}

	_, err = file.ReadAt(buf2, 512) // offset must be 512*n, n>=0
	if err != nil {
		fmt.Println("read error ", err)
	} else {
		fmt.Println("read succeed", buf2)
	}
}
