package main

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pre := head
	cur := pre.Next

	for cur != nil {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}
	return head
}
