package main

import "fmt"

func main() {
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc")
	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界
	s := []int{1, 2, 3}
	s2 := s[:0]
	fmt.Println("s: ", s)
	fmt.Println("s2: ", s2)
	num := 1 << 10
	fmt.Println("num: ", num)
}
