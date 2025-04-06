package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	dummyHead := &ListNode{}
//	curr := dummyHead
//	list1, list2 := l1, l2
//
//	flag := false
//	for list1 != nil && list2 != nil {
//		node := &ListNode{}
//		if flag {
//			node.Val = (list1.Val + list2.Val + 1) % 10
//			flag = false
//			if list1.Val+list2.Val+1 >= 10 {
//				flag = true
//			}
//		} else {
//			node.Val = (list1.Val + list2.Val) % 10
//			if list1.Val+list2.Val >= 10 {
//				flag = true
//			}
//		}
//
//		curr.Next = node
//		curr = curr.Next
//
//		list1 = list1.Next
//		list2 = list2.Next
//	}
//
//	if list1 != nil {
//		curr.Next = list1
//		for list1 != nil {
//			if flag {
//				val := (list1.Val + 1) % 10
//				flag = false
//				if list1.Val+1 >= 10 {
//					flag = true
//				}
//				list1.Val = val
//			}
//			list1 = list1.Next
//			curr = curr.Next
//		}
//		if flag {
//			node := &ListNode{Val: 1}
//			curr.Next = node
//		}
//	} else {
//		curr.Next = list2
//		for list2 != nil {
//			if flag {
//				val := (list2.Val + 1) % 10
//				flag = false
//				if list2.Val+1 >= 10 {
//					flag = true
//				}
//				list2.Val = val
//			}
//			list2 = list2.Next
//			curr = curr.Next
//		}
//		if flag {
//			node := &ListNode{Val: 1}
//			curr.Next = node
//		}
//	}
//
//	return dummyHead.Next
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
	}

	return dummyHead.Next
}

func main() {
	head1 := &ListNode{Val: 2}
	head1.Next = &ListNode{Val: 4}
	head1.Next.Next = &ListNode{Val: 9}

	head2 := &ListNode{Val: 5}
	head2.Next = &ListNode{Val: 6}
	head2.Next.Next = &ListNode{Val: 4}
	head2.Next.Next.Next = &ListNode{Val: 9}

	res := addTwoNumbers(head1, head2)
	curr := res
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
