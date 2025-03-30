/*
给定一个链表，判断其中是否存在环

双指针思路（龟兔赛跑）

关键点：为什么 fast := head.Next，因为循环初始会判断 slow 和 fast 是否相等，如果都从 head 开始，则会退出循环。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head

	fmt.Println(hasCycle(head)) // 输出 true
}
