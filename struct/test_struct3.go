package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (this *Human) Eat() {
	fmt.Println("human.Eat() ...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk() ...")
}

type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()")
}

func (this *SuperMan) Print() {
	fmt.Println("namne = ", this.name)
	fmt.Println("sex= ", this.sex)
	fmt.Println("level = ", this.level)
}

// func main() {
// 	h := Human{"zhangsan", "female"}
// 	h.Eat()
// 	h.Walk()
// 	// s := SuperMan{Human{"zhangsan", "male"},100}
// 	var s SuperMan
// 	s.name = "liyuanfang"
// 	s.sex = "male"
// 	s.level = 999
// 	s.Walk()
// 	s.Eat()
// 	s.Fly()
// 	s.Print()
// }