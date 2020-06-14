package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	var tmp []int
	pathDps(root, sum, tmp, &result)
	return result
}

func pathDps(root *TreeNode, sum int, tmp []int, result *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		tmp = append(tmp, root.Val)
		var res []int
		for _, v := range tmp {
			res = append(res, v)
		}
		*result = append(*result, res)
		return
	}

	tmp = append(tmp, root.Val)
	if root.Left != nil {
		pathDps(root.Left, sum-root.Val, tmp, result)
	}

	if root.Right != nil {
		pathDps(root.Right, sum-root.Val, tmp, result)
	}
}
