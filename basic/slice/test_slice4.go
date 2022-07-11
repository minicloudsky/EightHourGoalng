package main

import "fmt"

func main() {
	s := []int{1,2,3} // len=3,cap= 3
	s1 := s[0:2]
	fmt.Println("s: ", s)
	fmt.Println("s1: ", s1)
	s2 := make([]int, 3)
	// copy 可以将底层数组的slice一起进行拷贝
	copy(s2,s)
	s[0] = 999
	fmt.Println("s: ", s)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
}