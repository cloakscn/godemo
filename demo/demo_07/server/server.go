package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	stream "example.com/demo/common/stream/proto/v1"
)

type Server struct{}

func (s *Server) GetStream(req *stream.StreamReqData, res stream.Greeter_GetStreamServer) error {
	err := res.Send(&stream.StreamResData{
		Data: fmt.Sprintf("req data: %s", req.Data),
	})
	if err != nil {
		return err
	}
	for {
		err = res.Send(&stream.StreamResData{
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

func (s *Server) PutStream(req stream.Greeter_PutStreamServer) error {
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

func (s *Server) AllStream(req stream.Greeter_AllStreamServer) error {
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
			err := req.Send(&stream.StreamResData{
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
	stream.RegisterGreeterServer(s, &Server{})

	err = s.Serve(l)
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
}
