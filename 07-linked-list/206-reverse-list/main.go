package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func reverseList(head *Node) *Node {
	curr := head
	var prev *Node

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func main() {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}

	newHead := reverseList(head)
	curr := newHead
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
