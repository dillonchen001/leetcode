package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	start := &ListNode{0, nil}
	start.Next = head

	last := start
	pre := start

	i := 0
	for pre.Next != nil {
		cur := pre.Next
		i++
		if i >= m {
			pre.Next = cur.Next
			tmp := last.Next
			last.Next = cur
			cur.Next = tmp
			if i == n {
				return start.Next
			}
		} else {
			last = last.Next
			pre = pre.Next
		}
	}
	return start.Next
}
