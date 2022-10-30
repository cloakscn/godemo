package main

import (
	"fmt"

	"example.com/demo/demo_05/client_stub"
)

func main() {
	client := client_stub.NewHelloServiceClient("tcp", "localhost:8080")

	var reply string
	err := client.Hello("cloaks", &reply)
	if err != nil {
		panic(fmt.Sprintf("connect error: %s", err))
	}
	fmt.Println(reply)
}

// json {"method": "HelloService.Hello", "params": ["cloaks"], "id": 0}
