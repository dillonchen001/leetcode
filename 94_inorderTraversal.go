package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	var pre *TreeNode
	cur := root

	for cur != nil {
		//没有左子树
		if cur.Left == nil {
			result = append(result, cur.Val)
			cur = cur.Right
			continue
		}

		pre = cur.Left

		//找到左子树的最右节点
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		//最右节点
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
			continue
		}

		// 最右节点与cur相等（成环情况）
		// visit cur
		result = append(result, cur.Val)

		//破环
		pre.Right = nil

		cur = cur.Right

	}

	return result
}
