package main

import (
	"context"
	"fmt"
	"time"

	v1 "example.com/demo/demo_09/proto/gen/timestamp/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	cc, err := grpc.Dial(":80", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	
	hsc := v1.NewHelloServerClient(cc)
	res, err := hsc.SayHello(context.Background(), &v1.HelloRequest{
		Gender: v1.Gender_FEMALE,
		CreateTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Data)
}