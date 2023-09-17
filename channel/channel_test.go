package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	c := make(chan int, 1)
	c <- 10
	close(c)
	go func(c chan int) {
		value := <-c
		fmt.Println("value: ", value)
	}(c)
	// 已经关闭的channel发送值会panic
	//c <- 100
	v, ok := <-c
	fmt.Println("v: ", v)
	fmt.Println(ok)
	fmt.Println("end")
}

// deadlock 无缓冲channel需要有接收才能发送数据
func TestNoCacheChan(t *testing.T) {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func TestNoCacheChanRecv(t *testing.T) {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。
// 因此，无缓冲通道也被称为同步通道
func TestProducerConsumer(t *testing.T) {
	ch := make(chan int, 0)
	go func(c chan int) {
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Printf("生产者发送了 %d\n", i)
		}
	}(ch)

	go func(c chan int) {
		for {
			if data, ok := <-c; ok {
				fmt.Printf("消费者消费了 %d\n", data)
			}
		}
	}(ch)
	select {
	case <-time.After(1 * time.Second):
		defer close(ch)
	}
}

func TestCachedChannelSend(t *testing.T) {
	//创建一个容量为 1 的有缓冲区的通道
	ch := make(chan int, 1)
	ch <- 100
	//ch <- 100 // fatal error: all goroutines are asleep - deadlock!
	// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
	//就像你小区的快递柜只有那么个多格子，
	//格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个
	fmt.Println("send success!")
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}

// 单向发送 out 通道
func SingleSendChan(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

// 单向发送 out 通道， 单向接收 in 通道
func sqer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// 单向接收 in 通道
func SingleRecv(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func TestSingleChannel(t *testing.T) {
	out := make(chan int)
	in := make(chan int)
	go SingleSendChan(out)
	go sqer(out, in)
	SingleRecv(in)
	time.Sleep(1 * time.Second)
}

// channel 循环取值
//channel 有一个特性：close关闭之后，在发送的时候会 panic，但是在接收的时候，是可以正常接收的。
//这里介绍三种方式：
//for range
//for {}
//select{}

func TestChannelForRangeGetValue(t *testing.T) {
	ch := make(chan int, 1)
	// range c产生的迭代值为Channel中发送的值，它会一直迭代直到channel被关闭。
	//如果把close(c)注释掉，程序会一直阻塞在for …… range那一行
	go func(ch chan int) {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	go func(ch chan int) {
		for value := range ch {
			fmt.Println("reading channel value: ", value)
		}
	}(ch)
	time.Sleep(1 * time.Second)
}

func TestChannelForDeadLoopGetValue(t *testing.T) {
	ch := make(chan int, 1)
	go func(ch chan int) {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	for {
		value, ok := <-ch
		// 通道关闭后再取值ok=false
		if ok {
			fmt.Println("get value: ", value)
		}
		if !ok {
			break
		}
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// 如果有同时多个case去处理,比如同时有多个channel可以接收数据，那么Go会伪随机的选择一个case处理(pseudo-random)。
// 如果没有case需要处理，则会选择default去处理，如果default case存在的情况下。
// 如果没有default case，则select语句会阻塞，直到某个case需要处理。
// 需要注意的是，nil channel上的操作会一直被阻塞，如果没有default case,只有nil channel的select会一直被阻塞。
func TestChannelSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func TestChannelTimeout(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

func TestTimer(t *testing.T) {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")
}

func TestTimerStop(t *testing.T) {
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

// ticker是一个定时触发的计时器，它会以一个间隔(interval)往Channel发送一个事件(当前时间)，
// 而Channel的接收者可以以固定的时间间隔从Channel中读取事件。下面的例子中ticker每500毫秒触发一次，你可以观察输出的时间
func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(5 * time.Second)
}

/*
channel 异常情况总结
channel	 nil	非空	    空的	    满了	    没满
接收	     阻塞	接收值	阻塞	    接收值	接收值
发送  	 阻塞	发送值	发送值	阻塞	    发送值
关闭	panic	关闭成功，读完数据后返回零值	关闭成功，返回零值	关闭成功，读完数据后返回零值	关闭成功，读完数据后返回零值
*/
