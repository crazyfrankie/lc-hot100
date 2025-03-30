/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 1. 递归思路
	// if list1 == nil {
	//     return list2
	// }
	// if list2 == nil {
	//     return list1
	// }
	// if list1.Val < list2.Val {
	//     list1.Next = mergeTwoLists(list1.Next,list2)
	//     return list1
	// } else {
	//     list2.Next = mergeTwoLists(list1, list2.Next)
	//     return list2
	// }

	// 2. 迭代
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
