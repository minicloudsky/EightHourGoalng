package main

// function main(){
// 	// define a channel
// 	c := make(chan int)

// 	// 先发送下面再接收
// 	go function(){
// 		defer fmt.Println("goroutine 结束")
// 		fmt.Println("goroutine 正在运行...")
// 		c <- 666 // 将 666 发送给c
// 	}()

// 	num := <- c
// 	fmt.Println("num = ", num)
// 	fmt.Println("main goroutine end.")
// }
