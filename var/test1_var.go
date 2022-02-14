package main
/*
	四种变量声明方式
*/
import (
	"fmt"
)

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
	fmt.Printf("cc = %s, type of cc = %T\n", bb,bb)

}