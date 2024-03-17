package main

import (
	"fmt"
	"os"
	"time"
)

type Job struct {
	Data chan int
	Quit chan bool
}

func Send(c chan int) {
	for i := 0; i < 1000; i++ {
		select {
		case c <- i:
			fmt.Printf("sended data: %d\n", i)
		default:
			fmt.Println("channel is full, skipping data")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func Recv(c chan int) {
	for data := range c {
		fmt.Printf("received data: %d\n", data)
	}
}

func Quit(c chan bool) {
	c <- true
}

func main() {
	dataChan := make(chan int, 1)
	quitChan := make(chan bool, 1)
	job := Job{
		Data: dataChan,
		Quit: quitChan,
	}
	fmt.Printf("job: %v\n", job)
	go Send(dataChan)
	go Recv(dataChan)
	go func() {
		select {
		case <-quitChan:
			fmt.Println("receive exit signal, quiting...")
			os.Exit(0)
		}
	}()
	time.Sleep(5 * time.Second)
	Quit(quitChan)
	time.Sleep(5 * time.Second)
}
