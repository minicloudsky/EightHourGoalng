package main

import "fmt"

//var a string
//
//func f() {
//	fmt.Println(a)
//}
//
//func hello() {
//	a = "hello world"
//	go f()
//}
//
//func main() {
//	hello()
//}

func main() {
	done := make(chan int, 2)
	go func() {
		fmt.Println("hello world")
		//data := <-done
		done <- 111
		fmt.Println("data: ", <-done)
	}()

	done <- 1
}
