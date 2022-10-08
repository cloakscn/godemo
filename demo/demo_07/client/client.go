package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	stream "example.com/test/demo/common/stream/proto/v1"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
	defer cc.Close()

	gc := stream.NewGreeterClient(cc)

	// getStream(gc)

	// putStream(gc)

	allStream(gc)
}

func allStream(gc stream.GreeterClient) {
	g, err := gc.AllStream(context.Background())
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			srd, err := g.Recv()
			if err != nil {
				fmt.Printf("error: %s\n", err)
				break
			}
			fmt.Println(srd)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err = g.Send(&stream.StreamReqData{
				Data: fmt.Sprintf("client stream: %v", time.Now().Unix()),
			})
			if err != nil {
				fmt.Printf("error: %s\n", err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}

func putStream(gc stream.GreeterClient) {
	g, err := gc.PutStream(context.Background())
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
	for {
		err = g.Send(&stream.StreamReqData{
			Data: fmt.Sprintf("put stream: %v", time.Now().Unix()),
		})
		if err != nil {
			fmt.Printf("error: %s", err)
			break
		}
		time.Sleep(time.Second)
	}
}

func getStream(gc stream.GreeterClient) {
	res, err := gc.GetStream(context.Background(), &stream.StreamReqData{
		Data: "cloaks",
	})
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
	for {
		srd, err := res.Recv()
		if err != nil {
			panic(fmt.Sprintf("error: %s", err))
		}
		fmt.Println(srd)
	}
}
