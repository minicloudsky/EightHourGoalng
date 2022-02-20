package main

import "fmt"

// 本质是一个指针
type AnimailIF interface {
	Sleep()
	GetColor() string  // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 具体的类
type Cat struct {
	color string // 猫的颜色
}

func (this *Cat) Sleep () {
	fmt.Println("Cat.Sleep() zzzz ....")
}
func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

// 具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep () {
	fmt.Println("Dog.Sleep() zzzz ....")
}
func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimailIF) {
	animal.Sleep() // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

// func main() {
// 	// var animal AnimailIF // 接口数据类型，父类指针
// 	// animal = &Cat{"Green"}
// 	// animal.Sleep() // 调用的就是Cat的 Sleep() 方法, 多态的现象
// 	// animal  = &Dog{"Yellow"}
// 	// animal.Sleep()
// 	// fmt.Println("------------")
	
// 	cat := Cat{"green"}
// 	dog := Dog{"Yellow"}
// 	showAnimal(&cat)
// 	showAnimal(&dog)
// }