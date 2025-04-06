package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead

	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}

	return dummyHead.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

func main() {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	head = reverseKGroup(head, 2)
	curr := head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
