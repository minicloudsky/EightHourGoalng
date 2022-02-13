package main

import "fmt"

func printArray(myArray [4]string){
	// 值拷贝
	for index,value := range myArray {
		fmt.Println("index= ",index,",value= ",value)
	}
}


func main() {
	// 固定长度的数组
	var myArray [10]int
	var myArray2 =[10]int{1,2,3,4,5}
	myArray3 := [4]string{"zhangsan","lisi","wangwu"}
	for i:=0;i<10;i++{
		fmt.Println(myArray[i])
	}
	for index,value := range myArray2{
		fmt.Println("index: ", index,"value: ",value)
	}
	// 查看数据类型
	fmt.Printf("myArray types %T\n",myArray)
	fmt.Printf("myArray2 types %T\n",myArray2)
	fmt.Printf("myArray3 types %T\n",myArray3)
	printArray(myArray3)
}