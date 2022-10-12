package main

import (
	"context"
	"fmt"

	v1 "example.com/test/demo/demo_11/proto/gen/interceptor/v1"
	"google.golang.org/grpc"
)


func main() {

	var opts []grp 
	cc, err := grpc.Dial(":80", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("failed to dial: %v", err))
	}
	defer cc.Close()

	hsc := v1.NewHelloServiceClient(cc)

	reply, err := hsc.SayHello(context.Background(), &v1.HelloRequest{
		Name: "cloaks",
	})

	if err != nil {
		panic(fmt.Sprintf("failt to sayHello: %v", err))
	}

	fmt.Println(reply)


}