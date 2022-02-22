package main

import (
	"fmt"
	"time"
)

func GetTime(num int) {
	fmt.Printf("id: ", num,"now time: %v\n", time.Now())
}

// func main() {
// 	// 用 go 承载一个形参为空，返回值为空的函数
// 	defer fmt.Println("A.Defer")
// 	// go func() {
// 	// 	defer fmt.Println("B.defer")
// 	// 	// 退出当前 goroutine
// 	// 	runtime.Goexit()
// 	// 	fmt.Println("B")
// 	// }()
// 	fmt.Println("A")
// 	// go func(a int,b int) bool {
// 	// 	fmt.Printf("a= ",a," b= ", b)
// 	// 	return true
// 	// }(10,20)

// 	for i:=0;i<1000;i++ {
// 		go GetTime(i)
// 	}

// 	//死循环
// 	// for {
// 	// 	time.Sleep(1 * time.Second)
// 	// }
// }