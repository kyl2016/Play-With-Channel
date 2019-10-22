package main

import (
	"encoding/binary"
	"fmt"
	"github.com/cpmech/gosl/la/mkl"
	"io/ioutil"
	"math"
	"testing"
)

func Test_Compare(t *testing.T) {
	//fmt.Println(CompareEigens("李小龙3/caches/faceEigens", "李小龙/caches/faceEigens"))
	fmt.Println(CompareEigens("/home/lynxi/go/src/github.com/lynxitech/ivs/volumes/retrieval_images/deb09e95c02f22914f0828c6fb76c177/caches/faceEigens", "/home/lynxi/go/src/github.com/lynxitech/ivs/volumes/retrieval_images/deb09e95c02f22914f0828c6fb76c177/caches/faceEigens"))
}

func CompareEigens(file1, file2 string) float64 {
	buffer1, _ := ioutil.ReadFile(file1)
	data1 := readNums(buffer1)

	buffer2, _ := ioutil.ReadFile(file2)
	data2 := readNums(buffer2)

	n, incx, incy := 512, 1, 1
	return mkl.Ddot(n, data1, incx, data2, incy)
}

func readNums(buffer []byte) []float64 {
	fmt.Println()

	data := []float64{}

	for i := 0; i < 512; i++ {
		from := i * 4
		to := i*4 + 4

		r := float64(Float32frombytes(buffer[from:to]))
		data = append(data, r)
	}

	return data
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func Float32frombytes2(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}
