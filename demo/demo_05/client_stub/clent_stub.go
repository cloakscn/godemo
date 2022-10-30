package client_stub

import (
	"fmt"
	"net/rpc"

	"example.com/demo/demo_05/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol string, address string) HelloServiceStub {
	dial, err := rpc.Dial(protol, address)
	if err != nil {
		panic(fmt.Sprintf("dial error: %s", err))
	}
	return HelloServiceStub{dial}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
