package main

import (
	"net"
	"net/rpc"

	"example.com/demo/demo_05/handler"
	"example.com/demo/demo_05/server_stub"
)

func main() {
	// initialization server
	listener, _ := net.Listen("tcp", "localhost:8080")
	// register
	_ = server_stub.RegisterHelloService(&handler.NewHelloService{})
	// start server
	for {
		accept, _ := listener.Accept()
		go rpc.ServeConn(accept)
	}
}
