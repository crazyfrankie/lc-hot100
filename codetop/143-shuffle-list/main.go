package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func shuffleList(head *ListNode) {
	var shuffle func(head *ListNode)
	shuffle = func(head *ListNode) {
		if head == nil || head.Next == nil || head.Next.Next == nil {
			return
		}
		curr := head
		for curr.Next.Next != nil {
			curr = curr.Next
		}
		tail := curr.Next
		tail.Next = head.Next
		head.Next = tail
		curr.Next = nil

		shuffle(tail.Next)
	}

	shuffle(head)
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	shuffleList(head)

	curr := head
	res := make([]int, 0, 5)
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	fmt.Println(res)
}
