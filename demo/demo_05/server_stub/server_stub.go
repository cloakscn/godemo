package server_stub

import (
	"net/rpc"

	"example.com/test/demo/demo_05/handler"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(service HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, service)
}
