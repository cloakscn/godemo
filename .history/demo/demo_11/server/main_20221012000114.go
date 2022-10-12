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
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)  {
		fmt.Println("接收到一个新的请求")
		i, err := handler(ctx, req)
		if err != nil {
			panic(fmt.s)
		}
		fmt.Println("请求结束")
		return i, nil
	}

	grpc.UnaryInterceptor(interceptor)
	s := grpc.NewServer()

	v1.RegisterHelloServiceServer(s, &Server{})
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	err = s.Serve(l)
	if err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}