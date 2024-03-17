package main

import (
	"fmt"
	"time"
)

type Msg struct {
	TaskId   uint
	TaskName string
	Content  string
	Balance  float64
}

func main() {
	ch := make(chan *Msg, 2) // 创建一个带缓冲的通道，容量为2

	// 启动两个 goroutine 同时向通道发送数据
	go func() {
		msg := &Msg{
			TaskId:   100,
			TaskName: "task1",
			Content:  "ping",
			Balance:  10.12,
		}
		msg2 := &Msg{
			TaskId:   2,
			TaskName: "task-2",
			Content:  "pong",
			Balance:  90.12,
		}
		ch <- msg
		ch <- msg2
	}()

	go func() {
		time.Sleep(1 * time.Second) // 等待一秒钟
		msg3 := &Msg{
			TaskId:   3,
			TaskName: "task-3",
			Content:  "hello world",
			Balance:  18.1,
		}
		ch <- msg3
	}()

	// 从通道接收数据
	fmt.Println(<-ch) // 输出 1
	fmt.Println(<-ch) // 输出 2
	fmt.Println(<-ch) // 输出 3
}
