package main

import (
	"context"
	"fmt"
	"net"

	timestamp "example.com/demo/demo_09/proto/gen/timestamp/v1"
	"google.golang.org/grpc"
)

type Hello struct{}

func (h *Hello) SayHello(ctx context.Context, in *timestamp.HelloRequest) (*timestamp.HelloResponse, error) {
	return &timestamp.HelloResponse{
		Data: fmt.Sprintf("gender is %d and timestamp is %s", in.Gender, in.CreateTime),
	}, nil
}

func main() {
	s := grpc.NewServer()

	timestamp.RegisterHelloServerServer(s, &Hello{})

	l, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(err)
	}

	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}
