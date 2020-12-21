package main

import (
	"fmt"
	"github.com/kyl2016/Play-With-Golang/pkg/consul/discovery/hello"
	"github.com/kyl2016/Play-With-Golang/pkg/consul/discovery/hello/pb"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8500/v1/health/service/hello")
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))

	lis1, _ := net.Listen("tcp", ":8888")
	lis2, _ := net.Listen("tcp", ":8889")
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &hello.GreeterServerImp{})
	go grpcServer.Serve(lis1)
	go grpcServer.Serve(lis2)

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}

//func main() {
//	var config *consul.Config = &consul.Config{
//		Address:    "172.17.0.6:8301",
//		Scheme:     "",
//		Datacenter: "",
//		Transport:  nil,
//		HttpClient: nil,
//		HttpAuth:   nil,
//		WaitTime:   0,
//		Token:      "",
//		TokenFile:  "",
//		Namespace:  "",
//		TLSConfig:  consul.TLSConfig{},
//	}
//	consulClient, err := consul.NewClient(config)
//	if err != nil {
//		panic(err)
//	}
//
//	c := Client{Consul: consulClient}
//	err = c.Register("myService", 5039)
//	if err != nil {
//		panic(err)
//	}
//
//	entry, meta, err := c.Service("myService", "")
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(entry, meta)
//}
