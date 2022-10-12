package main

import (
	"context"
	"fmt"
	"net"

	v1 "example.com/test/demo/demo_11/proto/gen/interceptor/v1"
	"google.golang.org/grpc"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
	return &v1.HelloResponse{
		Data: fmt.Sprintf("Hello %s", in.Name),
	}, nil
}

func main() {
	s := grpc.NewServer()

	v1.RegisterHelloServiceServer(s, &Server{})
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v"))
	}
}