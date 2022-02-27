package main

import (
	"fmt"
	"sort"
)

func main() {
	userIds := []int{12, 34, -12, 45, 3, 8, 9, 111, 78, -31}
	fmt.Println("before sort: ", userIds)
	// 升序排序
	sort.Ints(userIds)
	fmt.Println("after sort: ", userIds)
}
