package main

import "fmt"

func main(){
	// 声明slice1 是一个切片，并且初始化，默认值是1,2,3，长度len是3
	slice1 := []int{1,2,3}
	
	fmt.Println(slice1)
	fmt.Printf("len = %d, slice=%v\n", len(slice1), slice1)
}