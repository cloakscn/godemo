package main

import (
	"context"
	"fmt"
	// "time"

	v1 "example.com/demo/demo_11/proto/gen/interceptor/v1"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/metadata"
)

type customCredential struct {}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string {
		"APPID": "2018061036",
		"APPKEY": "Ycu061036",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}


func main() {
	// interceptor := func (ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 	start := time.Now()
		
	// 	md := metadata.New(map[string]string{
	// 		"APPID": "2018061036",
	// 		"APPKEY": "Ycu@061036",
	// 	})
	// 	ctx = metadata.NewOutgoingContext(context.Background(), md)
	// 	err := invoker(ctx, method, req, reply, cc, opts...)
	// 	fmt.Printf("耗时：%s\n", time.Since(start))
	// 	return err
	// }

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// opts = append(opts, grpc.WithUnaryInterceptor(interceptor)) 
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{})) 
	cc, err := grpc.Dial(":80", opts...)
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