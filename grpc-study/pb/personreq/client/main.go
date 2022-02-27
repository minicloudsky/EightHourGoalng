package main

import (
	"context"
	"fmt"
	person "github.com/minicloudsky/golang-in-action/grpc-study/pb/personreq"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	addr := "127.0.0.1:8888"
	l, _ := grpc.Dial(addr, grpc.WithInsecure())
	client := person.NewSearchServiceClient(l)
	// search 普通的 request response
	//for i := 0; i < 100; i++ {
	//	res, err := client.Search(context.Background(), &person.PersonReq{
	//		Name: strconv.Itoa(i) + " - Hello,I'm minicloudsky",
	//		Age:  int32(i),
	//	})
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("res: ", res)
	//}
	// 客户端流 request
	//c, err := client.SearchIn(context.Background())
	//if err != nil {
	//	fmt.Println("Search In err")
	//	return
	//}
	//for i := 0; i < 10; i++ {
	//	time.Sleep(1 * time.Second)
	//	err := c.Send(&person.PersonReq{
	//		Name: "我是进来的信息",
	//		Age:  int32(i),
	//	})
	//	if err != nil {
	//		return
	//	}
	//	if i == 9 {
	//		res, _ := c.CloseAndRecv()
	//		fmt.Println("res: ", res)
	//	}
	//}
	// 服务端流
	//c, _ := client.SearchOut(context.Background(), &person.PersonReq{
	//	Name: "Hello,Minicloudsky",
	//	Age:  100,
	//})
	//for {
	//	req, err := c.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println("req: ", req)
	//}
	// 双向流
	wg := sync.WaitGroup{}
	wg.Add(2)
	c, _ := client.SearchIO(context.Background())
	counter := 0
	go func() {
		for {
			if counter != 20 {
				err := c.Send(&person.PersonReq{ // 发送
					Name: "minicloudsky",
					Age:  0,
				})
				if err != nil {
					wg.Done()
					break
				}
			} else {
				err := c.Send(&person.PersonReq{ // 发送
					Name: "结束",
					Age:  0,
				})
				if err != nil {
					wg.Done()
					break
				}
			}
			counter += 1
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			req, err := c.Recv() // 接收
			fmt.Println("req: ", req)
			if err != nil {
				wg.Done()
				break
			}
		}
	}()
	wg.Wait()
}
