package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	mid := findMid(head)
	node := &TreeNode{
		Val: mid.Val,
	}

	if head == mid { //just one element
		return node
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)
	return node
}

func findMid(head *ListNode) *TreeNode {
	fast := head
	slow := head
	pre := head

	for fast != nil && fast.Next != nil {
		pre = flow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if pre.Next != nil {
		pre.Next = nil
	}
	return slow
}
