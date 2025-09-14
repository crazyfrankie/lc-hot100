package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func removeNthFromEnd(head *Node, n int) *Node {
	dummyHead := &Node{0, head}
	slow, fast := dummyHead, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func removeNthFromEnd1(head *Node, n int) *Node {
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head
}

func main() {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = &Node{Val: 4}
	head.Next.Next.Next.Next = &Node{Val: 5}

	newHead := removeNthFromEnd1(head, 2)
	curr := newHead
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
	// 1 2 3 5
}
