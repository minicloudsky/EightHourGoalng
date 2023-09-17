package main

import "fmt"

// Duck
// 接口是一组方法的集合,实现了这些方法，就实现了这个接口，就是多态了
type Duck interface {
	Quack()  // 鸭子叫
	DuckGo() // 鸭子走
}

type Chicken struct {
}

func (c Chicken) IsChicken() bool {
	fmt.Print("我是小鸡")
	return false
}

func (c Chicken) Quack() {
	fmt.Println("嘎嘎")
}

func (c Chicken) DuckGo() {
	fmt.Println("大摇大摆的走")
}

func DoDuck(d Duck) {
	d.Quack()
	d.DuckGo()
}

func main() {
	c := Chicken{}
	DoDuck(c) // 类似其他语言的多态
}
