package main

import (
	"context"
	"fmt"

	v1 "example.com/test/demo/demo_11/proto/gen/metadata/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)


func main() {
	cc, err := grpc.Dial(":80", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("failed to dial: %v", err))
	}
	defer cc.Close()

	hsc := v1.NewHelloServerClient(cc)

	// md := metadata.Pairs("timestamp", time.Now().Format(time.Stamp))
	md := metadata.New(map[string]string{
		"name": "cloaks",
		"timestamp": "good morning",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	reply, err := hsc.SayHello(ctx, &v1.HelloRequest{
		Name: "cloaks",
	})

	if err != nil {
		panic(fmt.Sprintf("failt to sayHello: %v", err))
	}

	fmt.Println(reply)


}