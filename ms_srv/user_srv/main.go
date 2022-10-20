package main

import (
	"flag"
	"fmt"
	"net"

	"example.com/test/ms_srv/user_srv/handler"
	"example.com/test/ms_srv/user_srv/model/proto"
	"google.golang.org/grpc"
)

func main() {
	// 获取命令行 flag
	IP := flag.String("ip", "0.0.0.0", "ip 地址")
	Port := flag.Int("port", 8080, "端口号")
	flag.Parse()

	fmt.Printf("IP: %v\n", *IP)
	fmt.Printf("Port: %v\n", *Port)

	s := grpc.NewServer()
	proto.RegisterUserServer(s, &handler.UserServer{})
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic(err)
	}
	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}
