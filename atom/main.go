package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var balance = atomic.Int32{}

func add(delta int32) {
	time.Sleep(1 * time.Second)
	balance.Add(delta)
}

func main() {
	for i := 0; i < 100; i++ {
		go func(i int32) {
			add(i)
		}(int32(i))
	}
	for {
		fmt.Println(balance.Load())
	}
}
