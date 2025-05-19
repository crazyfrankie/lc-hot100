/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

总体的思路就是，两个链表同时遍历，当其中一个走完了，就从另外一个的头部开始走，那么在第二次，两个链表遍历时会相遇。
例如一个长 5, 一个长 6, 5 先走完, 6 慢一个，然后 5 从 6 的开头开始, 当它走一个后, 6 也走完了, 然后它也从 5 的头开始，此时两个链表剩下的长度都是 5, 必然会相遇,
*/
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func getIntersectionNode(headA *Node, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}

func main() {
	intersect := &Node{Val: 8}
	intersect.Next = &Node{Val: 4}
	intersect.Next.Next = &Node{Val: 5}
	// 4 1 8 4 5
	headA := &Node{Val: 4}
	headA.Next = &Node{Val: 1}
	headA.Next.Next = intersect
	// 5 6 1 8 4 5
	headB := &Node{Val: 5}
	headB.Next = &Node{Val: 6}
	headB.Next.Next = &Node{Val: 1}
	headB.Next.Next.Next = intersect

	fmt.Println(getIntersectionNode(headA, headB).Val) // print 8
}
