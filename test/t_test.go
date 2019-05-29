package test

import (
	"fmt"
	"log"
	"testing"
)

func TestT(t *testing.T) {
	fmt.Println("test")
}

func BenchmarkT(t *testing.B) {
	t.Error("error")
	fmt.Println("test benchmark ", t.N)
}

func ExampleT(){
	fmt.Println("example")

	log.Fatal("fatal error")
}