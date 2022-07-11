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
		// c<- 100 // panic: send on closed channel
	}()
	// 可以使用 range 来迭代不断操作 channel
	for data := range c {
		fmt.Println("data = ", data)
	}
	fmt.Println("Main Finished")
}