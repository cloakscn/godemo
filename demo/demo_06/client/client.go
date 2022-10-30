package main

import (
	"context"
	"fmt"

	"example.com/demo/demo_06/proto"
	"google.golang.org/grpc"
)

func main() {
	// connect remote port
	dial, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("listen error: %s", err))
	}
	defer dial.Close()

	// get client
	client := proto.NewGreeterClient(dial)
	// visit remote function
	reply, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "cloaks",
	})
	if err != nil {
		panic(fmt.Sprintf("listen error: %s", err))
	}
	fmt.Println(reply)
}
