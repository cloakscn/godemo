package main

import (
	"context"
	"fmt"
	"net"

	v1 "example.com/test/demo/demo_10/proto/gen/metadata/proto/v1"
	"google.golang.org/grpc"
)


type Welcome struct {}

func (s *Welcome) SayHello(ctx context.Context, request *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{
		Message: fmt.Sprintf("hello %s", request.Name),
	}, nil
}

func main() {
	s := grpc.NewServer()
	v1.RegisterHelloServerServer(s, &Welcome{})
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(fmt.Sprintf("failed to liten: %v", err))
	}
	err = s.Serve(l)
	if err != nil {
		panic(fmt.Sprintf("failed to start serve: %v", err))
	}
}