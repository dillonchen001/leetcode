package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	val := -1
	maxPathSumDps(root, &val)
	return val
}

func maxPathSumDps(root *TreeNode, val *int) int {
	if root == nil {
		return 0
	}

	left := maxPathSumDps(root.Left, *val)
	right := maxPathSumDps(root.Right, *val)
	lmr := root.Val + getMax(0, left) + getMax(0, right)
	ret := root.Val + getMax(0, max(left, right))
	*val = max(val, max(lmr, ret))
	return ret
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
