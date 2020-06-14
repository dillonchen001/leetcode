package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	ser := false
	for flag {
		var tmp []int
		ll := len(lastLevel)
		flag = false
		tStart := start
		if ser {
			for i := ll - 1; i >= tStart; i-- {
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
		} else {
			for i := ll - 1; i >= tStart; i-- {
				start++
				if lastLevel[i].Right != nil {
					tmp = append(tmp, lastLevel[i].Right.Val)
					lastLevel = append(lastLevel, lastLevel[i].Right)
					flag = true
				}

				if lastLevel[i].Left != nil {
					tmp = append(tmp, lastLevel[i].Left.Val)
					lastLevel = append(lastLevel, lastLevel[i].Left)
					flag = true
				}
			}
		}
		ser = !ser

		if len(tmp) > 0 {
			result = append(result, tmp)
		}
	}
	return result
}
