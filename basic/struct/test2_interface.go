package main

import "fmt"

// interface{} 是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("my Func is called ...")
	fmt.Println("arg: ", arg)
	// interface 如何区分引用的底层数据类型到底是什么？
	// 给 interface{} 提供 "类型断言" 的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println(arg, " arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf(" value type is %T\n", value)
	}
}

type Store struct {
	name string
}

func main() {
	s := Store{"zhangsan"}
	myFunc(s)
	myFunc(100)
	myFunc("abc")
	myFunc(11.11)
	myFunc(0x0012)
}