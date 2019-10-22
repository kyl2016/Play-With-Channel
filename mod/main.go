package main

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"fmt"
)

func main() {
	println(primitive.NewObjectID().Hex())
	fmt.Println("hello")
}
