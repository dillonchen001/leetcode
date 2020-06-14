package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lLevel := 0
	if root.Left != nil {
		lLevel = findLevel(root.Left)
	}

	rLevel := 0
	if root.Right != nil {
		rLevel = findLevel(root.Right)
	}

	if lLevel-rLevel > 1 || rLevel-lLevel > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func findLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lLevel := 0
	if root.Left != nil {
		lLevel = findLevel(root.Left)
	}

	rLevel := 0
	if root.Right != nil {
		rLevel = findLevel(root.Right)
	}

	if lLevel > rLevel {
		return lLevel + 1
	} else {
		return rLevel + 1
	}
}
