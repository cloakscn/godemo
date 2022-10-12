package main

import (
	v1 "example.com/test/demo/demo_09/proto/gen/timestamp/v1"
	"google.golang.org/grpc"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, in *v1.HelloRequest) (*HelloResponse, error) {

}

func main() {
	s := grpc.NewServer()

	v1.RegisterHelloServerServer(s, &)
}