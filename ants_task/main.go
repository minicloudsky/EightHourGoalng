package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

type Message struct {
	Content []byte
}

var (
	LoupanProcessChan chan []Message
	loupanChan        chan []Message
)

const (
	ChanQueueSize = 100
)

func InitLianjiaChan() {
	fmt.Println("---initializing lianjia channel...")
	LoupanProcessChan = make(chan []Message, ChanQueueSize)
	loupanChan = make(chan []Message, ChanQueueSize)
}

func Receive() error {
	fmt.Println("channel start receiving messages...")
	openChannels := 1
	for {
		select {
		case msgs, ok := <-loupanChan:
			if ok {
				LoupanProcessChan <- msgs
			} else {
				openChannels--
				if openChannels == 0 {
					fmt.Println("all channel closed, exiting receive...")
					return nil
				}
			}
		}
	}
}

func poolFunc(i interface{}) {
	err := ListCityLoupan(i)
	if err != nil {
		return
	}
}

func ListCityLoupan(i interface{}) error {
	time.Sleep(1 * time.Second)
	fmt.Printf("ListCityLoupan city: %d\n", i.(int))

	return nil
}

func FetchLoupan() {
	var numbers []int
	for i := 0; i < 200; i++ {
		numbers = append(numbers, i)
	}
	go func() {
		err := Receive()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(20, func(i interface{}) {
		poolFunc(i)
		wg.Done()
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, number := range numbers {
		wg.Add(1)
		err = p.Invoke(number)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
	wg.Wait()
}

func main() {
	InitLianjiaChan()
	go FetchLoupan()
	for {

	}
}
