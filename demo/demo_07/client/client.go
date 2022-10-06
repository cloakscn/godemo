package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"example.com/test/demo/demo_07/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
	defer cc.Close()

	gc := proto.NewGreeterClient(cc)

	// getStream(gc)

	// putStream(gc)

	allStream(gc)
}

func allStream(gc proto.GreeterClient) {
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
			err = g.Send(&proto.StreamReqData{
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

func putStream(gc proto.GreeterClient) {
	g, err := gc.PutStream(context.Background())
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
	for {
		err = g.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("put stream: %v", time.Now().Unix()),
		})
		if err != nil {
			fmt.Printf("error: %s", err)
			break
		}
		time.Sleep(time.Second)
	}
}

func getStream(gc proto.GreeterClient) {
	res, err := gc.GetStream(context.Background(), &proto.StreamReqData{
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
