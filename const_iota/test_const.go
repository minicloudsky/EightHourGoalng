package main

import "fmt"

// const定义枚举类型
const (
	// 可以在const() 添加一个关键字，iota，每行的iota都会累加1，第一行的iota默认值是0
	BEIJING = iota
	SHANGHAI
	SHENZHEN
	GUANGZHOU
)

const (
	a,b = iota+1,iota+2 // iota=0, a=iota+1,b=iota+2 a=1,b=2
	c,d // iota=1, c=iota+1,d=iota+2,c=2，d=3
	e,f // iota=2,e=iota+1,d=iota+2,e=3,f=4
	g,h = iota*2,iota*3 // iota = 3,g=iota*2,h=iota*3 g=6,h=9
	i,k // iota=4,i=iota*2,k=iota*3,i=8,k=12
)

func main() {
	// 常量(只读属性)
	const length int = 10
	fmt.Println("length = ", length)
	// 错误，常量不能修改cannot assign to length (constant 10 of type 
	// length = 100
	fmt.Println("BEIJING= ",BEIJING)
	fmt.Println("SHANGHAI= ", SHANGHAI)
	fmt.Println("SHENZHEN= ", SHENZHEN)
	fmt.Println("GUANGZHOU=", GUANGZHOU)
	fmt.Println(a,b,"\n",c,d,"\n",e,f,"\n",g,h,"\n",i,k)
	// iota 只能够配合const()一起使用，iota只有在const进行累加效果
	// var a int = iota
}