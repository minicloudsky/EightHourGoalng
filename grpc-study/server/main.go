package main

//
//import (
//	hellogrpc "github.com/minicloudsky/grpc-study/pb/hello"
//	"google.golang.org/grpc"
//	"net"
//)
//import "context"
//import "fmt"
//
//type server struct {
//	hellogrpc.UnimplementedHelloGRPCServer
//}
//
//func (s *server) apply(options *interface{}) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *server) SayHi(ctx context.Context, req *hellogrpc.Req) (res *hellogrpc.Res, err error) {
//	fmt.Println(req.GetMessage())
//	return &hellogrpc.Res{Message: "Hello,gRPC Response"}, nil
//}
//
//func main() {
//	l, _ := net.Listen("tcp", ":8888")
//	s := grpc.NewServer()
//	hellogrpc.RegisterHelloGRPCServer(s, &server{})
//	err := s.Serve(l)
//	if err != nil {
//		return
//	}
//}
