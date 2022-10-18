package main

import (
	"context"
	"net"
	"time"

	demo "example.com/test/demo/demo_14/proto/gen/demo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HelloService struct{}

func (s *HelloService) SayHello(ctx context.Context, in *demo.DemoRequest) (*demo.DemoResponse, error) {
	if in.Name == "cloaks" {
		return &demo.DemoResponse{
			Message: in.Name,
		}, nil
	} else {
		time.Sleep(time.Second*5)
		return nil, status.Error(codes.Unknown, "未知用户")
	}
}

func main() {
	s := grpc.NewServer()

	demo.RegisterDemoServiceServer(s, &HelloService{})

	l, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(err)
	}
	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}
