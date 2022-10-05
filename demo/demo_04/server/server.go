package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {
	// register
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// initialization server
	http.HandleFunc("/jsonRPC", func(w http.ResponseWriter, r *http.Request) {
		var connect io.ReadWriteCloser = struct {
			io.ReadCloser
			io.Writer
		} {
			ReadCloser: r.Body,
			Writer: w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(connect))
	})
	// start server
	http.ListenAndServe(":8080", nil)
}
