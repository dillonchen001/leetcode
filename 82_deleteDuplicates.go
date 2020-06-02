package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	start := &ListNode{0, nil}
	start.Next = head
	pre := start

	for pre.Next != nil {
		cur := pre.Next

		if cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}

		if cur != pre.Next {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}

	return start.Next
}
