package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lLevel := 0
	if root.Left != nil {
		lLevel = minDepth(root.Left)
	}

	rLevel := 0
	if root.Right != nil {
		rLevel = minDepth(root.Right)
	}

	if rLevel > 0 && lLevel > 0 {
		if rLevel > lLevel {
			return lLevel + 1
		} else {
			return rLevel + 1
		}
	} else if rLevel > 0 {
		return rLevel + 1
	} else if lLevel > 0 {
		return lLevel + 1
	} else {
		return 1
	}
}
