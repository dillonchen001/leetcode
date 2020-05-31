package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	total := 0
	var pre *ListNode
	var tail *ListNode

	for i := 1; i <= k; i++ {
		total = 0
		pre = head
		tail = head

		for head.Next != nil {
			tail = head
			head = head.Next
			total++
		}
		total++

		head.Next = pre
		tail.Next = nil

		if i%total == k%total {
			return head
		}
	}

	return head
}
