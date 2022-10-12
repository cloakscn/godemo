package main

import (
	v1 "example.com/test/demo/demo_11/proto/gen/interceptor/v1"
	"google.golang.org/grpc"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
	
}

func main() {
	s := grpc.NewServer()

	v1.RegisterHelloServerServer(s, &)
}