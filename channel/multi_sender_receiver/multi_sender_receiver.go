package main

import (
	"fmt"
	"time"
)

var senderNum int

func Send(c chan int, seq int) {
	for i := 0; i < 100*seq; i++ {
		ts := int(time.Now().UnixMilli())
		senderNum += 1
		select {
		case c <- ts:
			fmt.Printf("senderNum: %v send data %v success\n", senderNum, ts)
		default:
			fmt.Printf("chan is full, waiting...\n")
		}
		//time.Sleep(10 * time.Millisecond)
		if senderNum == 1000 {
			close(c)
			fmt.Println("---send finished, closing chan...")
		}
	}
}

func Receiver(c chan int) {
	for {
		select {
		case data, ok := <-c:
			if ok {
				fmt.Printf("receiving data %v\n", data)
			} else {
				fmt.Println("channel is closed")
			}
		}
	}
}

func main() {
	senderNum := 5
	dataChan := make(chan int, 1)
	for i := 0; i < senderNum; i++ {
		go Send(dataChan, i)
	}

	for i := 0; i < 1000; i++ {
		go Receiver(dataChan)
	}

	time.Sleep(100 * time.Second)
	close(dataChan)
}
