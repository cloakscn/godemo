package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"example.com/test/demo/demo_06/proto"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	// init
	server := grpc.NewServer()
	// register Server to grpc server
	proto.RegisterGreeterServer(server, &Server{})
	// listen port
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(fmt.Sprintf("listen error: %s", err))
	}
	// start server on listen port
	err = server.Serve(l)
	if err != nil {
		panic(fmt.Sprintf("Start Serve error: %s", err))
	}
}
