package main

/*
	四种变量声明方式
*/
import (
	"fmt"
)

// 声明全局变量，方法一、方法二、方法三可以
var username string = "minicloudsky"

// := 只能用在函数内
// userId := 100

func main() {
	// 方法一: 声明一个变量，默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)
	// 方法二: 声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b= %T\n", b)
	
	var user string = "tony"
	fmt.Println("user ", user)
	// 方法三: 在初始化时候，可以省去数据类型，通过值自动匹配当前的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)
	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc,cc)
	// 方法四: (常用的方法) 省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)
	f := "abcd"
	fmt.Println("f= ", f)
	fmt.Printf("type of f=%T\n", f)
	g := 3.14
	fmt.Println("g= ", g)
	fmt.Printf("type of g = %T\n",g)
	fmt.Printf("username=%s\n", username)
	// 声明多个变量
	var xx,yy int = 100,200
	var kk,ll = 100,"lily"
	fmt.Println(xx,yy,kk,ll)
	var (
		aaa int = 10
		bbb string = "bbb"
	)
	fmt.Println(aaa,bbb)
	fmt.Printf("cc = %s, type of cc = %T\n", bb,bb)

}