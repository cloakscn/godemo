package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {
	// initialization server
	listen, _ := net.Listen("tcp", "localhost:8080")
	// register
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// start server
	for {
		accept, _ := listen.Accept()
		go 
		
		rpc.ServeCodec(jsonrpc.NewServerCodec(accept))
	}
}
