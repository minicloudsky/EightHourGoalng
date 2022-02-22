package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i:=0;i<5;i++{
			c <- i
		}
		// close 可以关闭一个channel
		close(c)
		c<- 100 // panic: send on closed channel
	}()

	for {
		// ok 为 true表示channel没有关闭，如果为 false 表示 channel 已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("Main Finished")
}