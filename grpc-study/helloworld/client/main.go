package main

import (
	"context"
	"fmt"
	hellogrpc "github.com/minicloudsky/grpc-study/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	if err != nil {
		return
	}
	client := hellogrpc.NewHelloGRPCClient(conn)
	req, err := client.SayHi(context.Background(), &hellogrpc.Req{Message: "Hello,gRPC!"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(req.GetMessage())

}
