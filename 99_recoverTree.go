package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var pre *TreeNode
	var lastNode *TreeNode
	var x *TreeNode
	var y *TreeNode
	cur := root
	for cur != nil {
		if cur.Left == nil {
			if lastNode != nil && cur.Val <= lastNode.Val {
				y = cur
				if x == nil {
					x = lastNode
				}
			}

			lastNode = cur
			cur = cur.Right
			continue
		}

		pre = cur.Left

		//找到左子树的最右节点
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
			continue
		}

		// 最右节点与cur相等（成环情况）
		// visit cur
		if lastNode != nil && cur.Val <= lastNode.Val {
			y = cur
			if x == nil {
				x = lastNode
			}
		}

		lastNode = cur
		pre.Right = nil
		cur = cur.Right
	}

	x.Val, y.Val = y.Val, x.Val
}
