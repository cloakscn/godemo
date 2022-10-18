package main

import (
	"context"
	"fmt"
	"time"

	demo "example.com/test/demo/demo_14/proto/gen/demo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial(":80", grpc.WithInsecure())
	if err != nil  {
		panic(err)
	}
	defer cc.Close()
	dsc := demo.NewDemoServiceClient(cc)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	resp, err := dsc.SayHello(ctx, &demo.DemoRequest{
		Name: "cloak",
	})
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			panic(err)
		}
		fmt.Printf("s.Message(): %v\n", s.Message())
		fmt.Printf("s.Code(): %v\n", s.Code())
	} else {
		fmt.Printf("resp.Message: %v\n", resp)
	}
}