package main

import "fmt"

// 如果类名首字母大写，表示其他包也可以能够访问，小写的话，当前包可见，其他包不可见

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}

// function (this Hero) SetName(newName string) {
// 	// this 是调用该方法的对象的一个副本(copy)
// 	this.Name = newName
// }

func (this *Hero) SetName(newName string) {
	// this 是调用该方法的对象的一个副本(copy)
	this.Name = newName
}

func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

// function main(){
// 	// 创建一个对象
// 	hero := Hero{Name: "zhangsan", Ad:100,Level:1}
// 	hero.Show()
// 	hero.SetName("lisi")
// 	fmt.Println("------")
// 	hero.Show()
// }
