package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(fmt.Sprintf("dial error: %s", err))
	}

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(dial))
	err = client.Call("HelloService.Hello", "cloaks", &reply)
	if err != nil {
		panic(fmt.Sprintf("connect error: %s", err))
	}
	fmt.Println(reply)
}

// json {"method": "HelloService.Hello", "params": ["cloaks"], "id": 0}