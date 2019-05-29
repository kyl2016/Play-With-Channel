package main

import (
	"context"
	pb "github.com/kyl2016/Play-With-Golang/gRPC/helloWorld/helloWorld"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	//address     = "localhost:50051"
	address     = "192.168.8.223:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("couldn't greet: %v", err)
	}

	log.Println("Greeting:", r.Message)
}
