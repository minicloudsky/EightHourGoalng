package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

// function main() {
// 	user := User{1,"minicloudsky", 18}
// 	DoFileAndMethod(user)
// }

func DoFileAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is : ", inputType.Name())

	// 获取 input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is : ", inputValue)

	// 通过 type 获取里面的字段
	// 1. 获取interface的reflect.Type,通过 Type 得到 NumFiled,进行遍历
	// 2. 得到每个field数据类型
	// 3. 通过field有一个Interface()方法，得到对应的 value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println("field: ", field, "value: ", value)
	}
	fmt.Println("-----------")
	// 通过type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
