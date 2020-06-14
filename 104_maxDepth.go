package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLevel := maxDepth(root.Left)
	rightLevel := maxDepth(root.Right)
	if leftLevel > rightLevel {
		return leftLevel + 1
	} else {
		return rightLevel + 1
	}
}
