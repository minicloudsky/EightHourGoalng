package main

import (
	context "context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	person "github.com/minicloudsky/golang-in-action/grpc-study/grpc-gateway/personreq"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

type personServe struct {
	person.UnimplementedSearchServiceServer
}

func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{
		Name: "我收到了信息 - " + name + "I'm from grpc-gateway",
		Age:  26,
	}
	fmt.Println("client res: ", res)
	return res, nil
}

//function (*personServe) SearchIn(server person.SearchService_SearchInServer) error {
//
//	for {
//		req, err := server.Recv()
//		fmt.Println("req: ", req)
//		if err != nil {
//			_ = server.SendAndClose(&person.PersonRes{
//				Name: "完成了",
//				Age:  -1,
//			})
//
//			break
//		}
//	}
//	return nil
//
//}
//function (*personServe) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
//
//	name := req.Name
//	for i := 0; i < 10; i++ {
//		time.Sleep(1 * time.Second)
//		err := server.Send(&person.PersonRes{
//			Name: strconv.Itoa(i) + "-我拿到了 " + name,
//			Age:  int32(i),
//		})
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//function (*personServe) SearchIO(server person.SearchService_SearchIOServer) error {
//	str := make(chan string)
//	c := 0
//	go function() {
//		for {
//			req, err := server.Recv()
//			fmt.Println("req: ", req)
//			if err != nil {
//				str <- "结束"
//				break
//			}
//			str <- req.Name
//		}
//	}()
//	for {
//		s := <-str
//		err := server.Send(&person.PersonRes{
//			Name: strconv.Itoa(c) + "-" + s,
//			Age:  100,
//		})
//		c += 1
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateway(&wg)
	go registerGRPC(&wg)
	wg.Wait()
}

func registerGateway(wg *sync.WaitGroup) {
	conn, _ := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8888",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	mux := runtime.NewServeMux() // 一个对外开放的 mux
	gatewayServer := &http.Server{
		Handler: mux,
		Addr:    ":8090",
	}
	err := person.RegisterSearchServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return
	}
	err = gatewayServer.ListenAndServe()
	if err != nil {
		return
	}
	wg.Done()
}
func registerGRPC(wg *sync.WaitGroup) {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServe{})
	fmt.Println("grpc serer started!")
	err = s.Serve(listen)
	if err != nil {
		return
	}
	wg.Done()
}
