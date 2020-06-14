package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	var levelNode [][]*TreeNode
	if root == nil {
		return result
	}

	var tmp []*TreeNode
	tmp = append(tmp, root)

	levelNode = append(levelNode, tmp)
	level := 0

	for len(levelNode) > level {
		var tmp []*TreeNode
		for _, v := range levelNode[level] {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}

			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}

		if len(tmp) > 0 {
			levelNode = append(levelNode, tmp)
		}
		level++
	}

	for j := len(levelNode) - 1; j >= 0; j-- {
		var tNum []int
		for _, v := range levelNode[j] {
			tNum = append(tNum, v.Val)
		}

		result = append(result, tNum)
	}
	return result
}
