package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Wallet struct {
	sync.Mutex
	money int
}

var wallet Wallet

func Produce(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		money := rand.Intn(100)
		fmt.Printf("老公赚了%d元\n", money)
		wallet.Lock()
		wallet.money += money
		wallet.Unlock()
	}
}
func Consume(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		money := rand.Intn(100)
		fmt.Printf("老婆花了%d元\n", money)
		wallet.Lock()
		wallet.money -= money
		wallet.Unlock()
	}
}

func main() {
	fmt.Println(rand.Intn(100))
	wg := sync.WaitGroup{}
	wg.Add(2)
	go Produce(&wg)
	go Consume(&wg)
	wg.Wait()
	fmt.Println("本月结余: ", wallet.money)
}
