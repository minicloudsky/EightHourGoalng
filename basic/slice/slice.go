package main

import "fmt"

func printArray(myArray []int) {
	// 值拷贝
	for index, value := range myArray {
		fmt.Println("index= ", index, ",value= ", value)
	}
	myArray[0] = 1000
}

// function main(){
// 	var myArray0[10]int
// 	myArray2 := [10]int{1,2,3,4}
// 	myArray3 := [4]int{11,22,33,44}
// 	fmt.Println(myArray0)
// 	fmt.Println(myArray2)
// 	fmt.Println(myArray3)
//  	myArray := []int{1,2,3,4,5} // 动态数组，切片 slice 引用传递
// 	fmt.Printf("myArray type is %T\n", myArray)
// 	printArray(myArray[:])
// 	for index,value := range myArray {
// 		fmt.Println("index= ",index,",value= ",value)
// 	}
// }
