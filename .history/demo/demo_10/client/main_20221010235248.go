package main

import (
	"context"
	"fmt"
	"time"

	v1 "example.com/test/demo/demo_10/proto/gen/metadata/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/internal/metadata"
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
		"timestamp": string(time.Now().Day())
	})
	reply, err := hsc.SayHello(context.Background(), &v1.HelloRequest{
		Name: "cloaks",
	})

	if err != nil {
		panic(fmt.Sprintf("failt to sayHello: %v", err))
	}

	fmt.Println(reply)


}