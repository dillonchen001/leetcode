package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}

	for j := 0; j < len(inorder); j++ {
		if inorder[j] == root.Val {
			root.Left = buildTree(inorder[0:j], postorder[0:j])
			root.Right = buildTree(inorder[j+1:], postorder[j:len(inorder)-1])
		}
	}
	return root
}
