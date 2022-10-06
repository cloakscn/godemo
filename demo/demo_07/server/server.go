package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"example.com/test/demo/demo_07/proto"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	err := res.Send(&proto.StreamResData{
		Data: fmt.Sprintf("req data: %s", req.Data),
	})
	if err != nil {
		return err
	}
	for {
		err = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("get stream: %v", time.Now().Unix()),
		})
		if err != nil {
			fmt.Printf("error: %s\n", err)
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *Server) PutStream(req proto.Greeter_PutStreamServer) error {
	for {
		srd, err := req.Recv()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			break
		}
		fmt.Println(srd)
	}
	return nil
}

func (s *Server) AllStream(req proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func ()  {
		defer wg.Done()
		for {
			srd, err := req.Recv()
			if err != nil {
				fmt.Printf("error: %s\n", err)
				break
			}
			fmt.Println(srd)
		}
	}()

	go func ()  {
		defer wg.Done()
		for {
			err := req.Send(&proto.StreamResData{
				Data: fmt.Sprintf("server stream: %v", time.Now().Unix()),
			})
			if err != nil {
				fmt.Printf("error: %s\n", err)
				break
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})

	err = s.Serve(l)
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
}
