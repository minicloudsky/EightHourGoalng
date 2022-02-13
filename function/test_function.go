package main

import "fmt"

func fool(a string, b int) int{
	fmt.Println("a= ",a)
	fmt.Println("b= ",b)
	c:= 100
	return c
}

// 返回多个返回值，匿名的
func foo2(a string,b int) (int,int){
	fmt.Println("a= ",a)
	fmt.Println("b= ",b)
	return 666,777
}

// 返回多个返回值，有形参名称的
func foo3(a string,b int) (r1 int,r2 int) {
	fmt.Println("------- foo3 -----")
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)
	// r1 r2 属于foo3的形参，初始化默认的值是0
	// r1 r2 作用域空间是 foo3,整个函数整体的{} 空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ",r2)
	// 给有名称的返回值变量赋值
	r1 = 100
	r2 = 200
	return
}

func foo4(a string,b int) (r1,r2 int) {
	fmt.Println("------foo -----")
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	r1 =200
	r2 =300

	return
}

func main(){
	c := fool("abc", 555)
	fmt.Println(c)
	res1,res2 := foo2("bbb", 999)
	fmt.Println("res1= ",res1,"re2= ",res2)
	res3,res4 := foo3("qqq", 111)
	fmt.Println("res3= ", res3, "res4= ", res4)
	res5,res6 := foo4("ppp",111)
	fmt.Println("res5= ",res5,"res6= ", res6)
}