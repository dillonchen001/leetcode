package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var lastLevel []*TreeNode

	if root == nil {
		return result
	}

	tmp := []int{root.Val}
	result = append(result, tmp)
	lastLevel = append(lastLevel, root)
	start := 0
	flag := true
	for flag {
		var tmp []int
		ll := len(lastLevel)
		flag = false
		for i := start; i < ll; i++ {
			start++
			if lastLevel[i].Left != nil {
				tmp = append(tmp, lastLevel[i].Left.Val)
				lastLevel = append(lastLevel, lastLevel[i].Left)
				flag = true
			}

			if lastLevel[i].Right != nil {
				tmp = append(tmp, lastLevel[i].Right.Val)
				lastLevel = append(lastLevel, lastLevel[i].Right)
				flag = true
			}

		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

	}
	return result
}
