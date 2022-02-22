package main

import "fmt"

func fibnacii(c,quit chan int) {
	x,y := 1,1
	for {
		select {
		case c <- x:
			// 如果 c可写，则该 case 就会进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}

	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("fib(%d) is %d\n", i,<-c)
		}
		quit <- 0
	}()
	// main go
	fibnacii(c, quit)
}