package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesDFS(1, n)
}

func generateTreesDFS(start, end int) []*TreeNode {
	var ret []*TreeNode

	if start > end {
		ret = append(ret, nil)
		return ret
	}

	for i := start; i <= end; i++ {
		leftTrees := generateTreesDFS(start, i-1)
		rightTrees := generateTreesDFS(i+1, end)

		for j := 0; j < len(leftTrees); j++ {
			for k := 0; k < len(rightTrees); k++ {
				node := &TreeNode{i, nil, nil}
				node.Left = leftTrees[j]
				node.Right = rightTrees[k]
				ret = append(ret, node)
			}
		}
	}
	return ret
}
