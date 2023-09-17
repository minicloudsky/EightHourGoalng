package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer function called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return Func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc() // 先调用 return,后调用defer
}

func main() {
	// 写入 defer 关键字
	// defer fmt.Println("main end1")
	// defer fmt.Println("main end2")
	// fmt.Println("main hello  go1")
	// fmt.Println("main hello go2")
	// defer fun1()
	// defer fun2()
	// defer fun3()
	returnAndDefer()
}

// function fun1(){
// 	fmt.Println("fun1")
// }

// function fun2(){
// 	fmt.Println("fun2")
// }

// function fun3(){
// 	fmt.Println("fun3")
// }
