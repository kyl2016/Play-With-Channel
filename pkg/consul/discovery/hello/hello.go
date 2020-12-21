package hello

import (
	"github.com/kyl2016/Play-With-Golang/pkg/consul/discovery/hello/pb"
)

type GreeterServerImp struct{}

func (g *GreeterServerImp) SayHello(srv pb.Greeter_SayHelloServer) error {

	req, err := srv.Recv()
	if err != nil {
		panic(err)
	}
	srv.Send(&pb.HelloReply{
		Message: "hello " + req.Name,
	})

	return nil
}
