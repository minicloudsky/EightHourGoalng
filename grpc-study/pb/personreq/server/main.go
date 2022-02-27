package main

import (
	context "context"
	"fmt"
	person "github.com/minicloudsky/golang-in-action/grpc-study/pb/personreq"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"time"
)

type personServe struct {
	person.UnimplementedSearchServiceServer
}

func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{
		Name: "我收到了信息 - " + name,
		Age:  26,
	}
	fmt.Println("client res: ", res)
	return res, nil
}
func (*personServe) SearchIn(server person.SearchService_SearchInServer) error {

	for {
		req, err := server.Recv()
		fmt.Println("req: ", req)
		if err != nil {
			_ = server.SendAndClose(&person.PersonRes{
				Name: "完成了",
				Age:  -1,
			})

			break
		}
	}
	return nil

}
func (*personServe) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {

	name := req.Name
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		err := server.Send(&person.PersonRes{
			Name: strconv.Itoa(i) + "-我拿到了 " + name,
			Age:  int32(i),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func (*personServe) SearchIO(server person.SearchService_SearchIOServer) error {
	str := make(chan string)
	c := 0
	go func() {
		for {
			req, err := server.Recv()
			fmt.Println("req: ", req)
			if err != nil {
				str <- "结束"
				break
			}
			str <- req.Name
		}
	}()
	for {
		s := <-str
		err := server.Send(&person.PersonRes{
			Name: strconv.Itoa(c) + "-" + s,
			Age:  100,
		})
		c += 1
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
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
}
