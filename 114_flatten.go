package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	result := root
	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left
	if left != nil {
		result = flattenDps(left)
	}
	if result != nil && right != nil {
		result.Right = right
		result = flattenDps(right)
	}
}

func flattenDps(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}

	result := node
	right := node.Right
	if node.Left != nil {
		node.Right = node.Left
		node.Left = nil
		result = node.Right
		result = flattenDps(node.Right)
		fmt.Println(result.Val)
	}

	if right != nil {
		result.Right = right
		result = flattenDps(node.Right)
		fmt.Println(result.Val)
	}

	return result
}
