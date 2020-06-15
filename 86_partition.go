package main

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	start := &ListNode{0, nil}
	start.Next = head

	pre := start
	last := start
	flag := false

	for pre.Next != nil {
		cur := pre.Next

		if cur.Val < x {
			if flag {
				pre.Next = cur.Next
				tmp := last.Next
				last.Next = cur
				cur.Next = tmp
			} else {
				pre = cur
			}
			last = cur
		} else {
			pre = cur
			flag = true
		}
	}
}
