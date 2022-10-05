package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	dial, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(fmt.Sprintf("connect error: %s", err))
	}

	var reply string
	err = dial.Call("HelloService.Hello", "cloaks", &reply)
	if err != nil {
		panic(fmt.Sprintf("dial error: %s", err))
	}
	fmt.Println(reply)
}
