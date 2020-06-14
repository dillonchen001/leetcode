package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var pre *TreeNode
	cur := root
	var last int
	flag := false

	for cur != nil {
		//没有左子树
		if cur.Left == nil {
			if cur.Val <= last && flag {
				return false
			}

			last = cur.Val
			flag = true
			cur = cur.Right
			continue
		}

		pre = cur.Left

		//左子树的最右节点
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
			continue
		}

		if cur.Val <= last && flag {
			return false
		}

		last = cur.Val
		flag = true

		//最右节点与cur相等
		//破环
		pre.Right = nil
		cur = cur.Right
	}
	return true
}
