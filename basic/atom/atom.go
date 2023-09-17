package main

import (
	"fmt"
	"sync"
)

// Resource 临界资源
type Resource struct {
	sync.Mutex
	value int
}

var total Resource

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}
