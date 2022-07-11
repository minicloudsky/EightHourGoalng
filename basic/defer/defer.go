package main

import "fmt"

func deferFunc() int{
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int{
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

// func fun1(){
// 	fmt.Println("fun1")
// }

// func fun2(){
// 	fmt.Println("fun2")
// }

// func fun3(){
// 	fmt.Println("fun3")
// }