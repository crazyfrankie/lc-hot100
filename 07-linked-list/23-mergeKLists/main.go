/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

最简单的思路，采用 21 题合并两个有序链表的思路，创建一个结果链表（空链表），每次循环将结果链表和当前链表传入。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for i := 0; i < len(lists); i++ {
		res = mergetTwoList(res, lists[i])
	}

	return res
}

func mergetTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}

	return dummyHead.Next
}

func main() {

}
