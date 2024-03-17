package main

import (
	"fmt"
)

// TopKWithBitmap 使用位图解决 Top-K 问题
func TopKWithBitmap(nums []int, k int) []int {
	// 位图的长度，假设所有元素均小于等于100
	const bitmapLength = 101

	// 初始化位图
	bitmap := make([]int, bitmapLength)

	// 统计计数
	for _, num := range nums {
		bitmap[num]++
	}

	// 获取前 K 个元素
	topK := make([]int, 0, k)
	for i := bitmapLength - 1; i >= 0; i-- {
		if len(topK) == k {
			break
		}
		if bitmap[i] > 0 {
			topK = append(topK, i)
		}
	}

	return topK
}

func main() {
	nums := []int{4, 2, 5, 2, 5, 6, 1, 1, 4, 4, 4, 18}
	k := 3
	topK := TopKWithBitmap(nums, k)
	fmt.Println("Top", k, "elements:", topK)
}
