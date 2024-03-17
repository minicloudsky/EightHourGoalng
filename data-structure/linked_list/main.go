package main

// 在线面试平台。将链接分享给你的朋友以加入相同的房间。
// Author: tdzl2003<dengyun@meideng.net>
// QQ-Group: 839088532

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func BuildLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head
	for _, val := range nums[1:] {
		node := &ListNode{Val: val}
		curr.Next = node
		curr = curr.Next
	}
	return head
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(nums)
	reversedList := ReverseList(list)
	fmt.Println(list)
	fmt.Println(reversedList)
}
