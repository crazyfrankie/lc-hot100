package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func isPalindrome(head *Node) bool {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = &Node{Val: 2}
	head.Next.Next.Next.Next = &Node{Val: 1}

	fmt.Println(isPalindrome(head))
}
